-- +goose Up
CREATE TABLE IF NOT EXISTS sections (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(128) NOT NULL UNIQUE,
    description TEXT NOT NULL DEFAULT '',
    parent_id INTEGER NULL,
    strict_mode BOOLEAN NOT NULL DEFAULT 0,
    show_vlan BOOLEAN NOT NULL DEFAULT 1,
    show_vrf BOOLEAN NOT NULL DEFAULT 1,
    subnet_order VARCHAR(32) NOT NULL DEFAULT 'subnet',
    sort_order INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES sections(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS folders (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    section_id INTEGER NOT NULL,
    parent_id INTEGER NULL,
    name VARCHAR(128) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    sort_order INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (section_id) REFERENCES sections(id) ON DELETE CASCADE,
    FOREIGN KEY (parent_id) REFERENCES folders(id) ON DELETE SET NULL,
    UNIQUE (section_id, parent_id, name)
);

CREATE TABLE IF NOT EXISTS subnets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    section_id INTEGER NOT NULL,
    parent_id INTEGER NULL,
    cidr VARCHAR(64) NOT NULL,
    ip_version TINYINT NOT NULL,
    network BLOB NOT NULL,
    prefix_len TINYINT NOT NULL,
    name VARCHAR(128) NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '',
    vlan_id INTEGER NULL,
    vrf_id INTEGER NULL,
    location_id INTEGER NULL,
    nameserver_id INTEGER NULL,
    scan_enabled BOOLEAN NOT NULL DEFAULT 0,
    discover_enabled BOOLEAN NOT NULL DEFAULT 0,
    threshold_percent INTEGER NOT NULL DEFAULT 90,
    allow_request BOOLEAN NOT NULL DEFAULT 0,
    status VARCHAR(32) NOT NULL DEFAULT 'active',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (section_id) REFERENCES sections(id) ON DELETE CASCADE,
    FOREIGN KEY (parent_id) REFERENCES subnets(id) ON DELETE SET NULL,
    UNIQUE (section_id, vrf_id, network, prefix_len)
);

CREATE INDEX idx_subnets_section ON subnets(section_id);
CREATE INDEX idx_subnets_parent ON subnets(parent_id);
CREATE INDEX idx_subnets_vrf ON subnets(vrf_id);
CREATE INDEX idx_subnets_network ON subnets(network, prefix_len);

CREATE TABLE IF NOT EXISTS addresses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    subnet_id INTEGER NOT NULL,
    ip BLOB NOT NULL,
    ip_text VARCHAR(64) NOT NULL,
    hostname VARCHAR(255) NOT NULL DEFAULT '',
    mac VARCHAR(32) NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '',
    status VARCHAR(32) NOT NULL DEFAULT 'free',
    tag_id INTEGER NULL,
    is_gateway BOOLEAN NOT NULL DEFAULT 0,
    device_id INTEGER NULL,
    device_port VARCHAR(64) NOT NULL DEFAULT '',
    owner VARCHAR(128) NOT NULL DEFAULT '',
    last_seen_at DATETIME NULL,
    dns_checked_at DATETIME NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (subnet_id) REFERENCES subnets(id) ON DELETE CASCADE,
    UNIQUE (subnet_id, ip)
);

CREATE INDEX idx_addresses_subnet ON addresses(subnet_id);
CREATE INDEX idx_addresses_ip_text ON addresses(ip_text);
CREATE INDEX idx_addresses_hostname ON addresses(hostname);
CREATE INDEX idx_addresses_status ON addresses(status);
CREATE INDEX idx_addresses_device ON addresses(device_id);

CREATE TABLE IF NOT EXISTS tags (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(128) NOT NULL UNIQUE,
    color VARCHAR(16) NOT NULL DEFAULT '#1890ff',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS addresses;
DROP TABLE IF EXISTS subnets;
DROP TABLE IF EXISTS folders;
DROP TABLE IF EXISTS sections;
