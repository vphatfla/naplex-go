-- name: CountRawQuestion :one
SELECT COUNT(id) FROM raw_questions;

-- name: GetRawQuestionByID :one
SELECT * FROM raw_questions
WHERE id=$1 LIMIT 1;

-- name: GetRawQuestionWithRange :many
SELECT * FROM raw_questions
WHERE id >= $1 AND id <= $2;
-- name: InsertRawQuestion :one
INSERT INTO raw_questions (title, raw_question, link)
VALUES ($1, $2, $3)
RETURNING id;

-- name: UpdateRawQuestion :one
UPDATE raw_questions
SET title = $1, raw_question = $2, link = $3
WHERE id = $4
RETURNING id;

-- name: InsertProcessedQuestion :one
INSERT INTO processed_questions (title, question, multiple_choices, correct_answer, explanation, keywords)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: GetProcessedQuestionByID :one
SELECT * FROM processed_questions WHERE id = $1;

-- name: UpdateProcessedQuestion :one
UPDATE processed_questions
SET title = $1, question = $2, multiple_choices = $3, correct_answer = $4, explanation = $5, keywords = $6
WHERE id = $7
RETURNING id;


