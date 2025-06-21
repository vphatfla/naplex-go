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

	// UserQuestion record may not exist since the user never attempt the given question
	// therefore, the skeleton UserQuestion is needed
	if err == pgx.ErrNoRows {
		uq = database.UsersQuestion{
			Attempts: pgtype.Int4{Int32: 0, Valid: true},
			Saved: pgtype.Bool{Bool: false, Valid: true},
			Hidden: pgtype.Bool{Bool: false, Valid: true},
			Status: database.NullQuestionStatus{QuestionStatus: database.QuestionStatusNA},
		}
	}

	qDTO := GenerateDTO(&q, &uq)
	return qDTO, nil
}

func (s *Service) CreateOrUpdateUserQuestion(ctx context.Context, uid int32, qDTO *QuestionDTO) (*QuestionDTO, error) {
	params := &database.CreateOrUpdateUserQuestionParams{
		Column1: qDTO.Status,
		Attempts: pgtype.Int4{Int32: qDTO.Attempts, Valid: true},
		Saved: pgtype.Bool{Bool: qDTO.Saved, Valid: true},
		Hidden: pgtype.Bool{Bool: qDTO.Hidden, Valid: true},
		Uid: uid,
		Qid: qDTO.ID,
	}
	uq, err := s.q.CreateOrUpdateUserQuestion(ctx, *params)
	if err != nil {
		return nil, err
	}
	qDTO.Status = uq.Status.QuestionStatus
	qDTO.Attempts = uq.Attempts.Int32
	qDTO.Saved = uq.Saved.Bool
	qDTO.Hidden = uq.Hidden.Bool

	return qDTO, nil
}

func (s *Service) GetAllPassedQuestion(ctx context.Context, uid int32) ([]QuestionDTO, error) {
	temp, err := s.q.GetAllPassedQuestion(ctx, uid)
	if err != nil {
		return nil, err
	}
	var list []QuestionDTO

	for _,item := range temp {
		qDTO := &QuestionDTO{
			ID: item.Qid,
			Title: item.Title,
			Question: item.Question,
			Multiple_choices: strings.Split(item.MultipleChoices, "\n"),
			Correct_answer: item.CorrectAnswer,
			Explanation: item.Explanation.String,
			Keywords: strings.Split(item.Keywords.String, ","),
			Link: item.Link.String,
			Status: item.Status.QuestionStatus,
			Attempts: item.Attempts.Int32,
			Saved: item.Saved.Bool,
			Hidden: item.Hidden.Bool,
		}
		list = append(list, *qDTO)
	}

	return list, nil
}
