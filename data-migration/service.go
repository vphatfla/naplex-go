package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vphatfla/naplex-go/data-migration/database/dstDB"
	"github.com/vphatfla/naplex-go/data-migration/database/srcDB"
)

const (
	multipleChoicesPattern = `[A-Z]\.\s([^A-Z]+(?:[A-Z](?!\.\s)[^A-Z]*)*)`
)

// Migration service
type Service struct {
	srcRepository *srcDB.Queries
	dstRepository *dstDB.Queries
	batchSize     int
	numWorker     int
	logger        *log.Logger
	logFile       *os.File
}

type result struct {
	ids []int32
	err error
	msg string
}

func (r *result) ToString() string {
	return fmt.Sprintf("Ids lens = %d start = %d end = %d with error = %v and msg = %s", len(r.ids), r.ids[0], r.ids[len(r.ids)-1], r.err, r.msg)
}

func NewService(srcPool *pgxpool.Pool, dstPool *pgxpool.Pool, batchSize int, numWorker int, dir string) (*Service, error) {
	l, err := NewLogWriter(dir, "Service")
	if err != nil {
		return nil, err
	}
	return &Service{
		srcRepository: srcDB.New(srcPool),
		dstRepository: dstDB.New(dstPool),
		batchSize:     batchSize,
		numWorker:     numWorker,
		logger:        l.Logger,
		logFile:       l.File,
	}, nil
}

func (s *Service) StartMigration() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		if err := s.logFile.Close(); err != nil {
			log.Printf("Closing log file error :%s", err) // only os.Stdout
		}
	}()

	defer cancel()

	// handle interrupt with Ctrl + C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
		case <-sigChan:
			s.logger.Println("Received interrupt signal, cancelling migration")
			cancel()
		case <-ctx.Done():
			// do nothing
		}
	}()

	ids, err := s.srcRepository.GetAllIds(ctx)
	if err != nil {
		return err
	}

	idBatch := make(chan []int32)
	results := make(chan *result)
	var wg sync.WaitGroup

	// receiver of results chan
	// this need to be ready in order for the worker to send the result in
	resultDone := make(chan struct{})
	go func() {
		for r := range results {
			s.logger.Println(r.ToString())
		}
		s.logger.Printf("Results processing FINISHED")
		close(resultDone)
	}()

	// receiver of idBatch, sender of results chan
	// spin up worker
	// worker is a reciever of idBatch channel need to be ready before the main routine can send batch into the channel
	for i := 0; i < s.numWorker; i += 1 {
		wg.Add(1)
		go s.worker(ctx, i, idBatch, results, &wg)
	}

	// sender to idBatch
	// divide ids into batches and send to idBatch channel
	for i := 0; i < len(ids); i += s.batchSize {
		end := min(i+s.batchSize, len(ids))
		select {
		case idBatch <- ids[i:end]:
		case <-ctx.Done():
			// handle interrupt
			close(idBatch)
			wg.Wait()
			close(results)
			<-resultDone
			return fmt.Errorf("migration cancelled at batch start %d end  %d : %w", i, end, ctx.Err())
		}
	}

	close(idBatch)
	wg.Wait()
	close(results)
	<-resultDone
	return nil
}

func (s *Service) worker(ctx context.Context, workerId int, idBatch chan []int32, results chan *result, wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Printf("Worker %d recover from panic %v", workerId, r)
		}
	}()
	defer wg.Done()
	s.logger.Printf("Worker %d start ", workerId)
	for {
		select {
		case ids, ok := <-idBatch:
			if !ok {
				s.logger.Printf("Worker %d: no more work, shutting down", workerId)
				return
			}
			s.logger.Printf("Worker %d gets ids len %d id range from %d to %d", workerId, len(ids), ids[0], ids[len(ids)-1])
			srcQs, err := s.srcRepository.GetProcessedQuestionsInBatch(ctx, ids)
			s.logger.Printf("Worker %d queried srcDB get count = %d", workerId, len(srcQs))
			if err != nil {
				results <- &result{
					ids: ids,
					err: err,
				}
				continue // move to next batch
			}
			var dstQs []dstDB.Question
			for _, srcQ := range srcQs {
				dstQs = append(dstQs, processQuestion(&srcQ))
			}

			var dstQsParams []dstDB.CreateQuestionsBatchParams
			for _, dstQ := range dstQs {
				p := dstDB.CreateQuestionsBatchParams{
					Title:           dstQ.Title,
					Question:        dstQ.Question,
					MultipleChoices: dstQ.MultipleChoices,
					CorrectAnswer:   dstQ.CorrectAnswer,
					Explanation:     dstQ.Explanation,
					Keywords:        dstQ.Keywords,
					Link:            dstQ.Link,
				}
				dstQsParams = append(dstQsParams, p)
			}
			count, err := s.dstRepository.CreateQuestionsBatch(ctx, dstQsParams)

			r := &result{
				ids: ids,
				err: err,
				msg: fmt.Sprintf("COUNT = %d", count),
			}

			results <- r
			s.logger.Printf("Worker %d FINISHED batch size %d range from %d to %d with results msg = %s", workerId, len(ids), ids[0], ids[len(ids)-1], r.msg)
		case <-ctx.Done():
			s.logger.Printf("Worker %d: context canceled/done from parents, shutting down", workerId)
			return
		}
	}
}

func processQuestion(srcQ *srcDB.ProcessedQuestion) dstDB.Question {
	dstQ := dstDB.Question{
		Title:       srcQ.Title,
		Question:    srcQ.Question,
		Explanation: srcQ.Explanation,
		Keywords:    srcQ.Keywords,
		Link:        srcQ.Link,
	}

	// options := parseMultipleChoices(srcQ.MultipleChoices)
	// dstQ.MultipleChoices = strings.Join(options, ",")
	// dstQ.CorrectAnswer = strings.TrimSpace(srcQ.CorrectAnswer[2:])

	dstQ.MultipleChoices = srcQ.MultipleChoices;
	dstQ.CorrectAnswer = srcQ.CorrectAnswer;
	return dstQ
}

func parseMultipleChoices(text string) []string {
	// Pattern to match option markers: A. B. C. etc.
	optionPattern := regexp.MustCompile(`\b[A-Z]\.\s`)

	// Find all positions where options start
	matches := optionPattern.FindAllStringIndex(text, -1)
	if len(matches) == 0 {
		return []string{}
	}

	var options []string

	// Extract text between option markers
	for i := 0; i < len(matches); i++ {
		start := matches[i][1] // Start after "A. "

		var end int
		if i < len(matches)-1 {
			end = matches[i+1][0] // End before next "B. "
		} else {
			end = len(text) // Last option goes to end
		}

		option := strings.TrimSpace(text[start:end])
		if option != "" {
			options = append(options, option)
		}
	}

	return options
}
