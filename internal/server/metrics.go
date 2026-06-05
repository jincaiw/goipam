package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// setupMetrics configures Prometheus metrics endpoint.
func (s *Server) setupMetrics(r chi.Router) {
	r.Handle("/metrics", promhttp.Handler())
}
