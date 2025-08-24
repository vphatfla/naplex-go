package dataTransfer

import (
	"context"
	"encoding/json"
	"net/http"

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

func (h *Handler) HandleInsertQuestionsBatch(w http.ResponseWriter, r *http.Request) {
	var questions []Question
	err := json.NewDecoder(r.Body).Decode(&questions)
	if err != nil {
		utils.HTTPJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	count, err := h.s.InsertBulkQuestions(context.Background(), questions)
	if err != nil {
		utils.HTTPJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.HTTPJsonResponse(w, map[string]int{"count": int(count)})
}
