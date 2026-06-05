package audit

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jincaiw/goipam/internal/repository"
)

// AuditLogger handles audit log recording.
type AuditLogger struct {
	repo *repository.AuditRepository
}

// NewAuditLogger creates a new AuditLogger.
func NewAuditLogger(repo *repository.AuditRepository) *AuditLogger {
	return &AuditLogger{repo: repo}
}

// Entry represents an audit log entry.
type Entry struct {
	ActorID      *int64      `json:"actor_id"`
	Action       string      `json:"action"`
	ResourceType string      `json:"resource_type"`
	ResourceID   int64       `json:"resource_id"`
	Before       interface{} `json:"before,omitempty"`
	After        interface{} `json:"after,omitempty"`
	IP           string      `json:"ip"`
	UserAgent    string      `json:"user_agent"`
	RequestID    string      `json:"request_id"`
}

// Log records an audit log entry.
func (l *AuditLogger) Log(ctx context.Context, entry Entry) error {
	var beforeJSON, afterJSON string
	if entry.Before != nil {
		b, err := json.Marshal(entry.Before)
		if err == nil {
			beforeJSON = string(b)
		}
	}
	if entry.After != nil {
		b, err := json.Marshal(entry.After)
		if err == nil {
			afterJSON = string(b)
		}
	}

	_, err := l.repo.Create(
		ctx,
		entry.ActorID,
		entry.Action,
		entry.ResourceType,
		entry.ResourceID,
		beforeJSON,
		afterJSON,
		entry.IP,
		entry.UserAgent,
		entry.RequestID,
	)
	return err
}

// AuditEntry is a domain object for audit logs.
type AuditEntry struct {
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
