-- name: GetAddressByID :one
SELECT id, subnet_id, ip, ip_text, hostname, mac, description, status, tag_id, is_gateway, device_id, device_port, owner, last_seen_at, dns_checked_at, created_at, updated_at
FROM addresses
WHERE id = ?;

-- name: ListAddresses :many
SELECT id, subnet_id, ip, ip_text, hostname, mac, description, status, tag_id, is_gateway, device_id, device_port, owner, last_seen_at, dns_checked_at, created_at, updated_at
FROM addresses
WHERE subnet_id = ?
ORDER BY ip
LIMIT ? OFFSET ?;

-- name: CountAddresses :one
SELECT COUNT(*) FROM addresses WHERE subnet_id = ?;

-- name: CountAddressesByStatus :one
SELECT COUNT(*) FROM addresses WHERE subnet_id = ? AND status = ?;

-- name: CreateAddress :execresult
INSERT INTO addresses (subnet_id, ip, ip_text, hostname, mac, description, status, tag_id, is_gateway, device_id, device_port, owner)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateAddress :exec
UPDATE addresses
SET hostname = ?, mac = ?, description = ?, status = ?, tag_id = ?, is_gateway = ?, device_id = ?, device_port = ?, owner = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteAddress :exec
DELETE FROM addresses WHERE id = ?;

-- name: GetAddressByIP :one
SELECT id, subnet_id, ip, ip_text, hostname, mac, description, status, tag_id, is_gateway, device_id, device_port, owner, last_seen_at, dns_checked_at, created_at, updated_at
FROM addresses
WHERE subnet_id = ? AND ip = ?;

-- name: GetAddressByIPText :one
SELECT id, subnet_id, ip, ip_text, hostname, mac, description, status, tag_id, is_gateway, device_id, device_port, owner, last_seen_at, dns_checked_at, created_at, updated_at
FROM addresses
WHERE subnet_id = ? AND ip_text = ?;

-- name: ListAddressIPsBySubnet :many
SELECT ip_text FROM addresses WHERE subnet_id = ?;


