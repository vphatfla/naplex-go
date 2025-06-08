package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

type GoogleUserInfo struct {
	ID        string `json:"id"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	Email string `json:"email"`
}

func GetGoogleUserInfo(ctx context.Context, token *oauth2.Token) (*GoogleUserInfo, error) {
	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))

	res, err := client.Get(GoogleUserInfoLink)
	if err != nil {
		return nil, fmt.Errorf("GetGoogleUserInfo: error GET -> %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("GetGoogleUserInfo: error GET status NOT OK -> %v", body)
	}

	var u GoogleUserInfo
	if err := json.NewDecoder(res.Body).Decode(&u); err != nil {
		return nil, fmt.Errorf("GoogleUserInfo: error decode user info -> %v", err)
	}

	return &u, nil
}

// GenerateStateToken creates a secure random state token
func GenerateStateToken() (string, error) {
    bytes := make([]byte, 32)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(bytes), nil
}

