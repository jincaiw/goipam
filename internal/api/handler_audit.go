package api

import (
	"net/http"

	"github.com/jincaiw/goipam/internal/repository"
)

// AuditHandler handles audit log API endpoints.
type AuditHandler struct {
	auditRepo *repository.AuditRepository
}

// NewAuditHandler creates a new AuditHandler.
func NewAuditHandler(auditRepo *repository.AuditRepository) *AuditHandler {
	return &AuditHandler{auditRepo: auditRepo}
}

// List handles GET /api/v1/audit-logs
func (h *AuditHandler) List(w http.ResponseWriter, r *http.Request) {
	page, pageSize := parsePagination(r)
	offset := int32((page - 1) * pageSize)

	logs, err := h.auditRepo.List(r.Context(), int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list audit logs")
		return
	}

	total, err := h.auditRepo.Count(r.Context())
	if err != nil {
		InternalError(w, r, "failed to count audit logs")
		return
	}

	OKPaged(w, r, logs, total, page, pageSize)
}
