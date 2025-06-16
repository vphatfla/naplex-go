package user

import "github.com/vphatfla/naplex-go/backend/internal/shared/database"

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
