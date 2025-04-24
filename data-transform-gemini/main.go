package main

import (
	"context"
	"encoding/json"
	"log"
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
	total = 5

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
			wg.Add(1)
			go func(rq db.RawQuestion) {
				log.Printf("Requesting content for raw questions id %v", rQ.ID)
				m := gemini.NewModelJson(client, cfg)
				txt, err := gemini.GetContent(ctx, m, rQ.RawQuestion)
				if err != nil {
					log.Printf("err -> %v", err)
					return
				}

				var temp db.ProcessedQuestion
				err = json.Unmarshal([]byte(txt), &temp)
				if err != nil {
					log.Panic(err)
				}
				log.Printf("Successfully Unmarshal for raw question id = %v", temp.ID)
				log.Printf("%+v", temp)
				wg.Done()
				return
			}(rQ)
		}
	}

	wg.Wait()
	return
}
