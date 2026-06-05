-- name: GetScanTaskByID :one
SELECT id, subnet_id, task_type, status, started_at, finished_at, total_count, alive_count, new_count, error, created_at, updated_at FROM scan_tasks WHERE id = ?;

-- name: ListScanTasks :many
SELECT id, subnet_id, task_type, status, started_at, finished_at, total_count, alive_count, new_count, error, created_at, updated_at FROM scan_tasks ORDER BY created_at DESC LIMIT ? OFFSET ?;

-- name: CountScanTasks :one
SELECT COUNT(*) FROM scan_tasks;

-- name: CreateScanTask :execresult
INSERT INTO scan_tasks (subnet_id, task_type, status, total_count) VALUES (?, ?, 'pending', 0);

-- name: UpdateScanTaskStatus :exec
UPDATE scan_tasks SET status = ?, started_at = ?, finished_at = ?, alive_count = ?, new_count = ?, error = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;

-- name: DeleteScanTask :exec
DELETE FROM scan_tasks WHERE id = ?;

-- name: GetScanResults :many
SELECT id, task_id, ip_text, status, message, created_at FROM scan_results WHERE task_id = ? ORDER BY ip_text;

-- name: CreateScanResult :execresult
INSERT INTO scan_results (task_id, ip_text, status, message) VALUES (?, ?, ?, ?);
