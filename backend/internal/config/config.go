package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	OAuth2Config *oauth2.Config
	CookieSecret []byte
}

func LoadConfig() *Config {
	return &Config{
		OAuth2Config: &oauth2.Config{
			ClientID: os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
			Endpoint: google.Endpoint,
			Scopes: []string{
                "openid",
                "email",
                "profile",
            },
			RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		},
		CookieSecret: []byte(os.Getenv("COOKIE_SECRET")),
	}
}
