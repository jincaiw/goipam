-- name: GetUserByID :one
SELECT id, username, password_hash, display_name, email, status, failed_login_count, locked_until, last_login_at, created_at, updated_at
FROM users
WHERE id = ?;

-- name: GetUserByUsername :one
SELECT id, username, password_hash, display_name, email, status, failed_login_count, locked_until, last_login_at, created_at, updated_at
FROM users
WHERE username = ?;

-- name: ListUsers :many
SELECT id, username, display_name, email, status, last_login_at, created_at, updated_at
FROM users
ORDER BY id
LIMIT ? OFFSET ?;

-- name: CountUsers :one
SELECT COUNT(*) FROM users;

-- name: CreateUser :execresult
INSERT INTO users (username, password_hash, display_name, email, status)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateUser :exec
UPDATE users
SET display_name = ?, email = ?, status = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateUserPassword :exec
UPDATE users
SET password_hash = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateUserLoginInfo :exec
UPDATE users
SET failed_login_count = ?, locked_until = ?, last_login_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
