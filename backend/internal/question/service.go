package question

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/vphatfla/naplex-go/backend/internal/shared/database"
)

type Service struct {
	q *database.Queries
}

func NewService(queries *database.Queries) *Service {
	return &Service{
		q: queries,
	}
}

func (s *Service) GetQuestion(ctx context.Context, uid int32, qid int32) (*QuestionDTO, error) {
	q, err := s.q.GetQuestionByID(ctx, qid)
	if err != nil {
		return nil, err
	}

	params := &database.GetUserQuestionParams{
		Uid: uid,
		Qid: qid,
	}

	uq, err := s.q.GetUserQuestion(ctx, *params)
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}
	if err == pgx.ErrNoRows {
		uq = database.UsersQuestion{
			Attempts: pgtype.Int4{Int32: 0, Valid: true},
			Saved: pgtype.Bool{Bool: false, Valid: true},
			Hidden: pgtype.Bool{Bool: false, Valid: true},
		}
	}
	questionDTO := &QuestionDTO{
		ID:               q.ID,
		Title:            q.Title,
		Question:         q.Question,
		Multiple_choices: strings.Split(q.MultipleChoices, ""),
		Correct_answer:   q.CorrectAnswer,
		Explanation:      q.Explanation.String,
		Keywords:         strings.Split(q.Keywords.String, ""),
		Link:             q.Link.String,
		Status:           string(uq.Status.QuestionStatus),
		Attempts:         uq.Attempts.Int32,
		Saved:            uq.Saved.Bool,
		Hidden:           uq.Hidden.Bool,
	}

	return questionDTO, nil
}
