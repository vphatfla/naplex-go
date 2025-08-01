// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: sqlc_query.sql

package srcDB

import (
	"context"
)

const getAllIds = `-- name: GetAllIds :many
SELECT id FROM processed_questions
`

func (q *Queries) GetAllIds(ctx context.Context) ([]int32, error) {
	rows, err := q.db.Query(ctx, getAllIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var id int32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProcessedQuestionsInBatch = `-- name: GetProcessedQuestionsInBatch :many
SELECT id, title, question, multiple_choices, correct_answer, explanation, keywords, link
FROM processed_questions
WHERE id = ANY($1::int[]) AND correct_answer <> ''
`

func (q *Queries) GetProcessedQuestionsInBatch(ctx context.Context, dollar_1 []int32) ([]ProcessedQuestion, error) {
	rows, err := q.db.Query(ctx, getProcessedQuestionsInBatch, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProcessedQuestion
	for rows.Next() {
		var i ProcessedQuestion
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Question,
			&i.MultipleChoices,
			&i.CorrectAnswer,
			&i.Explanation,
			&i.Keywords,
			&i.Link,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
