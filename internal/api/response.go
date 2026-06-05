package api

import (
	"encoding/json"
	"net/http"

	"github.com/jincaiw/goipam/internal/api/errcode"
)

// Response is the unified API response structure.
type Response struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data,omitempty"`
	Code      string      `json:"code,omitempty"`
	Message   string      `json:"message,omitempty"`
	RequestID string      `json:"request_id,omitempty"`
}

// PagedData wraps list data with pagination info.
type PagedData struct {
	Items    interface{} `json:"items"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

// OK sends a successful response.
func OK(w http.ResponseWriter, r *http.Request, data interface{}) {
	reqID := requestIDFromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Success:   true,
		Data:      data,
		RequestID: reqID,
	})
}

// OKPaged sends a successful paged response.
func OKPaged(w http.ResponseWriter, r *http.Request, items interface{}, total int64, page, pageSize int) {
	OK(w, r, PagedData{
		Items:    items,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}

// Created sends a 201 response.
func Created(w http.ResponseWriter, r *http.Request, data interface{}) {
	reqID := requestIDFromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Success:   true,
		Data:      data,
		RequestID: reqID,
	})
}

// Error sends an error response.
func Error(w http.ResponseWriter, r *http.Request, status int, code string, message string) {
	reqID := requestIDFromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Success:   false,
		Code:      code,
		Message:   message,
		RequestID: reqID,
	})
}

// BadRequest sends a 400 error.
func BadRequest(w http.ResponseWriter, r *http.Request, code string, message string) {
	Error(w, r, http.StatusBadRequest, code, message)
}

// Unauthorized sends a 401 error.
func Unauthorized(w http.ResponseWriter, r *http.Request, message string) {
	Error(w, r, http.StatusUnauthorized, errcode.AuthUnauthorized, message)
}

// Forbidden sends a 403 error.
func Forbidden(w http.ResponseWriter, r *http.Request, message string) {
	Error(w, r, http.StatusForbidden, errcode.AuthForbidden, message)
}

// NotFound sends a 404 error.
func NotFound(w http.ResponseWriter, r *http.Request, resource string) {
	Error(w, r, http.StatusNotFound, errcode.ResourceNotFound, resource+" not found")
}

// InternalError sends a 500 error.
func InternalError(w http.ResponseWriter, r *http.Request, message string) {
	Error(w, r, http.StatusInternalServerError, errcode.InternalError, message)
}
