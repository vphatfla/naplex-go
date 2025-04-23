package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/vphatfla/naplex-go/data-transform-gemini/config"
	"github.com/vphatfla/naplex-go/data-transform-gemini/db"
)

func main() {
	fmt.Println("Hello from naplex data transformer")
	err := godotenv.Load()

	cfg := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	pool, err := db.NewPool(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database info >>>  Successfully created database connection pool!")
	defer pool.Close()

	queries := db.New(pool)

	r, err := queries.GetRawQuestioniByID(ctx, 2)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Result = %v", r)
}
