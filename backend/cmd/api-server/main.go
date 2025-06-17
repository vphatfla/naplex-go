package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/vphatfla/naplex-go/backend/internal/auth"
	"github.com/vphatfla/naplex-go/backend/internal/config"
	"github.com/vphatfla/naplex-go/backend/internal/middleware"
	"github.com/vphatfla/naplex-go/backend/internal/shared/database"
	"github.com/vphatfla/naplex-go/backend/internal/user"
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
	authModule := auth.NewModule(config, queries)
	userModule := user.NewModule(q)
	// chi router
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)

	// Auth
	r.Route("/auth", func(r chi.Router) {
		r.Get("/google/login", authModule.Handler.HandleGoogleLogin)
		r.Get("/google/callback", authModule.Handler.HandleGoogleCallback)
		r.Get("/logout", authModule.Handler.HandleLogout)
	})
	
	// User
	r.Route("/user", func(r chi.Router) {
		r.Get("/profile", userModule.Handler.)
	})

	port := ":8080"
	log.Printf("Http Server starting on %v", port)
	log.Fatal(http.ListenAndServe(port, r))
}
