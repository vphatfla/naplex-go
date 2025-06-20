package question

import (
	"context"
	"net/http"
	"strconv"

	"github.com/vphatfla/naplex-go/backend/internal/shared/assert"
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

func (h *Handler) HandleGetQuestion(w http.ResponseWriter, r *http.Request) {
	uid, err := assert.AssertAndGetValue[int32](r.Context().Value("user_id"))
	if err != nil {
		utils.HTTPJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	qidStr := r.URL.Query().Get("question_id")
	if qidStr == "" {
		utils.HTTPJsonError(w, "questions_id is required", http.StatusBadRequest)
		return
	}
	qid, err := strconv.Atoi(qidStr)
	if err != nil {
		utils.HTTPJsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	qRes, err := h.s.GetQuestion(context.Background(), uid, int32(qid))
	if err != nil {
		utils.HTTPJsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.HTTPJsonResponse(w, qRes)
}
