-- name: GetProcessedQuestion :one
SELECT id, title, question, multiple_choices, correct_answer, explanation, keywords
FROM processed_questions;

-- name: GetBatchProcessedQuestions :many
SELECT id, title, question, multiple_choices, correct_answer, explanation, keywords
FROM processed_questions
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: GetAllIds :many
SELECT id FROM processed_questions;

-- name: GetProcessedQuestionsInBatch :many
SELECT id, title, question, multiple_choices, correct_answer, explanation, keywords
FROM processed_questions
WHERE id = ANY($1::int[]);
