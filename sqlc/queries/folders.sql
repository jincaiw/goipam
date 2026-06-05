-- name: GetFolderByID :one
SELECT id, section_id, parent_id, name, description, sort_order, created_at, updated_at
FROM folders
WHERE id = ?;

-- name: ListFolders :many
SELECT id, section_id, parent_id, name, description, sort_order, created_at, updated_at
FROM folders
WHERE section_id = ?
ORDER BY sort_order, name
LIMIT ? OFFSET ?;

-- name: CountFolders :one
SELECT COUNT(*) FROM folders WHERE section_id = ?;

-- name: CreateFolder :execresult
INSERT INTO folders (section_id, parent_id, name, description, sort_order)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateFolder :exec
UPDATE folders
SET name = ?, description = ?, sort_order = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteFolder :exec
DELETE FROM folders WHERE id = ?;

-- name: GetSubFolders :many
SELECT id, section_id, parent_id, name, description, sort_order, created_at, updated_at
FROM folders
WHERE parent_id = ?
ORDER BY sort_order, name;
