package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/vphatfla/naplex-go/backend/internal/auth"
	"github.com/vphatfla/naplex-go/backend/internal/shared/database"
	"golang.org/x/oauth2"
)

type Service struct {
	q *database.Queries
}

func NewService(q *database.Queries) *Service {
	return &Service{
		q: q,
	}
}
type GoogleUserInfo struct {
	ID        string `json:"id"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	Email string `json:"email"`
}

func (s *Service) GetGoogleUserInfo(ctx context.Context, token *oauth2.Token) (*GoogleUserInfo, error) {
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
func (s *Service) GenerateStateToken() (string, error) {
    bytes := make([]byte, 32)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(bytes), nil
}

// CreateOrUpdateUser take GoogleUserInfo as an argurment and call the querier to create/update the user
func (s *Service) CreateOrUpdateUser(ctx context.Context, gU *GoogleUserInfo) (*database.User, error) {
	params := &database.CreateOrUpsertUserParams{
		GoogleID: gU.ID,
		Email: gU.Email,
		Name: gU.Name,
		FirstName: pgtype.Text{String: gU.FirstName, Valid: true},
		LastName: pgtype.Text{String: gU.LastName, Valid: true},
		Picture: pgtype.Text{String: gU.Picture, Valid: true},
	}

	u, err := s.q.CreateOrUpsertUser(ctx, *params)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
