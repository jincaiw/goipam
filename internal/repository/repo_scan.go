package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// ScanTask represents a scan task domain object.
type ScanTask struct {
	ID         int64      `json:"id"`
	SubnetID   int64      `json:"subnet_id"`
	TaskType   string     `json:"task_type"`
	Status     string     `json:"status"`
	StartedAt  *time.Time `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at"`
	TotalCount int32      `json:"total_count"`
	AliveCount int32      `json:"alive_count"`
	NewCount   int32      `json:"new_count"`
	Error      string     `json:"error"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// ScanResult represents a scan result domain object.
type ScanResult struct {
	ID        int64     `json:"id"`
	TaskID    int64     `json:"task_id"`
	IpText    string    `json:"ip_text"`
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

// ScanRepository handles scan data access.
type ScanRepository struct {
	q *sqlc.Queries
}

// NewScanRepository creates a new ScanRepository.
func NewScanRepository(q *sqlc.Queries) *ScanRepository {
	return &ScanRepository{q: q}
}

// GetTaskByID retrieves a scan task by ID.
func (r *ScanRepository) GetTaskByID(ctx context.Context, id int64) (*ScanTask, error) {
	row, err := r.q.GetScanTaskByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get scan task by id: %w", err)
	}
	return sqlcScanTaskToScanTask(row), nil
}

// ListTasks retrieves a paginated list of scan tasks.
func (r *ScanRepository) ListTasks(ctx context.Context, limit, offset int32) ([]*ScanTask, error) {
	rows, err := r.q.ListScanTasks(ctx, sqlc.ListScanTasksParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list scan tasks: %w", err)
	}
	items := make([]*ScanTask, len(rows))
	for i, row := range rows {
		items[i] = sqlcScanTaskToScanTask(row)
	}
	return items, nil
}

// CountTasks returns the total number of scan tasks.
func (r *ScanRepository) CountTasks(ctx context.Context) (int64, error) {
	return r.q.CountScanTasks(ctx)
}

// CreateTask creates a new scan task and returns the ID.
func (r *ScanRepository) CreateTask(ctx context.Context, subnetID int64, taskType string) (int64, error) {
	result, err := r.q.CreateScanTask(ctx, sqlc.CreateScanTaskParams{
		SubnetID: subnetID,
		TaskType: taskType,
	})
	if err != nil {
		return 0, fmt.Errorf("create scan task: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// UpdateTaskStatus updates a scan task's status and results.
func (r *ScanRepository) UpdateTaskStatus(ctx context.Context, id int64, status string, startedAt, finishedAt *time.Time, aliveCount, newCount int32, errMsg string) error {
	return r.q.UpdateScanTaskStatus(ctx, sqlc.UpdateScanTaskStatusParams{
		Status:     status,
		StartedAt:  NullTime(startedAt),
		FinishedAt: NullTime(finishedAt),
		AliveCount: aliveCount,
		NewCount:   newCount,
		Error:      sql.NullString{String: errMsg, Valid: errMsg != ""},
		ID:         id,
	})
}

// DeleteTask deletes a scan task.
func (r *ScanRepository) DeleteTask(ctx context.Context, id int64) error {
	return r.q.DeleteScanTask(ctx, id)
}

// GetResults retrieves scan results for a task.
func (r *ScanRepository) GetResults(ctx context.Context, taskID int64) ([]*ScanResult, error) {
	rows, err := r.q.GetScanResults(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("get scan results: %w", err)
	}
	items := make([]*ScanResult, len(rows))
	for i, row := range rows {
		items[i] = sqlcScanResultToScanResult(row)
	}
	return items, nil
}

// CreateResult creates a new scan result.
func (r *ScanRepository) CreateResult(ctx context.Context, taskID int64, ipText, status, message string) (int64, error) {
	result, err := r.q.CreateScanResult(ctx, sqlc.CreateScanResultParams{
		TaskID:  taskID,
		IpText:  ipText,
		Status:  status,
		Message: sql.NullString{String: message, Valid: message != ""},
	})
	if err != nil {
		return 0, fmt.Errorf("create scan result: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

func sqlcScanTaskToScanTask(row sqlc.ScanTask) *ScanTask {
	return &ScanTask{
		ID:         row.ID,
		SubnetID:   row.SubnetID,
		TaskType:   row.TaskType,
		Status:     row.Status,
		StartedAt:  nullTimePtr(row.StartedAt),
		FinishedAt: nullTimePtr(row.FinishedAt),
		TotalCount: row.TotalCount,
		AliveCount: row.AliveCount,
		NewCount:   row.NewCount,
		Error:      row.Error.String,
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}
}

func sqlcScanResultToScanResult(row sqlc.ScanResult) *ScanResult {
	return &ScanResult{
		ID:        row.ID,
		TaskID:    row.TaskID,
		IpText:    row.IpText,
		Status:    row.Status,
		Message:   row.Message.String,
		CreatedAt: row.CreatedAt,
	}
}
