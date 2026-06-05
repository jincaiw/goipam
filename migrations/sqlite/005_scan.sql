-- +goose Up
CREATE TABLE IF NOT EXISTS scan_tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    subnet_id INTEGER NOT NULL,
    task_type VARCHAR(32) NOT NULL DEFAULT 'ping',
    status VARCHAR(32) NOT NULL DEFAULT 'pending',
    started_at DATETIME NULL,
    finished_at DATETIME NULL,
    total_count INTEGER NOT NULL DEFAULT 0,
    alive_count INTEGER NOT NULL DEFAULT 0,
    new_count INTEGER NOT NULL DEFAULT 0,
    error TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (subnet_id) REFERENCES subnets(id) ON DELETE CASCADE
);

CREATE INDEX idx_scan_tasks_subnet ON scan_tasks(subnet_id);
CREATE INDEX idx_scan_tasks_status ON scan_tasks(status);

CREATE TABLE IF NOT EXISTS scan_results (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task_id INTEGER NOT NULL,
    ip_text VARCHAR(64) NOT NULL,
    status VARCHAR(32) NOT NULL DEFAULT 'unknown',
    message TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (task_id) REFERENCES scan_tasks(id) ON DELETE CASCADE,
    UNIQUE (task_id, ip_text)
);

CREATE INDEX idx_scan_results_task ON scan_results(task_id);

-- +goose Down
DROP TABLE IF EXISTS scan_results;
DROP TABLE IF EXISTS scan_tasks;
