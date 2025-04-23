package main

import (
	"context"
	"encoding/json"
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

	queries := db.New(pool)

	r, err := queries.GetRawQuestioniByID(ctx, 2)
	if err != nil {
		log.Fatal(err)
	}

	model := client.GenerativeModel(cfg.Gemini.Model)

	prompt := `You are an AI assistant tasked with processing pharmacy exam questions. Given the following pharmacy exam question text, clean the formatting and extract key information into a structured JSON object.

	Rules:
	1. Remove extraneous characters like '+' line endings and any database artifacts
	2. Analyze the content to identify question components
	3. Output a single valid JSON object with the following fields:
	- "title": Create a concise, descriptive title based on the question content (e.g., "Vancomycin Dosing Adjustment")
	- "question": Extract the full question text including patient case and the specific question being asked
	- "multipleChoices": Format as "A. [option A text] B. [option B text] C. [option C text] D. [option D text]" with the exact lettering and formatting preserved
	- "correctAnswer": Provide the correct answer as the capitalized letter followed by the answer text (e.g., "B. 1000 mg IV q12h")
	- "explanation": Extract the full explanation/rationale for the correct answer
	- "keywords": Identify 2-4 important pharmacy-related keywords/topics from the question (e.g., "vancomycin, dosing, trough levels, antimicrobial therapy")

	Please process the following text into the specified JSON format` + r.RawQuestion
	res, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response from gemini");
	for _, c := range res.Candidates {
		if c != nil {
			log.Printf("Token count for this candidates %v", c.TokenCount)
			for _, part := range c.Content.Parts {
				log.Println(part)

				if txt, ok := part.(genai.Text); ok {
					var temp  db.ProcessedQuestion
					err := json.Unmarshal([]byte(txt), &temp)
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("%v", temp)

				} else {
					log.Println("Can't cast res to txt")
				}
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
