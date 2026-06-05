-- name: GetSubnetByID :one
SELECT id, section_id, parent_id, cidr, ip_version, network, prefix_len, name, description, vlan_id, vrf_id, location_id, nameserver_id, scan_enabled, discover_enabled, threshold_percent, allow_request, status, created_at, updated_at
FROM subnets
WHERE id = ?;

-- name: ListSubnets :many
SELECT id, section_id, parent_id, cidr, ip_version, network, prefix_len, name, description, vlan_id, vrf_id, location_id, nameserver_id, scan_enabled, discover_enabled, threshold_percent, allow_request, status, created_at, updated_at
FROM subnets
WHERE section_id = ?
ORDER BY network
LIMIT ? OFFSET ?;

-- name: CountSubnets :one
SELECT COUNT(*) FROM subnets WHERE section_id = ?;

-- name: CreateSubnet :execresult
INSERT INTO subnets (section_id, parent_id, cidr, ip_version, network, prefix_len, name, description, vlan_id, vrf_id, location_id, nameserver_id, scan_enabled, discover_enabled, threshold_percent, allow_request, status)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateSubnet :exec
UPDATE subnets
SET name = ?, description = ?, vlan_id = ?, vrf_id = ?, location_id = ?, nameserver_id = ?, scan_enabled = ?, discover_enabled = ?, threshold_percent = ?, allow_request = ?, status = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteSubnet :exec
DELETE FROM subnets WHERE id = ?;

-- name: GetChildSubnets :many
SELECT id, section_id, parent_id, cidr, ip_version, network, prefix_len, name, description, vlan_id, vrf_id, location_id, nameserver_id, scan_enabled, discover_enabled, threshold_percent, allow_request, status, created_at, updated_at
FROM subnets
WHERE parent_id = ?
ORDER BY network;

-- name: FindOverlappingSubnets :many
SELECT id, section_id, parent_id, cidr, ip_version, network, prefix_len, name, description, vlan_id, vrf_id, location_id, nameserver_id, scan_enabled, discover_enabled, threshold_percent, allow_request, status, created_at, updated_at
FROM subnets
WHERE section_id = ? AND (vrf_id = ? OR (? IS NULL AND vrf_id IS NULL)) AND id != ?;
