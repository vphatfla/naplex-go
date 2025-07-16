-- name: GetAllIds :many
SELECT id FROM processed_questions;

-- name: GetProcessedQuestionsInBatch :many
SELECT id, title, question, multiple_choices, correct_answer, explanation, keywords, link
FROM processed_questions
WHERE id = ANY($1::int[]) AND correct_answer <> '';
