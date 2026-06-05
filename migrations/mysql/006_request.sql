-- +goose Up
CREATE TABLE IF NOT EXISTS ip_requests (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    subnet_id BIGINT NOT NULL,
    applicant_id BIGINT NOT NULL,
    requested_ip VARCHAR(64) NOT NULL DEFAULT '',
    hostname VARCHAR(255) NOT NULL DEFAULT '',
    description TEXT,
    status VARCHAR(32) NOT NULL DEFAULT 'pending',
    approver_id BIGINT NULL,
    approve_comment TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (subnet_id) REFERENCES subnets(id) ON DELETE CASCADE,
    FOREIGN KEY (applicant_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (approver_id) REFERENCES users(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE INDEX idx_ip_requests_subnet ON ip_requests(subnet_id);
CREATE INDEX idx_ip_requests_applicant ON ip_requests(applicant_id);
CREATE INDEX idx_ip_requests_status ON ip_requests(status);

-- +goose Down
DROP TABLE IF EXISTS ip_requests;
