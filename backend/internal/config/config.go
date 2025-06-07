package config

import "os"

type Config struct {
	Auth AuthConfig
}

type AuthConfig struct {
	GoogleClientID string
	GoogleClientSecret string
}

func LoadConfig() *Config {
	return &Config{
		AuthConfig{
			GoogleClientID: os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
			GoogleClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		},
	}
}
