package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// AuditLog represents an audit log domain object.
type AuditLog struct {
	ID           int64     `json:"id"`
	ActorID      *int64    `json:"actor_id"`
	Action       string    `json:"action"`
	ResourceType string    `json:"resource_type"`
	ResourceID   int64     `json:"resource_id"`
	BeforeJSON   string    `json:"before_json"`
	AfterJSON    string    `json:"after_json"`
	IP           string    `json:"ip"`
	UserAgent    string    `json:"user_agent"`
	RequestID    string    `json:"request_id"`
	CreatedAt    time.Time `json:"created_at"`
}

// AuditRepository handles audit log data access.
type AuditRepository struct {
	q *sqlc.Queries
}

// NewAuditRepository creates a new AuditRepository.
func NewAuditRepository(q *sqlc.Queries) *AuditRepository {
	return &AuditRepository{q: q}
}

// List retrieves a paginated list of audit logs.
func (r *AuditRepository) List(ctx context.Context, limit, offset int32) ([]*AuditLog, error) {
	rows, err := r.q.ListAuditLogs(ctx, sqlc.ListAuditLogsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list audit logs: %w", err)
	}
	items := make([]*AuditLog, len(rows))
	for i, row := range rows {
		items[i] = sqlcAuditLogToAuditLog(row)
	}
	return items, nil
}

// Count returns the total number of audit logs.
func (r *AuditRepository) Count(ctx context.Context) (int64, error) {
	return r.q.CountAuditLogs(ctx)
}

// Create creates a new audit log entry and returns the ID.
func (r *AuditRepository) Create(ctx context.Context, actorID *int64, action, resourceType string, resourceID int64, beforeJSON, afterJSON, ip, userAgent, requestID string) (int64, error) {
	result, err := r.q.CreateAuditLog(ctx, sqlc.CreateAuditLogParams{
		ActorID:      NullInt64(actorID),
		Action:       action,
		ResourceType: resourceType,
		ResourceID:   resourceID,
		BeforeJson:   sql.NullString{String: beforeJSON, Valid: beforeJSON != ""},
		AfterJson:    sql.NullString{String: afterJSON, Valid: afterJSON != ""},
		Ip:           ip,
		UserAgent:    userAgent,
		RequestID:    requestID,
	})
	if err != nil {
		return 0, fmt.Errorf("create audit log: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

func sqlcAuditLogToAuditLog(row sqlc.AuditLog) *AuditLog {
	return &AuditLog{
		ID:           row.ID,
		ActorID:      nullInt64Ptr(row.ActorID),
		Action:       row.Action,
		ResourceType: row.ResourceType,
		ResourceID:   row.ResourceID,
		BeforeJSON:   row.BeforeJson.String,
		AfterJSON:    row.AfterJson.String,
		IP:           row.Ip,
		UserAgent:    row.UserAgent,
		RequestID:    row.RequestID,
		CreatedAt:    row.CreatedAt,
	}
}
