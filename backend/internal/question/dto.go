package question

import (
	"strings"

	"github.com/vphatfla/naplex-go/backend/internal/shared/database"
)

// DTO stands for data transfer object
// DTO use to define a schema JSON structure for client's HTTP Response
type QuestionDTO struct {
	ID               int32    `json:"question_id"`
	Title            string   `json:"title"`
	Question         string   `json:"question"`
	Multiple_choices []string `json:"multiple_choices"`
	Correct_answer   string   `json:"correct_answer"`
	Explanation      string   `json:"explanation"`
	Keywords         []string `json:"keywords"`
	Link             string   `json:"link"`
	Status           database.QuestionStatus   `json:"status"`
	Attempts         int32    `json:"attempts"`
	Saved            bool     `json:"saved"`
	Hidden           bool     `json:"hidden"`
}

func GenerateDTO(q *database.Question, uq *database.UsersQuestion) *QuestionDTO {
	questionDTO := &QuestionDTO{
		ID:               q.ID,
		Title:            q.Title,
		Question:         q.Question,
		Multiple_choices: strings.Split(q.MultipleChoices, "\n"),
		Correct_answer:   q.CorrectAnswer,
		Explanation:      q.Explanation.String,
		Keywords:         strings.Split(q.Keywords.String, ","),
		Link:             q.Link.String,
		Status:           uq.Status.QuestionStatus,
		Attempts:         uq.Attempts.Int32,
		Saved:            uq.Saved.Bool,
		Hidden:           uq.Hidden.Bool,
	}
	return questionDTO
}
