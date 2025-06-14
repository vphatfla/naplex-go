package users

import (
	"log"
	"net/http"

	"github.com/vphatfla/naplex-go/backend/internal/utils"
)

type UserHandler struct {
}

func (h *UserHandler) RegisterRoutes() *http.ServeMux {
	m := http.NewServeMux()

	m.Handle("/info", http.HandlerFunc(h.HandleGetUserInfo))

	return m
}
func (h *UserHandler) HandleGetUserInfo(w http.ResponseWriter, r *http.Request) {
	log.Printf("user id = %v", r.Context().Value("user_id"))
	log.Printf("user session = %v", r.Cookies())
	log.Printf("r context = %v", r.Context())
	utils.HTTPJsonResponse(w, nil)
}
