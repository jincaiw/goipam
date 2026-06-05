# GoIPAM

[English](README.md) | 中文

GoIPAM 是一个轻量级 Web IP 地址管理系统，面向需要统一管理 IPv4/IPv6 网段、地址、VLAN、VRF、设备、位置、NAT、DNS、扫描、申请审批、用户、角色和审计日志的团队。

项目后端使用 Go，前端使用 Vue 3。发布版二进制会内嵌前端资源和数据库迁移文件，因此 Linux 生产部署只需要一个可执行文件和一个 YAML 配置文件。

## 功能特性

- IPv4 / IPv6 网段和地址管理。
- Section、Folder、Subnet 层级，首个空闲 IP、重叠检查、网段拆分、可用范围查询。
- VLAN、VRF、设备、机柜/位置、NAT、DNS/Nameserver、扫描、申请、用户、角色、用户组、设置、搜索、仪表盘和审计 API。
- 本地账号密码认证，支持 JWT Access Token 和 Refresh Token。
- 默认 SQLite，支持 MySQL 8.0+。
- 生产二进制内嵌 Vue 3 前端。
- `/healthz`、`/readyz` 健康检查和 `/metrics` Prometheus 指标。
- Docker 与 systemd 部署示例。

## 快速开始

从最新 Release 下载 Linux 二进制：

```bash
curl -L -o goipam https://github.com/jincaiw/goipam/releases/download/v0.1.0/goipam-linux-amd64
chmod +x goipam
```

创建最小配置：

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

启动：

```bash
export GOIPAM_ADMIN_PASSWORD='replace-with-a-strong-admin-password'
./goipam --config config.yaml
```

打开 `http://localhost:5050`，使用以下账号登录：

- 用户名：`admin`
- 密码：`GOIPAM_ADMIN_PASSWORD` 的值

仅本地开发时，如果 `GOIPAM_ENV` 不是 `production` 且未设置 `GOIPAM_ADMIN_PASSWORD`，初始密码会回退为 `admin`。

## Linux 单文件部署

发布版二进制内嵌前端和迁移文件。生产主机只需要：

- `/opt/goipam/goipam`
- `/opt/goipam/config.yaml`
- `/etc/goipam/goipam.env`
- 一个可写 SQLite 数据目录，例如 `/var/lib/goipam`

安装：

```bash
sudo useradd --system --home /var/lib/goipam --shell /usr/sbin/nologin goipam
sudo install -d -o goipam -g goipam -m 0750 /opt/goipam /var/lib/goipam /etc/goipam
sudo install -m 0755 goipam /opt/goipam/goipam
```

创建 `/opt/goipam/config.yaml`：

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

创建 `/etc/goipam/goipam.env`：

```bash
GOIPAM_ENV=production
GOIPAM_ADMIN_PASSWORD=replace-with-a-strong-initial-password
```

保护密钥文件并设置数据目录权限：

```bash
sudo chown root:goipam /etc/goipam/goipam.env
sudo chmod 0640 /etc/goipam/goipam.env
sudo chown -R goipam:goipam /var/lib/goipam
```

## systemd 服务

创建 `/etc/systemd/system/goipam.service`：

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

启用并启动：

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now goipam
sudo systemctl status goipam
journalctl -u goipam -f
```

健康检查：

```bash
curl http://127.0.0.1:5050/healthz
curl http://127.0.0.1:5050/readyz
curl http://127.0.0.1:5050/metrics
```

## 从源码构建

要求：

- Go 1.26+
- Node.js 20+
- npm

构建本机二进制：

```bash
make build VERSION=v0.1.0
./bin/goipam --version
```

构建 Linux 发布资产：

```bash
VERSION=v0.1.0 ./scripts/build-linux.sh
ls -lh dist/
cat dist/SHA256SUMS
```

运行测试：

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

## 配置

GoIPAM 会读取 `configs/config.example.yaml`、配置文件和环境变量覆盖：

- `GOIPAM_LISTEN`
- `GOIPAM_PUBLIC_URL`
- `GOIPAM_DB_DRIVER`
- `GOIPAM_DB_DSN`
- `GOIPAM_JWT_SECRET`
- `GOIPAM_LOG_LEVEL`
- `GOIPAM_ENV`
- `GOIPAM_ADMIN_PASSWORD`

生产环境中，`security.jwt_secret` 不能使用默认值，首次启动必须设置 `GOIPAM_ADMIN_PASSWORD`。

## 开源协议

GoIPAM 使用 [MIT License](LICENSE) 开源。
