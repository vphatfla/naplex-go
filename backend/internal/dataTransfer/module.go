package dataTransfer

import (
	"github.com/vphatfla/naplex-go/backend/internal/shared/database"
)

type Module struct {
	Handler *Handler
	Service *Service
}

func NewModule(q *database.Queries) *Module {
	s := NewService(q)
	h := NewHandler(s)

	return &Module{
		Handler: h,
		Service: s,
	}
}
