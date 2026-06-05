-- name: GetNatByID :one
SELECT id, name, type, inside_address_id, outside_address_id, inside_ip, outside_ip, inside_port, outside_port, protocol, device_id, description, enabled, created_at, updated_at FROM nats WHERE id = ?;
-- name: ListNats :many
SELECT id, name, type, inside_address_id, outside_address_id, inside_ip, outside_ip, inside_port, outside_port, protocol, device_id, description, enabled, created_at, updated_at FROM nats ORDER BY name LIMIT ? OFFSET ?;
-- name: CountNats :one
SELECT COUNT(*) FROM nats;
-- name: CreateNat :execresult
INSERT INTO nats (name, type, inside_address_id, outside_address_id, inside_ip, outside_ip, inside_port, outside_port, protocol, device_id, description, enabled) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
-- name: UpdateNat :exec
UPDATE nats SET name = ?, type = ?, inside_address_id = ?, outside_address_id = ?, inside_ip = ?, outside_ip = ?, inside_port = ?, outside_port = ?, protocol = ?, device_id = ?, description = ?, enabled = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;
-- name: DeleteNat :exec
DELETE FROM nats WHERE id = ?;
