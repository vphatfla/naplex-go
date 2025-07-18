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

-- name: CreateOrUpsertUser :one
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

-- name: InsertQuestion :one
INSERT INTO questions (title, question, multiple_choices, correct_answer, explanation, keywords, link)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;

-- name: GetQuestionByID :one
SELECT * FROM questions WHERE id = $1;

-- name: UpdateQuestion :one
UPDATE questions
SET title = $1, question = $2, multiple_choices = $3, correct_answer = $4, explanation = $5, keywords = $6, link = $7
WHERE id = $8
RETURNING id;

-- name: GetUserQuestion :one
SELECT * FROM users_questions
WHERE uid = $1 AND qid = $2;

-- name: CreateOrUpdateUserQuestion :one
INSERT INTO users_questions (uid, qid, status, attempts, saved, hidden)
VALUES ($5, $6, $1::question_status, $2, $3, $4)
ON CONFLICT (uid, qid)
DO UPDATE SET
    status = COALESCE(EXCLUDED.status, users_questions.status),
    attempts = COALESCE(EXCLUDED.attempts, users_questions.attempts),
    saved = COALESCE(EXCLUDED.saved, users_questions.saved),
    hidden = COALESCE(EXCLUDED.hidden, users_questions.hidden),
    updated_at = NOW()
RETURNING *;

-- name: GetAllPassedQuestion :many
SELECT
    uq.*,
    q.*
FROM users_questions uq
JOIN questions q ON q.id = uq.qid
WHERE uq.uid = $1 AND uq.status = 'PASSED'::question_status;

-- name: GetAllFailedQuestion :many
SELECT
    uq.*,
    q.*
FROM users_questions uq
JOIN questions q ON q.id = uq.qid
WHERE uq.uid = $1 AND uq.status = 'FAILED'::question_status;

-- name: GetRandomDailyQuestions :many
SELECT
    q.*,
    COALESCE(q.id, uq.qid) AS qid,
    COALESCE(uq.status, 'NA'::question_status) AS status,
    COALESCE(uq.attempts, 0) AS attempts,
    COALESCE(uq.saved, FALSE) AS saved,
    COALESCE(uq.hidden, FALSE) AS hidden
FROM
    questions q
LEFT JOIN
    users_questions uq ON uq.qid = q.id AND uq.uid = $1
WHERE
    uq.uid IS NULL -- Junction record do not exists
    OR (
        uq.status IN ('FAILED'::question_status, 'NA'::question_status)
        AND
        uq.hidden = FALSE
    )
ORDER BY RANDOM() --randomly order selection
LIMIT $2;
