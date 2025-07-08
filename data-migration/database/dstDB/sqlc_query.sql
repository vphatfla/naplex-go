-- name: CreateQuestion :one
INSERT INTO questions (title, question, multiple_choices, correct_answer, explanation, keywords, link)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;
