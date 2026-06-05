-- name: GetDeviceByID :one
SELECT id, name, hostname, mgmt_ip, type_id, location_id, rack_id, rack_start_u, rack_size_u, manufacturer, model, serial_number, description, created_at, updated_at FROM devices WHERE id = ?;
-- name: ListDevices :many
SELECT id, name, hostname, mgmt_ip, type_id, location_id, rack_id, rack_start_u, rack_size_u, manufacturer, model, serial_number, description, created_at, updated_at FROM devices ORDER BY name LIMIT ? OFFSET ?;
-- name: CountDevices :one
SELECT COUNT(*) FROM devices;
-- name: CreateDevice :execresult
INSERT INTO devices (name, hostname, mgmt_ip, type_id, location_id, rack_id, rack_start_u, rack_size_u, manufacturer, model, serial_number, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
-- name: UpdateDevice :exec
UPDATE devices SET name = ?, hostname = ?, mgmt_ip = ?, type_id = ?, location_id = ?, rack_id = ?, rack_start_u = ?, rack_size_u = ?, manufacturer = ?, model = ?, serial_number = ?, description = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;
-- name: DeleteDevice :exec
DELETE FROM devices WHERE id = ?;
-- name: GetDeviceTypeByID :one
SELECT id, name, icon, description, created_at, updated_at FROM device_types WHERE id = ?;
-- name: ListDeviceTypes :many
SELECT id, name, icon, description, created_at, updated_at FROM device_types ORDER BY name;
-- name: CreateDeviceType :execresult
INSERT INTO device_types (name, icon, description) VALUES (?, ?, ?);
-- name: UpdateDeviceType :exec
UPDATE device_types SET name = ?, icon = ?, description = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;
-- name: DeleteDeviceType :exec
DELETE FROM device_types WHERE id = ?;
