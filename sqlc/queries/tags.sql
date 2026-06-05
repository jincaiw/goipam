-- name: GetTagByID :one
SELECT id, name, color, created_at, updated_at FROM tags WHERE id = ?;

-- name: ListTags :many
SELECT id, name, color, created_at, updated_at FROM tags ORDER BY name;

-- name: CreateTag :execresult
INSERT INTO tags (name, color) VALUES (?, ?);

-- name: UpdateTag :exec
UPDATE tags SET name = ?, color = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;

-- name: DeleteTag :exec
DELETE FROM tags WHERE id = ?;
