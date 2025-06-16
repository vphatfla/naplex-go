package auth

import (
	"github.com/vphatfla/naplex-go/backend/internal/config"
	"github.com/vphatfla/naplex-go/backend/internal/shared/database"
)

type Module struct {
	Handler *Handler
	Service *Service
}

func NewModule(config *config.Config, q *database.Queries) *Module {
	s := NewService(q)
	h := NewHandler(config, s)

	return &Module{
		Handler: h,
		Service: s,
	}
}
