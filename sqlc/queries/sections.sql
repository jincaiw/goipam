-- name: GetSectionByID :one
SELECT id, name, description, parent_id, strict_mode, show_vlan, show_vrf, subnet_order, sort_order, created_at, updated_at
FROM sections
WHERE id = ?;

-- name: ListSections :many
SELECT id, name, description, parent_id, strict_mode, show_vlan, show_vrf, subnet_order, sort_order, created_at, updated_at
FROM sections
ORDER BY sort_order, name
LIMIT ? OFFSET ?;

-- name: CountSections :one
SELECT COUNT(*) FROM sections;

-- name: CreateSection :execresult
INSERT INTO sections (name, description, parent_id, strict_mode, show_vlan, show_vrf, subnet_order, sort_order)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateSection :exec
UPDATE sections
SET name = ?, description = ?, parent_id = ?, strict_mode = ?, show_vlan = ?, show_vrf = ?, subnet_order = ?, sort_order = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteSection :exec
DELETE FROM sections WHERE id = ?;

-- name: GetSubSections :many
SELECT id, name, description, parent_id, strict_mode, show_vlan, show_vrf, subnet_order, sort_order, created_at, updated_at
FROM sections
WHERE parent_id = ?
ORDER BY sort_order, name;
