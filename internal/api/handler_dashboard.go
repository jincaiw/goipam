package api

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"time"
)

// DashboardHandler handles dashboard API endpoints.
type DashboardHandler struct {
	db *sql.DB
}

// NewDashboardHandler creates a new DashboardHandler.
func NewDashboardHandler(db *sql.DB) *DashboardHandler {
	return &DashboardHandler{db: db}
}

// DashboardStats represents the dashboard statistics response.
type DashboardStats struct {
	TotalSections        int64                 `json:"total_sections"`
	TotalSubnets         int64                 `json:"total_subnets"`
	TotalAddresses       int64                 `json:"total_addresses"`
	AddressUsage         AddressUsageStats     `json:"address_usage"`
	RecentSubnets        []RecentSubnet        `json:"recent_subnets"`
	RecentAddressChanges []RecentAddressChange `json:"recent_address_changes"`
	ThresholdSubnets     []ThresholdSubnet     `json:"threshold_subnets"`
}

// AddressUsageStats represents address usage statistics.
type AddressUsageStats struct {
	Used      int64 `json:"used"`
	Available int64 `json:"available"`
	Free      int64 `json:"free"`
}

// RecentSubnet represents a recently created subnet.
type RecentSubnet struct {
	ID        int64     `json:"id"`
	Cidr      string    `json:"cidr"`
	Name      string    `json:"name"`
	SectionID int64     `json:"section_id"`
	CreatedAt time.Time `json:"created_at"`
}

// RecentAddressChange represents a recent address change.
type RecentAddressChange struct {
	ID        int64     `json:"id"`
	IpText    string    `json:"ip_text"`
	Hostname  string    `json:"hostname"`
	Status    string    `json:"status"`
	SubnetID  int64     `json:"subnet_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ThresholdSubnet represents a subnet approaching its usage threshold.
type ThresholdSubnet struct {
	ID               int64   `json:"id"`
	Cidr             string  `json:"cidr"`
	Name             string  `json:"name"`
	UsedCount        int64   `json:"used_count"`
	TotalCount       int64   `json:"total_count"`
	UsagePercent     float64 `json:"usage_percent"`
	ThresholdPercent int32   `json:"threshold_percent"`
}

// Dashboard handles GET /api/v1/dashboard
func (h *DashboardHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	stats := &DashboardStats{}

	// Get total counts
	stats.TotalSections = h.queryCount(ctx, "SELECT COUNT(*) FROM sections")
	stats.TotalSubnets = h.queryCount(ctx, "SELECT COUNT(*) FROM subnets")
	stats.TotalAddresses = h.queryCount(ctx, "SELECT COUNT(*) FROM addresses")

	// Get address usage stats
	stats.AddressUsage = h.getAddressUsage(ctx)

	// Get recent subnets
	stats.RecentSubnets = h.getRecentSubnets(ctx)

	// Get recent address changes
	stats.RecentAddressChanges = h.getRecentAddressChanges(ctx)

	// Get subnets approaching threshold
	stats.ThresholdSubnets = h.getThresholdSubnets(ctx)

	OK(w, r, stats)
}

func (h *DashboardHandler) queryCount(ctx context.Context, query string) int64 {
	var count int64
	err := h.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		slog.Error("dashboard count query failed", "query", query, "error", err)
		return 0
	}
	return count
}

func (h *DashboardHandler) getAddressUsage(ctx context.Context) AddressUsageStats {
	var usage AddressUsageStats

	rows, err := h.db.QueryContext(ctx, "SELECT status, COUNT(*) FROM addresses GROUP BY status")
	if err != nil {
		return usage
	}
	defer rows.Close()

	for rows.Next() {
		var status string
		var count int64
		if err := rows.Scan(&status, &count); err != nil {
			continue
		}
		switch status {
		case "used", "online", "offline", "reserved":
			usage.Used += count
		case "free":
			usage.Free += count
		default:
			usage.Available += count
		}
	}
	if err := rows.Err(); err != nil {
		slog.Error("dashboard address usage rows error", "error", err)
	}

	return usage
}

func (h *DashboardHandler) getRecentSubnets(ctx context.Context) []RecentSubnet {
	var result []RecentSubnet

	rows, err := h.db.QueryContext(ctx,
		"SELECT id, cidr, name, section_id, created_at FROM subnets ORDER BY created_at DESC LIMIT 5")
	if err != nil {
		return result
	}
	defer rows.Close()

	for rows.Next() {
		var s RecentSubnet
		if err := rows.Scan(&s.ID, &s.Cidr, &s.Name, &s.SectionID, &s.CreatedAt); err != nil {
			continue
		}
		result = append(result, s)
	}
	if err := rows.Err(); err != nil {
		slog.Error("dashboard recent subnets rows error", "error", err)
	}

	return result
}

func (h *DashboardHandler) getRecentAddressChanges(ctx context.Context) []RecentAddressChange {
	var result []RecentAddressChange

	rows, err := h.db.QueryContext(ctx,
		"SELECT id, ip_text, hostname, status, subnet_id, updated_at FROM addresses ORDER BY updated_at DESC LIMIT 10")
	if err != nil {
		return result
	}
	defer rows.Close()

	for rows.Next() {
		var a RecentAddressChange
		if err := rows.Scan(&a.ID, &a.IpText, &a.Hostname, &a.Status, &a.SubnetID, &a.UpdatedAt); err != nil {
			continue
		}
		result = append(result, a)
	}
	if err := rows.Err(); err != nil {
		slog.Error("dashboard recent address changes rows error", "error", err)
	}

	return result
}

func (h *DashboardHandler) getThresholdSubnets(ctx context.Context) []ThresholdSubnet {
	var result []ThresholdSubnet

	// Query subnets with their address usage, filtering those >80% used
	query := `
		SELECT
			s.id, s.cidr, s.name,
			COALESCE(used.cnt, 0) AS used_count,
			COALESCE(total.cnt, 0) AS total_count,
			s.threshold_percent
		FROM subnets s
		LEFT JOIN (SELECT subnet_id, COUNT(*) AS cnt FROM addresses WHERE status IN ('used','online','offline','reserved') GROUP BY subnet_id) used ON used.subnet_id = s.id
		LEFT JOIN (SELECT subnet_id, COUNT(*) AS cnt FROM addresses GROUP BY subnet_id) total ON total.subnet_id = s.id
		WHERE total.cnt > 0
			AND (CAST(COALESCE(used.cnt, 0) AS REAL) / CAST(total.cnt AS REAL) * 100) > 80
		ORDER BY (CAST(COALESCE(used.cnt, 0) AS REAL) / CAST(total.cnt AS REAL)) DESC
		LIMIT 20
	`

	rows, err := h.db.QueryContext(ctx, query)
	if err != nil {
		slog.Error("threshold query error", "error", err)
		return result
	}
	defer rows.Close()

	for rows.Next() {
		var t ThresholdSubnet
		if err := rows.Scan(&t.ID, &t.Cidr, &t.Name, &t.UsedCount, &t.TotalCount, &t.ThresholdPercent); err != nil {
			continue
		}
		if t.TotalCount > 0 {
			t.UsagePercent = float64(t.UsedCount) / float64(t.TotalCount) * 100
		}
		result = append(result, t)
	}
	if err := rows.Err(); err != nil {
		slog.Error("dashboard threshold subnets rows error", "error", err)
	}

	return result
}
