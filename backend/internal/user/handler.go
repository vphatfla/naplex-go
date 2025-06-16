package user

import (
	"log"
	"net/http"

	"github.com/vphatfla/naplex-go/backend/internal/utils"
)

type Handler struct {
	s Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		s: s,
	}
}
func (h *Handler) RegisterRoutes() *http.ServeMux {
	m := http.NewServeMux()

	m.Handle("/info", http.HandlerFunc(h.HandleGetUser))

	return m
}
func (h *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("user id = %v", r.Context().Value("user_id"))
	log.Printf("user session = %v", r.Cookies())
	log.Printf("r context = %v", r.Context())
	utils.HTTPJsonResponse(w, nil)
}
