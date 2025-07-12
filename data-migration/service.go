package main

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vphatfla/naplex-go/data-migration/database/dstDB"
	"github.com/vphatfla/naplex-go/data-migration/database/srcDB"
)

const (
	multipleChoicesPattern = `(?:^|\s)[A-Z]\.\s(.*?)(?=\s[A-Z]\.\s|$)`
)

// Migration service
type Service struct {
	srcRepository *srcDB.Queries
	dstRepository *dstDB.Queries
	batchSize     int
	numWorkder    int
}

type result struct {
	ids []int32
	err error
	msg string
}

func (r *result) ToString() string {
	return fmt.Sprintf("Ids lens = %d start = %d end = %d with error = %v and msg = %s", len(r.ids), r.ids[0], r.ids[len(r.ids)-1], r.err, r.msg)
}
func NewService(srcPool *pgxpool.Pool, dstPool *pgxpool.Pool, bacthSize int, numWorkder int) *Service {
	return &Service{
		srcRepository: srcDB.New(srcPool),
		dstRepository: dstDB.New(dstPool),
		batchSize:     bacthSize,
		numWorkder:    numWorkder,
	}
}

func (s *Service) StartMigration() error {

	ids, err := s.srcRepository.GetAllIds(context.Background())
	if err != nil {
		return err
	}

	idBatch := make(chan []int32)
	results := make(chan *result)
	var wg sync.WaitGroup

	// spin up worker
	for i := 0; i < s.numWorkder; i += 1 {
		wg.Add(1)
		go s.worker(i, idBatch, results, &wg)
	}

	// divide ids into batches and send to idBatch channel
	start := 0
	for {
		if start >= len(ids) {
			break
		}

		lenList := min(len(ids)-start, s.batchSize)
		var list []int32

		for i := start; i < start+lenList; i += 1 {
			list = append(list, ids[i])
		}

		idBatch <- list
		start += lenList
	}
	close(idBatch)
	wg.Wait()

	for r := range results {
		log.Println(r.ToString())
	}
	return nil
}

func (s *Service) worker(workerId int, idBatch chan []int32, results chan *result, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Worker %d start ", workerId)
	re := regexp.MustCompile(multipleChoicesPattern)
	for ids := range idBatch {
		log.Printf("Worker %d STARTING get  batch size %d range from %d to %d", workerId, len(ids), ids[0], ids[len(ids)-1])
		//logic
		srcQs, err := s.srcRepository.GetProcessedQuestionsInBatch(context.Background(), ids)
		if err != nil {
			results <- &result{
				ids: ids,
				err: err,
			}
		}

		var dstQs []dstDB.Question
		for _, srcQ := range srcQs {
			dstQs = append(dstQs, processQuestion(&srcQ, re))
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
		count, err := s.dstRepository.CreateQuestionsBatch(context.Background(), dstQsParams)

		r := &result{
			ids: ids,
			err: nil,
			msg: fmt.Sprintf("COUNT = %d", count),
		}

		results <- r
		log.Printf("Worker %d FINSIHED batch size %d range from %d to %d with results msg = %s", workerId, len(ids), ids[0], ids[len(ids)-1], r.msg)
	}

	log.Printf("Worker %d routines FINISHED, closing", workerId)
}

func processQuestion(srcQ *srcDB.ProcessedQuestion, re *regexp.Regexp) dstDB.Question {
	dstQ := dstDB.Question{
		Title:       srcQ.Title,
		Question:    srcQ.Question,
		Explanation: srcQ.Explanation,
		Keywords:    srcQ.Keywords,
		Link:        srcQ.Link,
	}

	matches := re.FindAllStringSubmatch(srcQ.MultipleChoices, -1)

	var options []string
	for _, match := range matches {
		if len(match) > 1 {
			options = append(options, strings.TrimSpace(match[1]))
		}
	}

	dstQ.MultipleChoices = strings.Join(options, ",")
	dstQ.CorrectAnswer = strings.TrimSpace(srcQ.CorrectAnswer[2:])

	return dstQ
}
