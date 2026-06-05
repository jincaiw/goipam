-- +goose Up
CREATE TABLE IF NOT EXISTS scan_tasks (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    subnet_id BIGINT NOT NULL,
    task_type VARCHAR(32) NOT NULL DEFAULT 'ping',
    status VARCHAR(32) NOT NULL DEFAULT 'pending',
    started_at DATETIME NULL,
    finished_at DATETIME NULL,
    total_count INT NOT NULL DEFAULT 0,
    alive_count INT NOT NULL DEFAULT 0,
    new_count INT NOT NULL DEFAULT 0,
    error TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (subnet_id) REFERENCES subnets(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE INDEX idx_scan_tasks_subnet ON scan_tasks(subnet_id);
CREATE INDEX idx_scan_tasks_status ON scan_tasks(status);

CREATE TABLE IF NOT EXISTS scan_results (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    task_id BIGINT NOT NULL,
    ip_text VARCHAR(64) NOT NULL,
    status VARCHAR(32) NOT NULL DEFAULT 'unknown',
    message TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (task_id) REFERENCES scan_tasks(id) ON DELETE CASCADE,
    UNIQUE (task_id, ip_text)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE INDEX idx_scan_results_task ON scan_results(task_id);

-- +goose Down
DROP TABLE IF EXISTS scan_results;
DROP TABLE IF EXISTS scan_tasks;
