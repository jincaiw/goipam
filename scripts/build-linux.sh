#!/usr/bin/env bash
set -euo pipefail

VERSION="${VERSION:-v0.1.0}"
COMMIT="${COMMIT:-$(git rev-parse --short HEAD 2>/dev/null || echo unknown)}"
DATE="${DATE:-$(date -u +%Y-%m-%dT%H:%M:%SZ)}"

mkdir -p dist internal/server/web

cd frontend
npm ci
npm run build
cd ..

rm -rf internal/server/web/dist
cp -R frontend/dist internal/server/web/dist

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
  -trimpath \
  -ldflags "-s -w -X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE}" \
  -o dist/goipam-linux-amd64 \
  ./cmd/goipam

(
  cd dist
  shasum -a 256 goipam-linux-amd64 > SHA256SUMS
)
