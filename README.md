# GoIPAM

English | [中文](README.zh-CN.md)

GoIPAM is a lightweight web IP address management system for teams that need a single source of truth for IPv4/IPv6 subnets, addresses, VLANs, VRFs, devices, locations, NAT mappings, DNS records, scans, requests, users, roles, and audit logs.

The project is built with a Go backend and a Vue 3 frontend. Release binaries embed the frontend and database migrations, so a Linux deployment can run from one executable file plus a small YAML configuration file.

## Features

- IPv4 and IPv6 subnet/address management.
- Section, folder, subnet hierarchy, first-free IP, overlap checks, split and available-range helpers.
- VLAN, VRF, device, rack/location, NAT, DNS/nameserver, scan, request, user, role, group, settings, search, dashboard, and audit APIs.
- Local username/password authentication with JWT access and refresh tokens.
- SQLite by default, with MySQL 8.0+ support.
- Embedded Vue 3 UI for production single-binary deployment.
- Health checks at `/healthz` and `/readyz`, plus Prometheus metrics at `/metrics`.
- Docker and systemd deployment examples.

## Quick Start

Download the Linux binary from the latest release:

```bash
curl -L -o goipam https://github.com/jincaiw/goipam/releases/download/v0.1.0/goipam-linux-amd64
chmod +x goipam
```

Create a minimal config:

```bash
cat > config.yaml <<'EOF'
server:
  listen: "0.0.0.0:5050"
  public_url: "http://localhost:5050"

database:
  driver: "sqlite"
  dsn: "./data/goipam.db"

security:
  jwt_secret: "replace-with-a-long-random-secret"
  access_token_ttl: "30m"
  refresh_token_ttl: "168h"

scan:
  enabled: true
  max_concurrency: 128
  timeout: "1s"
  schedule: "0 */6 * * *"

log:
  level: "info"
  format: "json"
EOF
```

Start GoIPAM:

```bash
export GOIPAM_ADMIN_PASSWORD='replace-with-a-strong-admin-password'
./goipam --config config.yaml
```

Open `http://localhost:5050` and sign in with:

- Username: `admin`
- Password: the value of `GOIPAM_ADMIN_PASSWORD`

For local development only, if `GOIPAM_ENV` is not `production` and `GOIPAM_ADMIN_PASSWORD` is omitted, the initial password falls back to `admin`.

## Linux Single-File Deployment

The release binary embeds the frontend and migrations. A production host only needs:

- `/opt/goipam/goipam`
- `/opt/goipam/config.yaml`
- `/etc/goipam/goipam.env`
- A writable SQLite data directory, for example `/var/lib/goipam`

Install:

```bash
sudo useradd --system --home /var/lib/goipam --shell /usr/sbin/nologin goipam
sudo install -d -o goipam -g goipam -m 0750 /opt/goipam /var/lib/goipam /etc/goipam
sudo install -m 0755 goipam /opt/goipam/goipam
```

Create `/opt/goipam/config.yaml`:

```yaml
server:
  listen: "0.0.0.0:5050"
  public_url: "https://ipam.example.com"

database:
  driver: "sqlite"
  dsn: "/var/lib/goipam/goipam.db"

security:
  jwt_secret: "replace-with-at-least-32-random-bytes"
  access_token_ttl: "30m"
  refresh_token_ttl: "168h"

scan:
  enabled: true
  max_concurrency: 128
  timeout: "1s"
  schedule: "0 */6 * * *"

log:
  level: "info"
  format: "json"
```

Create `/etc/goipam/goipam.env`:

```bash
GOIPAM_ENV=production
GOIPAM_ADMIN_PASSWORD=replace-with-a-strong-initial-password
```

Protect secrets:

```bash
sudo chown root:goipam /etc/goipam/goipam.env
sudo chmod 0640 /etc/goipam/goipam.env
sudo chown -R goipam:goipam /var/lib/goipam
```

## systemd Service

Create `/etc/systemd/system/goipam.service`:

```ini
[Unit]
Description=GoIPAM - IP Address Management
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
User=goipam
Group=goipam
WorkingDirectory=/opt/goipam
EnvironmentFile=/etc/goipam/goipam.env
ExecStart=/opt/goipam/goipam --config /opt/goipam/config.yaml
Restart=on-failure
RestartSec=5
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
```

Enable and start:

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now goipam
sudo systemctl status goipam
journalctl -u goipam -f
```

Health checks:

```bash
curl http://127.0.0.1:5050/healthz
curl http://127.0.0.1:5050/readyz
curl http://127.0.0.1:5050/metrics
```

## Build From Source

Requirements:

- Go 1.26+
- Node.js 20+
- npm

Build the local binary:

```bash
make build VERSION=v0.1.0
./bin/goipam --version
```

Build Linux release assets:

```bash
VERSION=v0.1.0 ./scripts/build-linux.sh
ls -lh dist/
cat dist/SHA256SUMS
```

Run tests:

```bash
go test ./... -count=1
go vet ./...
cd frontend && npm run build
```

## Docker Compose

```bash
export GOIPAM_JWT_SECRET='replace-with-a-long-random-secret'
export GOIPAM_ADMIN_PASSWORD='replace-with-a-strong-admin-password'
docker compose -f deploy/docker-compose.yaml up -d
```

## Configuration

GoIPAM reads `configs/config.example.yaml` by default values plus optional config files and environment overrides:

- `GOIPAM_LISTEN`
- `GOIPAM_PUBLIC_URL`
- `GOIPAM_DB_DRIVER`
- `GOIPAM_DB_DSN`
- `GOIPAM_JWT_SECRET`
- `GOIPAM_LOG_LEVEL`
- `GOIPAM_ENV`
- `GOIPAM_ADMIN_PASSWORD`

In production, `security.jwt_secret` must not be the default value, and the first startup requires `GOIPAM_ADMIN_PASSWORD`.

## License

GoIPAM is released under the [MIT License](LICENSE).
