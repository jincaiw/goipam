-- name: GetIPRequestByID :one
SELECT id, subnet_id, applicant_id, requested_ip, hostname, description, status, approver_id, approve_comment, created_at, updated_at FROM ip_requests WHERE id = ?;

-- name: ListIPRequests :many
SELECT id, subnet_id, applicant_id, requested_ip, hostname, description, status, approver_id, approve_comment, created_at, updated_at FROM ip_requests ORDER BY created_at DESC LIMIT ? OFFSET ?;

-- name: CountIPRequests :one
SELECT COUNT(*) FROM ip_requests;

-- name: CreateIPRequest :execresult
INSERT INTO ip_requests (subnet_id, applicant_id, requested_ip, hostname, description, status) VALUES (?, ?, ?, ?, ?, 'pending');

-- name: UpdateIPRequestStatus :exec
UPDATE ip_requests SET status = ?, approver_id = ?, approve_comment = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?;

-- name: DeleteIPRequest :exec
DELETE FROM ip_requests WHERE id = ?;
