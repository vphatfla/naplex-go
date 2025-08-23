package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/vphatfla/naplex-go/backend/internal/auth"
	"github.com/vphatfla/naplex-go/backend/internal/shared/database"
)

// Service is a struct that implements business logic for users
// q stands for querier, or repository, which is used to store/retrieve data
type Service struct {
	q *database.Queries
}

func NewService(queries *database.Queries) *Service {
	return &Service{
		q: queries,
	}
}

// CreateOrUpdateUser take GoogleUserInfo as an argurment and call the querier to create/update the user
func (s *Service) CreateOrUpdateUser(ctx context.Context, gU *auth.GoogleUserInfo) (*database.User, error) {
	params := &database.CreateOrUpsertUserParams{
		GoogleID:  gU.ID,
		Email:     gU.Email,
		Name:      gU.Name,
		FirstName: pgtype.Text{String: gU.FirstName, Valid: true},
		LastName:  pgtype.Text{String: gU.LastName, Valid: true},
		Picture:   pgtype.Text{String: gU.Picture, Valid: true},
	}

	u, err := s.q.CreateOrUpsertUser(ctx, *params)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *Service) GetUserProfile(ctx context.Context, id int32) (*database.User, error) {
	u, err := s.q.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *Service) UpdateUserProfile(ctx context.Context, u *database.User) (*database.User, error) {
	params := &database.UpdateUserProfileParams{
		ID:        u.ID,
		Name:      u.Name,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Picture:   u.Picture,
	}

	newU, err := s.q.UpdateUserProfile(ctx, *params)
	if err != nil {
		return nil, err
	}
	return &newU, nil
}
