-- ============================================
-- User Creation and Authentication
-- ============================================

-- name: CreateUser :one
INSERT INTO users (
    google_id,
    email,
    name,
    first_name,
    last_name,
    picture,
    last_login_at
) VALUES (
    $1, $2, $3, $4, $5, $6, NOW()
)
RETURNING *;

-- name: UpsertUser :one
-- Used for Google OAuth login - creates user if not exists, updates if exists
INSERT INTO users (
    google_id,
    email,
    name,
    first_name,
    last_name,
    picture,
    last_login_at
) VALUES (
    $1, $2, $3, $4, $5, $6, NOW()
)
ON CONFLICT (google_id)
DO UPDATE SET
    email = EXCLUDED.email,
    name = EXCLUDED.name,
    first_name = EXCLUDED.first_name,
    last_name = EXCLUDED.last_name,
    picture = EXCLUDED.picture,
    last_login_at = NOW()
RETURNING *;

-- ============================================
-- User Retrieval
-- ============================================

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByGoogleID :one
SELECT * FROM users
WHERE google_id = $1;

-- name: GetUsersByIDs :many
SELECT * FROM users
WHERE id = ANY($1::int[])
ORDER BY created_at DESC;

-- ============================================
-- User Existence Checks
-- ============================================

-- name: CheckUserExistsByEmail :one
SELECT EXISTS(
    SELECT 1 FROM users WHERE email = $1
);

-- name: CheckUserExistsByGoogleID :one
SELECT EXISTS(
    SELECT 1 FROM users WHERE google_id = $1
);

-- ============================================
-- User Updates
-- ============================================

-- name: UpdateUserProfile :one
UPDATE users
SET
    name = COALESCE($2, name),
    first_name = COALESCE($3, first_name),
    last_name = COALESCE($4, last_name),
    picture = COALESCE($5, picture)
WHERE id = $1
RETURNING *;

-- name: UpdateUserEmail :one
UPDATE users
SET email = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUserLastLogin :exec
UPDATE users
SET last_login_at = NOW()
WHERE id = $1;

-- name: UpdateUserPicture :one
UPDATE users
SET picture = $2
WHERE id = $1
RETURNING *;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CountUsers :one
SELECT COUNT(*) FROM users;

-- name: SearchUsersByName :many
SELECT * FROM users
WHERE
    name ILIKE '%' || $1 || '%' OR
    first_name ILIKE '%' || $1 || '%' OR
    last_name ILIKE '%' || $1 || '%'
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: SearchUsersByEmail :many
SELECT * FROM users
WHERE email ILIKE '%' || $1 || '%'
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListUsersByCreatedDate :many
SELECT * FROM users
WHERE created_at >= $1 AND created_at <= $2
ORDER BY created_at DESC;

-- name: GetDailyNewUsers :many
SELECT
    DATE(created_at) as date,
    COUNT(*) as new_users
FROM users
WHERE created_at >= $1
GROUP BY DATE(created_at)
ORDER BY date DESC;

-- ============================================
-- User Deletion
-- ============================================

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: DeleteUserByEmail :exec
DELETE FROM users
WHERE email = $1;

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
INSERT INTO processed_questions (title, question, multiple_choices, correct_answer, explanation, keywords, link)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;

-- name: GetProcessedQuestionByID :one
SELECT * FROM processed_questions WHERE id = $1;

-- name: UpdateProcessedQuestion :one
UPDATE processed_questions
SET title = $1, question = $2, multiple_choices = $3, correct_answer = $4, explanation = $5, keywords = $6, link = $7
WHERE id = $8
RETURNING id;
