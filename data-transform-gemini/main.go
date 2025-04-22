package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/vphatfla/naplex-go/data-transform-gemini/db"
)

func main() {
	fmt.Println("Hello from naplex data transformer")
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://naplex_user:password@localhost:5432/app_db")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)
	
	r, err := queries.GetRawQuestioniByID(ctx, 2)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Result = %v", r)
}
