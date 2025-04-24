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
	model.ResponseMIMEType = "application/json"
	prompt := `Process this pharmacy exam question into a JSON object:
	- "title": Brief descriptive title (e.g., "Vancomycin Dosing")
	- "question": Full question text with patient case
	- "multipleChoices": Format as "A. [text] B. [text] C. [text] D. [text]"
	- "correctAnswer": Letter + text (e.g., "B. 1000 mg IV q12h")
	- "explanation": Rationale for correct answer
	- "keywords": Single string with 2-4 terms separated by commas only without spaces (e.g., "vancomycin,dosing,antimicrobial")

	Remove any '+' line endings and database artifacts.`+ r.RawQuestion
	res, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response from gemini");
	for _, c := range res.Candidates {
		if c != nil {
			log.Printf("Token count for this candidates %v", c.TokenCount)
			for _, part := range c.Content.Parts {
				// log.Println(part)

				if txt, ok := part.(genai.Text); ok {
					var temp  db.ProcessedQuestion
					err := json.Unmarshal([]byte(txt), &temp)
					if err != nil {
						log.Fatal(err)
					}
					
					jsonData, err := json.Marshal(temp)
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("%v", string(jsonData))
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
