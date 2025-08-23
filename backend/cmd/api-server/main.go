package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/vphatfla/naplex-go/backend/internal/auth"
	"github.com/vphatfla/naplex-go/backend/internal/config"
	"github.com/vphatfla/naplex-go/backend/internal/dataTransfer"
	"github.com/vphatfla/naplex-go/backend/internal/middleware"
	"github.com/vphatfla/naplex-go/backend/internal/question"
	"github.com/vphatfla/naplex-go/backend/internal/shared/database"
	"github.com/vphatfla/naplex-go/backend/internal/user"
)

func main() {
	log.Printf("Hello from naplex go backend")

	if os.Getenv("DOCKER") == "true" {
		log.Println("Docker container running, no need to use godotenv to load")
	} else if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed loading env %v", err)
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
	userModule := user.NewModule(queries)
	questionModule := question.NewModule(queries)
	dataTransferModule := dataTransfer.NewModule(queries)
	// Middleware declare
	authM := auth.NewMiddleware(config)
	// chi router
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.AddCors)
	// Auth
	r.Route("/auth", func(r chi.Router) {
		r.Get("/google/login", authModule.Handler.HandleGoogleLogin)
		r.Get("/google/callback", authModule.Handler.HandleGoogleCallback)
		r.Get("/logout", authModule.Handler.HandleLogout)
	})

	// User
	r.Route("/user", func(r chi.Router) {
		r.Use(authM.RequireAuth)
		r.Get("/profile", userModule.Handler.HandleGetUser)
		r.Post("/profile", userModule.Handler.HandleUpdateUser)
	})

	// Question
	r.Route("/question", func(r chi.Router) {
		r.Use(authM.RequireAuth)

		r.Get("/", questionModule.Handler.HandleGetQuestion)
		r.Post("/", questionModule.Handler.HandleCreateOrUpdateUserQuestion)
		r.Get("/passed", questionModule.Handler.HandlerGetAllPassedQuestion)
		r.Get("/failed", questionModule.Handler.HandlerGetAllFailedQuestion)

		r.Get("/daily", questionModule.Handler.HandlerGetRandomDailyQuestion)
	})

	// Strictly Internal
	r.Route("/internal", func(r chi.Router) {
		r.Use(authM.RequireAPIKEYInternal)
		r.Post("/data/upload", dataTransferModule.Handler.HandleInsertQuestionsBatch)
	})

	port := ":8080"
	log.Printf("Http Server starting on %v", port)
	log.Fatal(http.ListenAndServe(port, r))
}
