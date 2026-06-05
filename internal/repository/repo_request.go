package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// IPRequest represents an IP request domain object.
type IPRequest struct {
	ID             int64     `json:"id"`
	SubnetID       int64     `json:"subnet_id"`
	ApplicantID    int64     `json:"applicant_id"`
	RequestedIP    string    `json:"requested_ip"`
	Hostname       string    `json:"hostname"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
	ApproverID     *int64    `json:"approver_id"`
	ApproveComment string    `json:"approve_comment"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// RequestRepository handles IP request data access.
type RequestRepository struct {
	q *sqlc.Queries
}

// NewRequestRepository creates a new RequestRepository.
func NewRequestRepository(q *sqlc.Queries) *RequestRepository {
	return &RequestRepository{q: q}
}

// GetByID retrieves an IP request by ID.
func (r *RequestRepository) GetByID(ctx context.Context, id int64) (*IPRequest, error) {
	row, err := r.q.GetIPRequestByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get ip request by id: %w", err)
	}
	return sqlcIPRequestToIPRequest(row), nil
}

// List retrieves a paginated list of IP requests.
func (r *RequestRepository) List(ctx context.Context, limit, offset int32) ([]*IPRequest, error) {
	rows, err := r.q.ListIPRequests(ctx, sqlc.ListIPRequestsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list ip requests: %w", err)
	}
	items := make([]*IPRequest, len(rows))
	for i, row := range rows {
		items[i] = sqlcIPRequestToIPRequest(row)
	}
	return items, nil
}

// Count returns the total number of IP requests.
func (r *RequestRepository) Count(ctx context.Context) (int64, error) {
	return r.q.CountIPRequests(ctx)
}

// Create creates a new IP request and returns the ID.
func (r *RequestRepository) Create(ctx context.Context, subnetID, applicantID int64, requestedIP, hostname, description string) (int64, error) {
	result, err := r.q.CreateIPRequest(ctx, sqlc.CreateIPRequestParams{
		SubnetID:    subnetID,
		ApplicantID: applicantID,
		RequestedIp: requestedIP,
		Hostname:    hostname,
		Description: NullString(&description),
	})
	if err != nil {
		return 0, fmt.Errorf("create ip request: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// UpdateStatus updates an IP request's status.
func (r *RequestRepository) UpdateStatus(ctx context.Context, id int64, status string, approverID *int64, approveComment string) error {
	return r.q.UpdateIPRequestStatus(ctx, sqlc.UpdateIPRequestStatusParams{
		Status:         status,
		ApproverID:     NullInt64(approverID),
		ApproveComment: sql.NullString{String: approveComment, Valid: approveComment != ""},
		ID:             id,
	})
}

// Delete deletes an IP request.
func (r *RequestRepository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteIPRequest(ctx, id)
}

func sqlcIPRequestToIPRequest(row sqlc.IpRequest) *IPRequest {
	return &IPRequest{
		ID:             row.ID,
		SubnetID:       row.SubnetID,
		ApplicantID:    row.ApplicantID,
		RequestedIP:    row.RequestedIp,
		Hostname:       row.Hostname,
		Description:    row.Description.String,
		Status:         row.Status,
		ApproverID:     nullInt64Ptr(row.ApproverID),
		ApproveComment: row.ApproveComment.String,
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
	}
}
