-- +goose Up
CREATE TABLE IF NOT EXISTS ip_requests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    subnet_id INTEGER NOT NULL,
    applicant_id INTEGER NOT NULL,
    requested_ip VARCHAR(64) NOT NULL DEFAULT '',
    hostname VARCHAR(255) NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '',
    status VARCHAR(32) NOT NULL DEFAULT 'pending',
    approver_id INTEGER NULL,
    approve_comment TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (subnet_id) REFERENCES subnets(id) ON DELETE CASCADE,
    FOREIGN KEY (applicant_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (approver_id) REFERENCES users(id) ON DELETE SET NULL
);

CREATE INDEX idx_ip_requests_subnet ON ip_requests(subnet_id);
CREATE INDEX idx_ip_requests_applicant ON ip_requests(applicant_id);
CREATE INDEX idx_ip_requests_status ON ip_requests(status);

-- +goose Down
DROP TABLE IF EXISTS ip_requests;
