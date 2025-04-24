package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/vphatfla/naplex-go/data-transform-gemini/config"
	"github.com/vphatfla/naplex-go/data-transform-gemini/db"
	"github.com/vphatfla/naplex-go/data-transform-gemini/gemini"
)

func main() {
	log.Println("Hello from naplex data transformer")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	cfg := config.LoadConfig()

	ctx := context.Background()

	pool, err := db.NewPool(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database >>>  Successfully created database connection pool!")
	defer pool.Close()

	client, err := gemini.NewClient(ctx, cfg)
	if err != nil {
		log.Fatal(nil)
	}
	defer client.Close()
	log.Println("Gemini >>> Successfully initilize gemini ai client!")

	queries := db.New(pool)

	t, err  := queries.CountRawQuestion(ctx)
	if err != nil {
		log.Fatal(err)
	}
	total := int(t)
	total = 10

	var wg sync.WaitGroup

	for start  := 1; start <= total; start += 10{
		end := min(start + 9, total)
		rawQuestions, err := queries.GetRawQuestionWithRange(ctx,
			db.GetRawQuestionWithRangeParams{
				ID: int32(start),
				ID_2: int32(end),
			})
		if err != nil {
			log.Printf("Error getting raw questions from %v to %v -->  %v", start, end, err)
			continue
		}

		for _,rQ := range rawQuestions {
			if strings.Contains(rQ.Title, "Preparing for Success") {
				log.Printf("Skipping id %v title %v", rQ.ID, rQ.Title)
				continue
			}
			wg.Add(1)
			go func(rQ db.RawQuestion) {
				log.Printf("Requesting content for raw questions id %v title %v", rQ.ID, rQ.Title)
				m := gemini.NewModelJson(client, cfg)
				txt, err := gemini.GetContent(ctx, m, rQ.RawQuestion)
				if err != nil {
					log.Printf("err -> %v", err)
					return
				}

				// log.Printf("id %v & title %v --> \n%v", rQ.ID, rQ.Title, string(txt))
				var temp db.ProcessedQuestion
				temp.Title = rQ.Title

				err = json.Unmarshal([]byte(txt), &temp)
				if err != nil {
					log.Panic(err)
				}
				log.Printf("Successfully Unmarshal for raw question id = %v", rQ.ID)

				id, err := queries.InsertProcessedQuestion(ctx, db.InsertProcessedQuestionParams{
					Title: temp.Title,
					Question: temp.Question,
					MultipleChoices: temp.MultipleChoices,
					CorrectAnswer: temp.CorrectAnswer,
					Explanation: temp.Explanation,
					Keywords: temp.Keywords,
				})
				if err != nil {
					log.Panic(err)
				}
				log.Printf("Inserted data into processed table id %v", id)
				wg.Done()
				return
			}(rQ)
		}
	}

	wg.Wait()
	return
}
