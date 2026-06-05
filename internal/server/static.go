package server

import (
	"bytes"
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

// embeddedFrontend stores the production SPA in the Go binary for single-file
// deployments. The release build script refreshes this directory from frontend/dist.
//
//go:embed web/dist/*
var embeddedFrontend embed.FS

// setupStaticFiles serves the frontend static files.
// In development, the frontend runs on a separate Vite dev server.
// In production, the built frontend is served from the filesystem when present,
// otherwise the embedded copy is used so the Linux binary is self-contained.
func (s *Server) setupStaticFiles() {
	distDir := "frontend/dist"

	assets, err := staticAssets(distDir)
	if err != nil {
		return
	}
	fileServer := http.FileServer(http.FS(assets))

	// Use chi's wildcard to catch all non-API routes
	s.mux.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqPath := r.URL.Path

		// Don't interfere with API routes
		if strings.HasPrefix(reqPath, "/api/") {
			http.NotFound(w, r)
			return
		}
		if reqPath == "/healthz" || reqPath == "/readyz" || strings.HasPrefix(reqPath, "/metrics") {
			http.NotFound(w, r)
			return
		}

		// Try serving the file, fall back to index.html for SPA routing
		name := strings.TrimPrefix(path.Clean(reqPath), "/")
		if name == "." || name == "" {
			name = "index.html"
		}
		if _, err := fs.Stat(assets, name); err != nil {
			serveStaticFile(w, r, assets, "index.html")
		} else {
			fileServer.ServeHTTP(w, r)
		}
	}))
}

func staticAssets(distDir string) (fs.FS, error) {
	if _, err := os.Stat(distDir); err == nil {
		return os.DirFS(distDir), nil
	}
	return fs.Sub(embeddedFrontend, "web/dist")
}

func serveStaticFile(w http.ResponseWriter, r *http.Request, assets fs.FS, name string) {
	content, err := fs.ReadFile(assets, name)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.ServeContent(w, r, name, time.Time{}, bytes.NewReader(content))
}
