-- name: ListAuditLogs :many
SELECT id, actor_id, action, resource_type, resource_id, before_json, after_json, ip, user_agent, request_id, created_at FROM audit_logs ORDER BY created_at DESC LIMIT ? OFFSET ?;

-- name: CountAuditLogs :one
SELECT COUNT(*) FROM audit_logs;

-- name: CreateAuditLog :execresult
INSERT INTO audit_logs (actor_id, action, resource_type, resource_id, before_json, after_json, ip, user_agent, request_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
