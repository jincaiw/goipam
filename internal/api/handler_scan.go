package api

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"log/slog"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/domain/ipam"
	"github.com/jincaiw/goipam/internal/domain/scan"
	"github.com/jincaiw/goipam/internal/repository"
)

// ScanHandler handles scan API endpoints.
type ScanHandler struct {
	scanRepo    *repository.ScanRepository
	subnetRepo  *repository.SubnetRepository
	addressRepo *repository.AddressRepository
	mu          sync.Mutex
	running     map[int64]context.CancelFunc
}

// NewScanHandler creates a new ScanHandler.
func NewScanHandler(scanRepo *repository.ScanRepository, subnetRepo *repository.SubnetRepository, addressRepo *repository.AddressRepository) *ScanHandler {
	return &ScanHandler{
		scanRepo:    scanRepo,
		subnetRepo:  subnetRepo,
		addressRepo: addressRepo,
		running:     make(map[int64]context.CancelFunc),
	}
}

// CreateScan handles POST /api/v1/scans/subnets/{id}
func (h *ScanHandler) CreateScan(w http.ResponseWriter, r *http.Request) {
	subnetID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid subnet id")
		return
	}

	subnet, err := h.subnetRepo.GetByID(r.Context(), subnetID)
	if err != nil {
		NotFound(w, r, "subnet")
		return
	}

	// Parse CIDR to get total addresses
	info, err := ipam.ParseCIDR(subnet.Cidr)
	if err != nil {
		BadRequest(w, r, errcode.SubnetInvalidCIDR, "invalid subnet CIDR")
		return
	}

	if info.TotalAddresses > 65536 {
		Error(w, r, http.StatusBadRequest, errcode.ScanTooLarge, "subnet is too large to scan (max 65536 addresses)")
		return
	}

	// Check if there is already a running scan for this subnet
	h.mu.Lock()
	if cancel, ok := h.running[subnetID]; ok {
		h.mu.Unlock()
		_ = cancel
		Error(w, r, http.StatusConflict, errcode.ScanAlreadyRunning, "a scan is already running for this subnet")
		return
	}

	// Create scan task
	taskID, err := h.scanRepo.CreateTask(r.Context(), subnetID, "ping")
	if err != nil {
		h.mu.Unlock()
		InternalError(w, r, "failed to create scan task")
		return
	}

	// Launch scan in background
	ctx, cancel := context.WithCancel(r.Context())
	h.running[subnetID] = cancel
	h.mu.Unlock()

	go h.runScan(ctx, taskID, subnetID, subnet.Cidr, info.TotalAddresses)

	Created(w, r, map[string]int64{"id": taskID})
}

// ListTasks handles GET /api/v1/scans/tasks
func (h *ScanHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	page, pageSize := parsePagination(r)
	offset := int32((page - 1) * pageSize)

	tasks, err := h.scanRepo.ListTasks(r.Context(), int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list scan tasks")
		return
	}

	total, err := h.scanRepo.CountTasks(r.Context())
	if err != nil {
		InternalError(w, r, "failed to count scan tasks")
		return
	}

	OKPaged(w, r, tasks, total, page, pageSize)
}

// GetTask handles GET /api/v1/scans/tasks/{id}
func (h *ScanHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid task id")
		return
	}

	task, err := h.scanRepo.GetTaskByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "scan task")
		return
	}

	// Also get results if task is completed or running
	results, _ := h.scanRepo.GetResults(r.Context(), id)

	OK(w, r, map[string]interface{}{
		"task":    task,
		"results": results,
	})
}

