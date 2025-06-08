package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vphatfla/naplex-go/backend/internal/config"
)

func main() {
	log.Printf("Hello from naplex go backend")

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed loading env %v", err);
	}

	config := config.LoadConfig()
	log.Printf("Client id %s", config.OAuth2Config.ClientID)
	log.Printf("Cookie = %s", config.CookieSecret)
}
