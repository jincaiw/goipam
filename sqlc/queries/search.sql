-- name: SearchAddresses :many
SELECT id, subnet_id, ip, ip_text, hostname, mac, description, status, tag_id, is_gateway, device_id, device_port, owner, last_seen_at, dns_checked_at, created_at, updated_at FROM addresses WHERE hostname LIKE ? OR ip_text LIKE ? OR mac LIKE ? OR description LIKE ? LIMIT ? OFFSET ?;

-- name: SearchSubnets :many
SELECT id, section_id, parent_id, cidr, ip_version, network, prefix_len, name, description, vlan_id, vrf_id, location_id, nameserver_id, scan_enabled, discover_enabled, threshold_percent, allow_request, status, created_at, updated_at FROM subnets WHERE name LIKE ? OR cidr LIKE ? OR description LIKE ? LIMIT ?;

-- name: SearchDevices :many
SELECT id, name, hostname, mgmt_ip, type_id, location_id, rack_id, rack_start_u, rack_size_u, manufacturer, model, serial_number, description, created_at, updated_at FROM devices WHERE name LIKE ? OR hostname LIKE ? OR mgmt_ip LIKE ? LIMIT ?;