// CancelTask handles POST /api/v1/scans/tasks/{id}/cancel
func (h *ScanHandler) CancelTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid task id")
		return
	}

	task, err := h.scanRepo.GetTaskByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "scan task")
		return
	}

	if task.Status != "running" && task.Status != "pending" {
		BadRequest(w, r, errcode.ValidationError, "task is not running or pending")
		return
	}

	// Cancel the running goroutine
	h.mu.Lock()
	if cancel, ok := h.running[task.SubnetID]; ok {
		cancel()
		delete(h.running, task.SubnetID)
	}
	h.mu.Unlock()

	// Update task status to cancelled
	now := time.Now()
	h.scanRepo.UpdateTaskStatus(r.Context(), id, "cancelled", task.StartedAt, &now, task.AliveCount, task.NewCount, "cancelled by user")

	OK(w, r, map[string]string{"message": "scan task cancelled"})
}

// runScan executes the actual scan in a background goroutine.
func (h *ScanHandler) runScan(ctx context.Context, taskID, subnetID int64, cidr string, totalAddresses int64) {
	defer func() {
		h.mu.Lock()
		delete(h.running, subnetID)
		h.mu.Unlock()
	}()

	slog.Info("starting scan", "task_id", taskID, "subnet_id", subnetID, "cidr", cidr)

	// Update task to running
	now := time.Now()
	if err := h.scanRepo.UpdateTaskStatus(ctx, taskID, "running", &now, nil, 0, 0, ""); err != nil {
		slog.Error("failed to update scan task status", "error", err)
		return
	}

	// Get existing IPs in the subnet
	existingIPs, _ := h.addressRepo.ListIPsBySubnet(ctx, subnetID)
	existingSet := make(map[string]bool)
	for _, ip := range existingIPs {
		existingSet[ip] = true
	}

	// Parse CIDR and iterate through all IPs
	info, err := ipam.ParseCIDR(cidr)
	if err != nil {
		errMsg := err.Error()
		h.scanRepo.UpdateTaskStatus(ctx, taskID, "failed", &now, nil, 0, 0, errMsg)
		return
	}

	// Generate all IPs in the subnet
	ip := parseIP(info.FirstIP)
	lastIP := parseIP(info.LastIP)
	if ip == nil || lastIP == nil {
		h.scanRepo.UpdateTaskStatus(ctx, taskID, "failed", &now, nil, 0, 0, "invalid IP range")
		return
	}

	var aliveCount int32
	var newCount int32
	probeTimeout := 500 * time.Millisecond

	for !ip.Equal(lastIP) {
		select {
		case <-ctx.Done():
			finishedAt := time.Now()
			h.scanRepo.UpdateTaskStatus(ctx, taskID, "cancelled", &now, &finishedAt, aliveCount, newCount, "cancelled")
			return
		default:
		}

		ipStr := ip.String()
		result := scan.PingProbe(ctx, ipStr, probeTimeout)

		if result.Alive {
			aliveCount++
			isNew := !existingSet[ipStr]
			status := "alive"
			if isNew {
				newCount++
				status = "alive_new"
			}
			h.scanRepo.CreateResult(ctx, taskID, ipStr, status, result.Message)
		} else {
			h.scanRepo.CreateResult(ctx, taskID, ipStr, "dead", result.Message)
		}

		incrementIP(ip)
	}

	// Probe the last IP
	ipStr := lastIP.String()
	result := scan.PingProbe(ctx, ipStr, probeTimeout)
	if result.Alive {
		aliveCount++
		isNew := !existingSet[ipStr]
		status := "alive"
		if isNew {
			newCount++
			status = "alive_new"
		}
		h.scanRepo.CreateResult(ctx, taskID, ipStr, status, result.Message)
	} else {
		h.scanRepo.CreateResult(ctx, taskID, ipStr, "dead", result.Message)
	}

	// Mark task as completed
	finishedAt := time.Now()
	h.scanRepo.UpdateTaskStatus(ctx, taskID, "completed", &now, &finishedAt, aliveCount, newCount, "")

	slog.Info("scan completed", "task_id", taskID, "alive", aliveCount, "new", newCount)
}

func parseIP(s string) netIP {
	ip := net.ParseIP(s)
	if ip == nil {
		return nil
	}
	if v4 := ip.To4(); v4 != nil {
		return v4
	}
	return ip
}

type netIP = net.IP

func incrementIP(ip netIP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
}
