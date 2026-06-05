.PHONY: build build-linux release-assets run dev test lint clean migrate-up migrate-down sqlc-gen frontend-build

# Go parameters
BINARY_NAME=goipam
GO=go
MAIN=./cmd/goipam
VERSION ?= dev
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
DATE ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS=-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)

# Build
build: frontend-build
	rm -rf internal/server/web/dist
	cp -R frontend/dist internal/server/web/dist
	$(GO) build -ldflags "$(LDFLAGS)" -o bin/$(BINARY_NAME) $(MAIN)

build-linux: frontend-build
	rm -rf internal/server/web/dist
	cp -R frontend/dist internal/server/web/dist
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -trimpath -ldflags "$(LDFLAGS)" -o dist/$(BINARY_NAME)-linux-amd64 $(MAIN)

release-assets: build-linux
	cd dist && shasum -a 256 $(BINARY_NAME)-linux-amd64 > SHA256SUMS

# Run in development mode
run: build
	./bin/$(BINARY_NAME) --config configs/config.example.yaml

# Development with auto-reload (requires air)
dev:
	air -c .air.toml

# Tests
test:
	$(GO) test ./... -v -count=1

test-short:
	$(GO) test ./... -short -count=1

# Lint
lint:
	golangci-lint run ./...

# Clean
clean:
	rm -rf bin/ data/

# Database migrations
migrate-up:
	goose -dir migrations/sqlite sqlite3 ./data/goipam.db up

migrate-down:
	goose -dir migrations/sqlite sqlite3 ./data/goipam.db down

migrate-up-mysql:
	goose -dir migrations/mysql mysql "root:password@tcp(localhost:3306)/goipam?parseTime=true" up

migrate-down-mysql:
	goose -dir migrations/mysql mysql "root:password@tcp(localhost:3306)/goipam?parseTime=true" down

# SQLC code generation
sqlc-gen:
	sqlc generate -f sqlc/sqlc.yaml

# Frontend
frontend-install:
	cd frontend && npm install

frontend-dev:
	cd frontend && npm run dev

frontend-build:
	cd frontend && npm run build

# Docker
docker-build:
	docker build -t goipam:latest -f deploy/Dockerfile .

docker-up:
	docker compose -f deploy/docker-compose.yaml up -d

docker-down:
	docker compose -f deploy/docker-compose.yaml down
