-- name: GetNameserverByID :one
SELECT id, name, type, server, port, timeout_ms, enabled, description, created_at, updated_at FROM nameservers WHERE id = ?;

-- name: ListNameservers :many
SELECT id, name, type, server, port, timeout_ms, enabled, description, created_at, updated_at FROM nameservers ORDER BY name LIMIT ? OFFSET ?;

-- name: CountNameservers :one
SELECT COUNT(*) FROM nameservers;

-- name: CreateNameserver :execresult
INSERT INTO nameservers (name, type, server, port, timeout_ms, enabled, description) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateNameserver :exec
UPDATE nameservers SET name = ?, type = ?, server = ?, port = ?, timeout_ms = ?, enabled = ?, description = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;

-- name: DeleteNameserver :exec
DELETE FROM nameservers WHERE id = ?;
