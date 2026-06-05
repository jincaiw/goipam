-- +goose Up
CREATE TABLE IF NOT EXISTS sections (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(128) NOT NULL UNIQUE,
    description TEXT,
    parent_id BIGINT NULL,
    strict_mode TINYINT(1) NOT NULL DEFAULT 0,
    show_vlan TINYINT(1) NOT NULL DEFAULT 1,
    show_vrf TINYINT(1) NOT NULL DEFAULT 1,
    subnet_order VARCHAR(32) NOT NULL DEFAULT 'subnet',
    sort_order INT NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES sections(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS folders (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    section_id BIGINT NOT NULL,
    parent_id BIGINT NULL,
    name VARCHAR(128) NOT NULL,
    description TEXT,
    sort_order INT NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (section_id) REFERENCES sections(id) ON DELETE CASCADE,
    FOREIGN KEY (parent_id) REFERENCES folders(id) ON DELETE SET NULL,
    UNIQUE (section_id, parent_id, name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS subnets (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    section_id BIGINT NOT NULL,
    parent_id BIGINT NULL,
    cidr VARCHAR(64) NOT NULL,
    ip_version TINYINT UNSIGNED NOT NULL,
    network VARBINARY(16) NOT NULL,
    prefix_len TINYINT UNSIGNED NOT NULL,
    name VARCHAR(128) NOT NULL DEFAULT '',
    description TEXT,
    vlan_id BIGINT NULL,
    vrf_id BIGINT NULL,
    location_id BIGINT NULL,
    nameserver_id BIGINT NULL,
    scan_enabled TINYINT(1) NOT NULL DEFAULT 0,
    discover_enabled TINYINT(1) NOT NULL DEFAULT 0,
    threshold_percent INT NOT NULL DEFAULT 90,
    allow_request TINYINT(1) NOT NULL DEFAULT 0,
    status VARCHAR(32) NOT NULL DEFAULT 'active',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (section_id) REFERENCES sections(id) ON DELETE CASCADE,
    FOREIGN KEY (parent_id) REFERENCES subnets(id) ON DELETE SET NULL,
    UNIQUE (section_id, vrf_id, network, prefix_len)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE INDEX idx_subnets_section ON subnets(section_id);
CREATE INDEX idx_subnets_parent ON subnets(parent_id);
CREATE INDEX idx_subnets_vrf ON subnets(vrf_id);
CREATE INDEX idx_subnets_network ON subnets(network, prefix_len);

CREATE TABLE IF NOT EXISTS addresses (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    subnet_id BIGINT NOT NULL,
    ip VARBINARY(16) NOT NULL,
    ip_text VARCHAR(64) NOT NULL,
    hostname VARCHAR(255) NOT NULL DEFAULT '',
    mac VARCHAR(32) NOT NULL DEFAULT '',
    description TEXT,
    status VARCHAR(32) NOT NULL DEFAULT 'free',
    tag_id BIGINT NULL,
    is_gateway TINYINT(1) NOT NULL DEFAULT 0,
    device_id BIGINT NULL,
    device_port VARCHAR(64) NOT NULL DEFAULT '',
    owner VARCHAR(128) NOT NULL DEFAULT '',
    last_seen_at DATETIME NULL,
    dns_checked_at DATETIME NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (subnet_id) REFERENCES subnets(id) ON DELETE CASCADE,
    UNIQUE (subnet_id, ip)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE INDEX idx_addresses_subnet ON addresses(subnet_id);
CREATE INDEX idx_addresses_ip_text ON addresses(ip_text);
CREATE INDEX idx_addresses_hostname ON addresses(hostname);
CREATE INDEX idx_addresses_status ON addresses(status);
CREATE INDEX idx_addresses_device ON addresses(device_id);

CREATE TABLE IF NOT EXISTS tags (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(128) NOT NULL UNIQUE,
    color VARCHAR(16) NOT NULL DEFAULT '#1890ff',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS addresses;
DROP TABLE IF EXISTS subnets;
DROP TABLE IF EXISTS folders;
DROP TABLE IF EXISTS sections;
