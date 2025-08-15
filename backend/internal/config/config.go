package config

import (
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	OAuth2Config *oauth2.Config
	CookieSecret []byte
	DBConfig     *DBConfig
	FrontEndRedirectPageURI string
}

type DBConfig struct {
	Username string
	Password string
	DBName   string
	Host     string
	Port     string
}

func (c *DBConfig) ToURLString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.Username, c.Password, c.Host, c.Port, c.DBName)
}

func LoadConfig() *Config {
	/* if os.Getenv("DOCKER_PROD") == "true" {
		os.Setenv("POSTGRES_HOST", "naplex-postgres-db")
	} */
	return &Config{
		OAuth2Config: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
			Endpoint:     google.Endpoint,
			Scopes: []string{
				"openid",
				"email",
				"profile",
			},
			RedirectURL: os.Getenv("GOOGLE_REDIRECT_URI"),
		},
		CookieSecret: []byte(os.Getenv("COOKIE_SECRET")),
		DBConfig: &DBConfig{
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DB"),
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
		},
		FrontEndRedirectPageURI: os.Getenv("FRONTEND_REDIRECT_URI"),
	}
}
