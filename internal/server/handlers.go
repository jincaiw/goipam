package server

import (
	"net/http"

	"github.com/jincaiw/goipam/internal/api"
)

// healthz returns process liveness.
func (s *Server) healthz(w http.ResponseWriter, r *http.Request) {
	api.OK(w, r, map[string]string{"status": "alive"})
}

// readyz returns database readiness.
func (s *Server) readyz(w http.ResponseWriter, r *http.Request) {
	// TODO: Check database connectivity
	api.OK(w, r, map[string]string{"status": "ready"})
}

// placeholderHandler returns a handler that responds with a not-implemented message.
// This is used during initial scaffolding and will be replaced with real handlers.
func (s *Server) placeholderHandler(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api.Error(w, r, http.StatusNotImplemented, "NOT_IMPLEMENTED", name+" is not yet implemented")
	}
}
