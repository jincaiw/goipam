-- name: GetVrfByID :one
SELECT id, name, rd, description, created_at, updated_at FROM vrfs WHERE id = ?;
-- name: ListVrfs :many
SELECT id, name, rd, description, created_at, updated_at FROM vrfs ORDER BY name LIMIT ? OFFSET ?;
-- name: CountVrfs :one
SELECT COUNT(*) FROM vrfs;
-- name: CreateVrf :execresult
INSERT INTO vrfs (name, rd, description) VALUES (?, ?, ?);
-- name: UpdateVrf :exec
UPDATE vrfs SET name = ?, rd = ?, description = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;
-- name: DeleteVrf :exec
DELETE FROM vrfs WHERE id = ?;
