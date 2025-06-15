package users

import (
	"github.com/vphatfla/naplex-go/backend/internal/shared/database"
	"github.com/vphatfla/naplex-go/backend/internal/users"
)

type Module struct {
	h *users.Handler
	s *users.Service
}

func NewModule(q *database.Queries) *Module {
	s := NewService(q)
	h := NewHandler(s)

	return &Module{
		h: h,
		s: s,
	}
}
