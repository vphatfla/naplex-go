package main

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/vphatfla/naplex-go/backend/internal/auth"
	"github.com/vphatfla/naplex-go/backend/internal/config"
	"github.com/vphatfla/naplex-go/backend/internal/shared/database"
)

func main() {
	log.Printf("Hello from naplex go backend")

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed loading env %v", err);
	}

	ctx := context.Background()
	config := config.LoadConfig()

	pool, err := database.NewPool(ctx, config)
	if err != nil {
		log.Fatalf("Failed to create database pool -> %v", err)
	}
	defer pool.Close()

	queries := database.New(pool)
	authHandler := auth.NewAuthHandler(config, queries)
	authRouter := authHandler.RegisterRouter()

	http.Handle("/auth/", http.StripPrefix("/auth", authRouter))

	log.Fatal(http.ListenAndServe(":8080", nil))
	/*
	querier := database.New(pool)

	querier.CheckUserExistsByEmail(ctx, "abc")
	*/
}
