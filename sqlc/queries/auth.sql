-- name: CreateAPIToken :execresult
INSERT INTO api_tokens (name, token_hash, user_id, permissions, expires_at, status)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetAPITokenByHash :one
SELECT id, name, token_hash, user_id, permissions, expires_at, last_used_at, status, created_at, updated_at
FROM api_tokens
WHERE token_hash = ?;

-- name: ListAPITokens :many
SELECT id, name, user_id, permissions, expires_at, last_used_at, status, created_at, updated_at
FROM api_tokens
WHERE user_id = ?
ORDER BY id;

-- name: UpdateAPITokenStatus :exec
UPDATE api_tokens
SET status = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateAPITokenLastUsed :exec
UPDATE api_tokens
SET last_used_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteAPIToken :exec
DELETE FROM api_tokens WHERE id = ?;

-- name: GetResourcePermissions :many
SELECT id, user_id, group_id, resource_type, resource_id, permission_level, created_at, updated_at
FROM resource_permissions
WHERE resource_type = ? AND resource_id = ?;

-- name: GetUserResourcePermissions :many
SELECT id, user_id, group_id, resource_type, resource_id, permission_level, created_at, updated_at
FROM resource_permissions
WHERE user_id = ?;

-- name: SetResourcePermission :execresult
INSERT INTO resource_permissions (user_id, group_id, resource_type, resource_id, permission_level)
VALUES (?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE permission_level = VALUES(permission_level), updated_at = CURRENT_TIMESTAMP;

-- name: DeleteResourcePermission :exec
DELETE FROM resource_permissions WHERE id = ?;
