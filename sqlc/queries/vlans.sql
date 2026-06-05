-- name: GetVlanByID :one
SELECT id, l2_domain_id, vlan_id, name, description, created_at, updated_at FROM vlans WHERE id = ?;
-- name: ListVlans :many
SELECT id, l2_domain_id, vlan_id, name, description, created_at, updated_at FROM vlans ORDER BY vlan_id LIMIT ? OFFSET ?;
-- name: CountVlans :one
SELECT COUNT(*) FROM vlans;
-- name: CreateVlan :execresult
INSERT INTO vlans (l2_domain_id, vlan_id, name, description) VALUES (?, ?, ?, ?);
-- name: UpdateVlan :exec
UPDATE vlans SET l2_domain_id = ?, vlan_id = ?, name = ?, description = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;
-- name: DeleteVlan :exec
DELETE FROM vlans WHERE id = ?;
-- name: GetL2DomainByID :one
SELECT id, name, description, created_at, updated_at FROM l2_domains WHERE id = ?;
-- name: ListL2Domains :many
SELECT id, name, description, created_at, updated_at FROM l2_domains ORDER BY name LIMIT ? OFFSET ?;
-- name: CountL2Domains :one
SELECT COUNT(*) FROM l2_domains;
-- name: CreateL2Domain :execresult
INSERT INTO l2_domains (name, description) VALUES (?, ?);
-- name: UpdateL2Domain :exec
UPDATE l2_domains SET name = ?, description = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;
-- name: DeleteL2Domain :exec
DELETE FROM l2_domains WHERE id = ?;
