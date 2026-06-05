-- name: GetLocationByID :one
SELECT id, name, address, room, row_name, description, created_at, updated_at FROM locations WHERE id = ?;
-- name: ListLocations :many
SELECT id, name, address, room, row_name, description, created_at, updated_at FROM locations ORDER BY name LIMIT ? OFFSET ?;
-- name: CountLocations :one
SELECT COUNT(*) FROM locations;
-- name: CreateLocation :execresult
INSERT INTO locations (name, address, room, row_name, description) VALUES (?, ?, ?, ?, ?);
-- name: UpdateLocation :exec
UPDATE locations SET name = ?, address = ?, room = ?, row_name = ?, description = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;
-- name: DeleteLocation :exec
DELETE FROM locations WHERE id = ?;
-- name: GetRackByID :one
SELECT id, location_id, name, size_u, row_name, position, direction, description, created_at, updated_at FROM racks WHERE id = ?;
-- name: ListRacks :many
SELECT id, location_id, name, size_u, row_name, position, direction, description, created_at, updated_at FROM racks WHERE location_id = ? ORDER BY name LIMIT ? OFFSET ?;
-- name: ListAllRacks :many
SELECT id, location_id, name, size_u, row_name, position, direction, description, created_at, updated_at FROM racks ORDER BY name LIMIT ? OFFSET ?;
-- name: CountRacks :one
SELECT COUNT(*) FROM racks WHERE location_id = ?;
-- name: CountAllRacks :one
SELECT COUNT(*) FROM racks;
-- name: CreateRack :execresult
INSERT INTO racks (location_id, name, size_u, row_name, position, direction, description) VALUES (?, ?, ?, ?, ?, ?, ?);
-- name: UpdateRack :exec
UPDATE racks SET name = ?, size_u = ?, row_name = ?, position = ?, direction = ?, description = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;
-- name: DeleteRack :exec
DELETE FROM racks WHERE id = ?;
