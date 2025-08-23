package dataTransfer

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/vphatfla/naplex-go/backend/internal/shared/database"
)

type Service struct {
	q *database.Queries
}

func NewService(q *database.Queries) *Service {
	return &Service{
		q: q,
	}
}

func (s *Service) InsertBulkQuestions(ctx context.Context, questions []Question) (int64, error) {
	var params []database.CreateQuestionsBatchParams
	for _, q := range questions {
		p := database.CreateQuestionsBatchParams{
			Title:           q.Title,
			Question:        q.Question,
			MultipleChoices: q.MultipleChoices,
			CorrectAnswer:   q.CorrectAnswer,
			Explanation:     pgtype.Text{String: q.Explanation, Valid: true},
			Keywords:        pgtype.Text{String: q.Keywords, Valid: true},
			Link:            pgtype.Text{String: q.Link, Valid: true},
		}
		params = append(params, p)
	}

	count, err := s.q.CreateQuestionsBatch(ctx, params)
	if err != nil {
		return -1, err
	}

	return count, nil
}
