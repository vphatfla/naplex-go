package main

import (
	"context"
	"log"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/vphatfla/naplex-go/backend/internal/config"
	"github.com/vphatfla/naplex-go/backend/internal/middleware"
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
	
	// Modules declare
	authModule := auth.
	// chi router
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)

	// Auth
	r.Route("/auth", func(r chi.Router) {
		
	})
}
