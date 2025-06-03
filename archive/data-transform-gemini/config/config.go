package config

import (
	"fmt"
	"os"
)

type Config struct {
	DB PostgresConfig
	Gemini GeminiConfig
}

type PostgresConfig struct {
	Username string
	Password string
	DBName string
	Host string
	Port string
}

type GeminiConfig struct {
	APIKEY string
	Model string
}

func (pc *PostgresConfig) ToURLString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pc.Username, pc.Password, pc.Host, pc.Port, pc.DBName)
}

func LoadConfig() (*Config) {
	return &Config{
		PostgresConfig{
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName: os.Getenv("POSTGRES_DB"),
			Host: os.Getenv("POSTGRES_HOST"),
			Port: os.Getenv("POSTGRES_PORT"),
		},
		GeminiConfig{
			APIKEY: os.Getenv("GEMINI_API_KEY"),
			Model: os.Getenv("GEMINI_MODEL"),
		},
	}
}
