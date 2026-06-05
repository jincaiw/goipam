package db

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSQLitePathFromDSN(t *testing.T) {
	tests := []struct {
		name string
		dsn  string
		want string
	}{
		{name: "plain path", dsn: "./data/goipam.db", want: "./data/goipam.db"},
		{name: "path with query", dsn: "/var/lib/goipam/goipam.db?_pragma=busy_timeout(5000)", want: "/var/lib/goipam/goipam.db"},
		{name: "file URI", dsn: "file:/var/lib/goipam/goipam.db?cache=shared", want: "/var/lib/goipam/goipam.db"},
		{name: "memory", dsn: ":memory:", want: ""},
		{name: "file memory", dsn: "file::memory:?cache=shared", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sqlitePathFromDSN(tt.dsn); got != tt.want {
				t.Fatalf("sqlitePathFromDSN(%q) = %q, want %q", tt.dsn, got, tt.want)
			}
		})
	}
}

func TestEnsureSQLiteDirCreatesParentDirectory(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "nested", "data")
	dsn := filepath.Join(dir, "goipam.db")

	if err := ensureSQLiteDir(dsn); err != nil {
		t.Fatalf("ensureSQLiteDir() error = %v", err)
	}

	if !dirExists(dir) {
		t.Fatalf("expected %q to exist", dir)
	}
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}
