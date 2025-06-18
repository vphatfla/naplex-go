package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/vphatfla/naplex-go/backend/internal/shared/assert"
	"github.com/vphatfla/naplex-go/backend/internal/shared/database"
	"github.com/vphatfla/naplex-go/backend/internal/utils"
)

type Handler struct {
	s *Service
}

func NewHandler(s *Service) *Handler {
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
	id, err := assert.AssertAndGetValue[int32](r.Context().Value("user_id"))
	if err != nil {
		utils.HTTPJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.s.GetUserProfile(context.Background(), id)
	if err != nil {
		utils.HTTPJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uRes := &User{
		ID: u.ID,
		Name: u.Name,
		FirstName: u.FirstName.String,
		LastName: u.LastName.String,
		Email: u.Email,
		Picture: u.Picture.String,
	}
	utils.HTTPJsonResponse(w, uRes)
	return
}

func (h *Handler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := assert.AssertAndGetValue[int32](r.Context().Value("user_id"))
	if err != nil {
		utils.HTTPJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var u User
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		utils.HTTPJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if u.ID != id {
		utils.HTTPJsonError(w, "user ID do not match, unauthorized", http.StatusNonAuthoritativeInfo)
		return
	}

	uInput := &database.User{
		ID: u.ID,
		GoogleID: "", // empty, update query do not allow/use google_id
		FirstName: pgtype.Text{ String: u.FirstName, Valid: true },
		LastName: pgtype.Text{ String: u.LastName, Valid: true },
		Name: u.Name,
		Picture: pgtype.Text{ String: u.Picture, Valid: true },
	}

	uRes, err := h.s.UpdateUserProfile(context.Background(), uInput)
	if err != nil {
		utils.HTTPJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.HTTPJsonResponse(w, uRes)
}
