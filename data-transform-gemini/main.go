package main

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
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
	log.Println("Database info >>>  Successfully created database connection pool!")
	defer pool.Close()

	client, err := gemini.NewClient(ctx, cfg)
	if err != nil {
		log.Fatal(nil)
	}
	defer client.Close()
	log.Println("Gemini info >>> Successfully initilize gemini ai client!")

	model := client.GenerativeModel(cfg.Gemini.Model)
	res, err := model.GenerateContent(ctx, genai.Text("Tell me something about horse"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response from gemini");
	for _, c := range res.Candidates {
		if c != nil {
			log.Printf("Token count for this candidates %v", c.TokenCount)
			for _, part := range c.Content.Parts {
				log.Println(part)
			}
		}
	}
	/* queries := db.New(pool)

	r, err := queries.GetRawQuestioniByID(ctx, 2)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Result = %v", r)
	*/
}
