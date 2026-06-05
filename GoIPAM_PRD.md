# Go IPAM 产品需求文档

| 项目 | 内容 |
|---|---|
| 产品名称 | GoIPAM |
| 文档版本 | 1.0 |
| 文档日期 | 2026-06-03 |
| 建设方式 | 使用 Go + Vue 3 + TypeScript + Naive UI 重新设计 |
| 目标形态 | 前后端分离 Web 系统，后端 Go 单体优先，默认 SQLite，支持 MySQL 8.0+ |
| 适用对象 | 产品、研发、测试、运维、安全、交付、二次开发团队 |

---

## 1. 文档目的

本文档用于指导将核心 Web IPAM 功能使用 Go 语言重新设计和实现。本文档所定义的目标不是逐行翻译旧代码，而是在保留成熟 IPAM 产品能力的基础上，形成一套更适合生产部署、二次开发、API 集成和现代前端体验的新系统需求基线。

本版本仅覆盖以下核心范围：

1. IPv4 / IPv6 地址管理。
2. Section / Folder / Subnet 管理。
3. IP 地址管理。
4. VLAN / L2 Domain 管理。
5. VRF 管理。
6. 设备、设备类型管理。
7. Rack / Location 管理。
8. NAT 管理。
9. DNS / Nameserver 集成。
10. 扫描、探活、新主机发现。
11. IP 申请与审批。
12. 用户、角色、用户组、权限。
13. REST API。
14. 日志、审计、安全、运维部署。

---

## 2. 产品定位

GoIPAM 是面向企业、运营商、数据中心、园区网络、云平台和运维团队的轻量级 Web IP 地址管理系统。系统以 IP 地址事实库为核心，统一管理网段、地址、VLAN、VRF、设备、机柜、位置、NAT、DNS、扫描结果、申请审批、权限和开放 API。

系统重点解决：

1. IP 地址台账分散在 Excel、Wiki、个人文档中，缺乏统一事实源。
2. 网段层级、可用空间、重叠冲突不可视。
3. IP 分配、回收、申请审批没有流程。
4. VLAN、VRF、设备、位置和地址之间缺少关联。
5. NAT 映射关系难以追踪。
6. 地址实际在线状态与台账状态不一致。
7. 外部系统缺少稳定 API 接入 IPAM 数据。

---

## 3. 建设目标

### 3.1 产品目标

1. 建立统一 IP 地址事实库，支持 IPv4 与 IPv6。
2. 支持多 Section、多层级 Subnet、Folder 和 Supernet 管理。
3. 支持 IP 地址全生命周期：申请、分配、占用、保留、回收、删除、审计。
4. 支持 VLAN、VRF、设备、机柜、位置、NAT、DNS 的关联管理。
5. 支持自动扫描、探活、新主机发现、地址状态校验。
6. 支持细粒度权限控制，最小可控到 Section / Subnet / 操作级。
7. 提供 REST API，便于 CMDB、自动化平台、监控系统、工单系统集成。
8. 支持默认 SQLite 单文件部署，同时支持 MySQL 8.0+ 生产部署。
9. 支持 Go 后端单二进制运行，前端静态资源可嵌入 Go 二进制。

### 3.2 技术目标

1. 后端采用 Go 1.22+。
2. Web 框架推荐 Chi，保持轻量、清晰、标准库友好。
3. 数据库默认 SQLite，支持 MySQL 8.0+。
4. 数据访问采用 SQLC，迁移采用 Goose。
5. 前端采用 Vue 3 + TypeScript + Naive UI + Vite。
6. 前后端严格以 REST API 为边界。
7. 生产运行不依赖 Node.js。
8. 支持 Linux 单文件部署，systemd 管理。
9. 支持 Docker / Docker Compose 部署。
10. 支持结构化日志、健康检查、Prometheus 指标。

---

## 4. 总体架构

### 4.1 架构形态

```text
Browser
  ↓ HTTPS
Vue 3 + Naive UI SPA
  ↓ REST API /api/v1
Go API Server
  ├── Auth / RBAC
  ├── IPAM Domain Service
  ├── Scan Scheduler
  ├── DNS Integration
  ├── Audit / Log
  ├── Import / Export
  └── Metrics / Health
  ↓ SQLC Repository
SQLite / MySQL
```

### 4.2 后端分层

```text
cmd/goipam/                  程序入口
internal/config/             配置加载
internal/server/             HTTP Server、Router、中间件
internal/auth/               登录、JWT、Session、密码、Token
internal/rbac/               权限、角色、资源授权
internal/domain/ipam/        Section、Subnet、Address 核心领域
internal/domain/network/     VLAN、VRF、NAT、Device、Rack、Location
internal/domain/dns/         Nameserver、DNS 解析、PTR 记录
internal/domain/scan/        Ping、TCP、发现、调度任务
internal/domain/request/     IP 申请审批
internal/repository/         Repository 接口封装
internal/db/                 SQLC 生成代码
internal/audit/              操作日志、审计日志
internal/api/                REST Handler、DTO、校验
internal/job/                后台任务调度
internal/exporter/           CSV/XLSX 导入导出
frontend/                    Vue 3 前端工程
migrations/sqlite/           SQLite Goose 迁移
migrations/mysql/            MySQL Goose 迁移
configs/                     示例配置
scripts/                     systemd、安装、备份脚本
```

### 4.3 部署形态

| 部署方式 | 说明 | 适用场景 |
|---|---|---|
| 单二进制 + SQLite | Go 二进制内嵌前端资源，SQLite 本地文件 | 小团队、单机、快速落地 |
| 单二进制 + MySQL | Go 二进制连接外部 MySQL | 企业生产、多人协作 |
| Docker Compose | app + mysql 可选 | 标准化交付、测试环境 |
| systemd | Linux 服务守护 | 生产运行 |

---

## 5. 用户角色

| 角色 | 说明 | 核心权限 |
|---|---|---|
| 超级管理员 | 系统最高权限 | 所有功能、系统设置、用户权限、数据库维护 |
| IPAM 管理员 | 网络/IPAM 负责人 | 管理所有 Section、Subnet、Address、VLAN、VRF、NAT、设备 |
| Section 管理员 | 某区域/部门管理员 | 管理授权 Section 下资源 |
| Subnet 管理员 | 某网段管理员 | 管理授权 Subnet 下地址和申请 |
| 普通用户 | 使用者/申请人 | 查看授权资源，提交 IP 申请 |
| 只读用户 | 审计或查看人员 | 只读查看授权资源 |
| API 调用方 | 外部系统 | 按 API Token 权限访问 |

---

## 6. 功能模块需求

## 6.1 登录、认证与会话

### 功能需求

1. 支持本地账号密码登录。
2. 支持管理员创建、禁用、锁定用户。
3. 支持密码复杂度策略：最小长度、大小写、数字、特殊字符。
4. 支持密码哈希：Argon2id，保留 bcrypt 兼容能力。
5. 支持登录失败次数限制与临时锁定。
6. 支持 JWT Access Token + Refresh Token。
7. 支持 Web Session 模式，可配置 Token 过期时间。
8. 支持 API Token 独立管理。
9. 支持退出登录、刷新 Token、修改密码、重置密码。
10. 支持单用户强制下线。
11. 支持操作审计记录登录、退出、失败登录、密码变更。

### 不纳入首版

LDAP、AD、Radius、SAML、Passkeys、MFA 可作为后续扩展，本版本优先完成本地认证闭环。

### 验收标准

1. 用户可登录、退出、刷新会话。
2. 禁用用户不可登录。
3. 连续失败登录达到阈值后账号被临时锁定。
4. 所有受保护 API 未登录时返回 401。
5. 登录、退出、失败登录均写入审计日志。

---

## 6.2 仪表盘

### 功能需求

1. 展示 Section 数量、Subnet 数量、IP 地址数量、已用地址、空闲地址。
2. 展示 IPv4 / IPv6 使用率。
3. 展示最近变更、最近申请、最近扫描任务。
4. 展示使用率最高的 Top N Subnet。
5. 展示即将耗尽的 Subnet，按阈值预警。
6. 展示待审批 IP 申请数量。
7. 展示扫描发现的新主机数量。
8. 按用户权限过滤数据范围。

### 验收标准

1. 超级管理员看到全局统计。
2. 普通用户只看到授权范围内统计。
3. 使用率和容量统计与 Subnet / Address 数据一致。

---

## 6.3 Section 管理

### 功能说明

Section 是最高级地址空间逻辑分区，可按组织、区域、数据中心、环境、业务系统、租户或网络域划分。

### 字段需求

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint | 主键 |
| name | varchar(128) | Section 名称，唯一 |
| description | text | 描述 |
| parent_id | bigint/null | 父级 Section，可选 |
| strict_mode | bool | 是否启用严格模式 |
| show_vlan | bool | 是否展示 VLAN |
| show_vrf | bool | 是否展示 VRF |
| subnet_order | varchar(32) | 子网排序方式 |
| sort_order | int | 展示排序 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

### 功能需求

1. 支持新增、编辑、删除 Section。
2. 支持 Section 嵌套。
3. 支持 Section 下挂载 Folder / Subnet。
4. 支持 Section 级权限授权。
5. 支持严格模式：限制非法重叠、父子网段关系错误。
6. 支持 Section 统计：Subnet 数量、地址容量、使用率。
7. 支持按名称、描述搜索。
8. 删除 Section 前必须检测下属资源。

### 验收标准

1. 创建 Section 后可在菜单、列表、API 中查询。
2. 有子资源的 Section 不允许直接删除，除非显式选择级联删除。
3. 无权限用户不可看到该 Section。

---

## 6.4 Folder 管理

### 功能说明

Folder 是 Subnet 树中的目录节点，只用于组织网段，不参与 IP 地址分配。

### 字段需求

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint | 主键 |
| section_id | bigint | 所属 Section |
| parent_id | bigint/null | 父 Folder 或父 Subnet |
| name | varchar(128) | 名称 |
| description | text | 描述 |
| sort_order | int | 排序 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

### 功能需求

1. 支持创建、编辑、删除 Folder。
2. 支持 Folder 嵌套。
3. 支持 Folder 下挂载 Subnet。
4. 支持权限继承 Section，也支持独立授权。
5. 支持拖拽或选择父级调整层级。

### 验收标准

1. Folder 不允许创建 IP Address。
2. Folder 下可展示子 Folder 和子 Subnet。
3. 删除 Folder 前必须确认其下资源处理方式。

---

## 6.5 Subnet 管理

### 功能说明

Subnet 是 IPAM 核心对象，用于管理 IPv4/IPv6 网段、掩码、层级、容量、使用率、扫描、DNS、VLAN、VRF、位置和权限。

### 字段需求

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint | 主键 |
| section_id | bigint | 所属 Section |
| parent_id | bigint/null | 父 Subnet 或 Folder |
| cidr | varchar(64) | CIDR，如 192.168.1.0/24 |
| ip_version | tinyint | 4 或 6 |
| network | varbinary(16) | 网络地址二进制，便于排序和比较 |
| prefix_len | tinyint | 掩码长度 |
| name | varchar(128) | 名称 |
| description | text | 描述 |
| vlan_id | bigint/null | 关联 VLAN |
| vrf_id | bigint/null | 关联 VRF |
| location_id | bigint/null | 关联位置 |
| nameserver_id | bigint/null | 关联 DNS |
| scan_enabled | bool | 是否探活 |
| discover_enabled | bool | 是否新主机发现 |
| threshold_percent | int | 使用率阈值 |
| allow_request | bool | 是否允许 IP 申请 |
| status | varchar(32) | active/reserved/deprecated |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

### 功能需求

1. 支持 IPv4 / IPv6 Subnet 创建。
2. 支持父子网段层级。
3. 支持 Supernet / Subnet 树形展示。
4. 支持 CIDR 合法性校验。
5. 支持重叠网段检测。
6. 支持 VRF 维度隔离：不同 VRF 可允许相同 CIDR。
7. 支持查询第一个可用 IP。
8. 支持查询所有可用 IP 范围。
9. 支持查询第一个可用子网。
10. 支持子网拆分，例如 /24 拆为多个 /28。
11. 支持调整掩码，但必须校验已有地址是否仍在范围内。
12. 支持统计总地址数、已用数、可用数、保留数、使用率。
13. 支持阈值预警。
14. 支持绑定 VLAN、VRF、Location、Nameserver。
15. 支持启用/禁用扫描。
16. 支持启用/禁用新主机发现。
17. 支持启用/禁用 IP 申请。
18. 支持 Subnet 权限配置。
19. 支持可视化地址占用图。
20. 支持 CSV/XLSX 导入导出。

### 业务规则

1. Subnet 必须属于 Section。
2. 子 Subnet 必须在父 Subnet 范围内。
3. 严格模式下，同一 Section + VRF 内不允许非法重叠。
4. 默认不允许创建与已有 Subnet 完全相同的 CIDR。
5. 删除 Subnet 前必须检测子网段和地址。
6. 有地址的 Subnet 不允许直接删除，除非选择级联删除。

### 验收标准

1. IPv4 / IPv6 CIDR 校验正确。
2. 使用率计算正确。
3. 第一个可用 IP 返回正确。
4. 重叠网段检测符合 Section + VRF 规则。
5. API 和 Web 操作结果一致。

---

## 6.6 IP 地址管理

### 字段需求

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint | 主键 |
| subnet_id | bigint | 所属 Subnet |
| ip | varbinary(16) | IP 二进制 |
| ip_text | varchar(64) | IP 字符串 |
| hostname | varchar(255) | 主机名 |
| mac | varchar(32) | MAC 地址 |
| description | text | 描述 |
| status | varchar(32) | used/free/reserved/offline/deprecated |
| tag_id | bigint/null | 地址标签 |
| is_gateway | bool | 是否网关 |
| device_id | bigint/null | 关联设备 |
| device_port | varchar(64) | 设备端口 |
| owner | varchar(128) | 负责人或部门 |
| last_seen_at | datetime/null | 最后探活时间 |
| dns_checked_at | datetime/null | DNS 检查时间 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

### 功能需求

1. 支持新增、编辑、删除 IP 地址。
2. 支持在指定 Subnet 内创建第一个可用 IP。
3. 支持手动指定 IP 创建。
4. 支持保留地址、网关地址、废弃地址标记。
5. 支持主机名、MAC、描述、负责人维护。
6. 支持关联设备和端口。
7. 支持关联 NAT 映射。
8. 支持单个 IP 探活。
9. 支持 DNS 正向/反向解析展示。
10. 支持地址变更审计。
11. 支持批量导入、导出。
12. 支持按 IP、主机名、MAC、描述、设备搜索。
13. 支持地址状态自动更新：在线、离线、未知。

### 业务规则

1. IP 必须属于所属 Subnet。
2. 同一 Subnet 内同一 IP 默认唯一。
3. 不同 VRF 下允许相同 IP。
4. 网段地址、广播地址默认不可分配，可通过配置允许。
5. 删除 IP 后释放地址，但保留审计日志。

### 验收标准

1. 创建重复 IP 时按规则阻止。
2. 创建不属于 Subnet 的 IP 时返回错误。
3. 删除地址后可再次分配。
4. 搜索结果受权限控制。

---

## 6.7 VLAN / L2 Domain 管理

### VLAN 字段

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint | 主键 |
| l2_domain_id | bigint/null | 所属二层域 |
| vlan_id | int | VLAN 编号 1-4094 |
| name | varchar(128) | 名称 |
| description | text | 描述 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

### L2 Domain 字段

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint | 主键 |
| name | varchar(128) | 名称 |
| description | text | 描述 |

### 功能需求

1. 支持新增、编辑、删除 VLAN。
2. 支持 L2 Domain 管理。
3. 支持 VLAN 关联多个 Subnet。
4. 支持按 VLAN 查看关联 Subnet。
5. 支持 VLAN 导入导出。
6. 支持同一 L2 Domain 内 VLAN ID 唯一。
7. 支持跨 L2 Domain 复用 VLAN ID。

### 验收标准

1. VLAN ID 必须在 1-4094。
2. 删除 VLAN 前提示关联 Subnet。
3. VLAN 页面可查看使用该 VLAN 的所有 Subnet。

---

## 6.8 VRF 管理

### 字段需求

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint | 主键 |
| name | varchar(128) | VRF 名称 |
| rd | varchar(128) | Route Distinguisher |
| description | text | 描述 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

### 功能需求

1. 支持新增、编辑、删除 VRF。
2. 支持 VRF 与 Subnet 关联。
3. 支持按 VRF 过滤 Subnet 和 IP。
4. 支持同 CIDR 在不同 VRF 中复用。
5. 支持 VRF 使用率统计。

### 验收标准

1. 删除 VRF 前必须提示关联 Subnet。
2. 同一 VRF 内重叠规则受严格模式控制。
3. 不同 VRF 的相同地址空间不会互相冲突。

---

## 6.9 设备与设备类型管理

### 设备字段

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint | 主键 |
| name | varchar(128) | 设备名称 |
| hostname | varchar(255) | 主机名 |
| mgmt_ip | varchar(64) | 管理 IP |
| type_id | bigint/null | 设备类型 |
| location_id | bigint/null | 位置 |
| rack_id | bigint/null | 机柜 |
| rack_start_u | int/null | 起始 U 位 |
| rack_size_u | int/null | 占用 U 数 |
| manufacturer | varchar(128) | 厂商 |
| model | varchar(128) | 型号 |
| serial_number | varchar(128) | 序列号 |
| description | text | 描述 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

### 设备类型字段

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint | 主键 |
| name | varchar(128) | 类型名称，如 switch/router/server/firewall |
| icon | varchar(64) | 图标 |
| description | text | 描述 |

### 功能需求

1. 支持设备增删改查。
2. 支持设备类型管理。
3. 支持设备与 IP 地址关联。
4. 支持设备与 Location、Rack 关联。
5. 支持设备上架信息维护。
6. 支持设备搜索和筛选。
7. 支持设备详情页展示关联 IP、Subnet、NAT、Rack。
8. 支持设备导入导出。

### 不纳入首版

设备接口自动发现、SNMP 接口采集、配置备份、拓扑发现不纳入首版，仅保留字段和扩展接口。

---

## 6.10 Rack / Location 管理

### Location 字段

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint | 主键 |
| name | varchar(128) | 位置名称 |
| address | varchar(255) | 地址 |
| room | varchar(128) | 房间 |
| row_name | varchar(64) | 行 |
| description | text | 描述 |

### Rack 字段

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint | 主键 |
| location_id | bigint | 所属位置 |
| name | varchar(128) | 机柜名称 |
| size_u | int | U 位容量，如 42 |
| row_name | varchar(64) | 所在行 |
| position | varchar(64) | 位置编号 |
| direction | varchar(16) | top_down/bottom_up |
| description | text | 描述 |

### 功能需求

1. 支持 Location 增删改查。
2. 支持 Rack 增删改查。
3. 支持设备上架、下架。
4. 支持校验 U 位占用冲突。
5. 支持 Rack 可视化展示。
6. 支持按位置查看机柜、设备、Subnet。
7. 支持按机柜查看设备和管理 IP。

### 验收标准

1. 设备上架时不能与已有设备 U 位冲突。
2. Rack 容量修改时不能小于现有占用最大 U 位。
3. 删除 Location / Rack 前必须提示关联对象。

---

## 6.11 NAT 管理

### 字段需求

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint | 主键 |
| name | varchar(128) | NAT 名称 |
| type | varchar(32) | static/dynamic/policy/port_forward |
| inside_address_id | bigint/null | 内部地址对象 |
| outside_address_id | bigint/null | 外部地址对象 |
| inside_ip | varchar(64) | 内部 IP，兼容未建地址对象场景 |
| outside_ip | varchar(64) | 外部 IP |
| inside_port | int/null | 内部端口 |
| outside_port | int/null | 外部端口 |
| protocol | varchar(16) | tcp/udp/icmp/any |
| device_id | bigint/null | 关联设备/防火墙 |
| description | text | 描述 |
| enabled | bool | 是否启用 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

### 功能需求

1. 支持静态 NAT。
2. 支持动态 NAT 记录。
3. 支持端口映射。
4. 支持 Policy NAT 基础字段。
5. 支持 NAT 与 IP 地址对象关联。
6. 支持通过 IP 地址详情查看 NAT 关系。
7. 支持 NAT 与设备关联。
8. 支持 NAT 查询、筛选、导出。
9. 支持 NAT 变更审计。

### 验收标准

1. NAT 创建后可从内部 IP 和外部 IP 两侧查询。
2. 删除关联 IP 前提示 NAT 关系。
3. NAT 端口映射端口范围必须合法。

---

## 6.12 DNS / Nameserver 集成

### Nameserver 字段

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint | 主键 |
| name | varchar(128) | 名称 |
| type | varchar(32) | udp/tcp/http_api/powerdns |
| server | varchar(255) | 服务器地址 |
| port | int | 端口 |
| timeout_ms | int | 超时时间 |
| enabled | bool | 是否启用 |
| description | text | 描述 |

### 功能需求

1. 支持 Nameserver 配置。
2. 支持 Section / Subnet 绑定 Nameserver。
3. 支持地址详情页执行正向解析。
4. 支持地址详情页执行反向解析。
5. 支持 DNS 解析结果写入 IP hostname。
6. 支持 DNS 检查时间记录。
7. 支持 PowerDNS API 预留扩展字段。

### 首版边界

首版只要求 DNS 查询和 hostname 同步；PowerDNS Zone / Record 创建、删除、自动 PTR 管理作为后续增强。

---

## 6.13 扫描、探活与新主机发现

### 功能需求

1. 支持 Subnet 级启用 Ping 探活。
2. 支持 Subnet 级启用 TCP 探活。
3. 支持 Subnet 级启用新主机发现。
4. 支持手动扫描单个 Subnet。
5. 支持定时扫描任务。
6. 支持扫描并发数配置。
7. 支持扫描超时配置。
8. 支持扫描结果写入 last_seen_at。
9. 支持发现台账中不存在但在线的 IP。
10. 支持新发现 IP 进入“待确认”状态。
11. 支持离线地址标记，不自动删除。
12. 支持扫描日志和任务状态。

### 扫描任务字段

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint | 主键 |
| subnet_id | bigint | 扫描网段 |
| task_type | varchar(32) | ping/tcp/discover/dns |
| status | varchar(32) | pending/running/success/failed/canceled |
| started_at | datetime/null | 开始时间 |
| finished_at | datetime/null | 结束时间 |
| total_count | int | 总数量 |
| alive_count | int | 在线数量 |
| new_count | int | 新发现数量 |
| error | text | 错误信息 |

### 验收标准

1. 扫描任务不阻塞 Web 请求。
2. 大网段扫描必须受最大地址数和并发数限制。
3. 新发现 IP 不直接变为正式分配，需管理员确认。
4. 扫描失败有错误日志。

---

## 6.14 IP 申请与审批

### 申请字段

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint | 主键 |
| subnet_id | bigint | 申请网段 |
| applicant_id | bigint | 申请人 |
| requested_ip | varchar(64)/null | 指定申请 IP，可为空 |
| hostname | varchar(255) | 主机名 |
| description | text | 用途说明 |
| status | varchar(32) | pending/approved/rejected/canceled |
| approver_id | bigint/null | 审批人 |
| approve_comment | text | 审批意见 |
| created_at | datetime | 申请时间 |
| updated_at | datetime | 更新时间 |

### 功能需求

1. 普通用户可在允许申请的 Subnet 中提交申请。
2. 可指定 IP，也可申请系统自动分配。
3. 管理员可审批通过、拒绝、取消。
4. 审批通过后自动创建 IP 地址。
5. 审批拒绝必须填写原因。
6. 支持申请列表、筛选、导出。
7. 支持申请审计。
8. 支持邮件通知预留接口。

### 验收标准

1. 未开启申请的 Subnet 不允许提交申请。
2. 申请指定 IP 被占用时不可提交或审批失败。
3. 审批通过后 Address 数据正确生成。

---

## 6.15 搜索与查询

### 功能需求

1. 支持全局搜索。
2. 支持 IP 精确搜索。
3. 支持 CIDR 搜索。
4. 支持主机名模糊搜索。
5. 支持 MAC 搜索。
6. 支持设备名称搜索。
7. 支持 VLAN / VRF / Location 条件筛选。
8. 支持搜索结果权限过滤。
9. 支持 API 搜索。

### 验收标准

1. 无权限资源不出现在搜索结果中。
2. 搜索结果可跳转到详情页。
3. 搜索响应时间在合理范围内。

---

## 6.16 导入导出

### 功能需求

1. 支持 Section 导出。
2. 支持 Subnet CSV/XLSX 导入导出。
3. 支持 IP 地址 CSV/XLSX 导入导出。
4. 支持 VLAN / VRF / Device 导入导出。
5. 支持导入前预校验。
6. 支持导入错误行报告。
7. 支持导入任务审计。

### 验收标准

1. 合法数据可导入。
2. 非法数据必须返回行号、字段、错误原因。
3. 导入不会绕过权限和业务规则。

---

## 7. 权限模型

### 7.1 权限维度

权限由三层共同决定：

1. 系统角色权限：决定用户能使用哪些模块和操作。
2. 资源权限：决定用户能访问哪些 Section / Subnet。
3. API Token 权限：决定外部调用方可执行哪些 API。

### 7.2 资源权限级别

| 权限级别 | 说明 |
|---|---|
| none | 无权限 |
| read | 只读 |
| write | 可新增和编辑地址 |
| manage | 可管理 Subnet、权限、扫描、申请审批 |
| admin | 资源管理员 |

### 7.3 操作权限清单

| 权限码 | 说明 |
|---|---|
| section.read | 查看 Section |
| section.manage | 管理 Section |
| subnet.read | 查看 Subnet |
| subnet.write | 新增/编辑 Subnet |
| subnet.delete | 删除 Subnet |
| address.read | 查看地址 |
| address.write | 新增/编辑地址 |
| address.delete | 删除地址 |
| vlan.manage | 管理 VLAN |
| vrf.manage | 管理 VRF |
| device.manage | 管理设备 |
| rack.manage | 管理机柜 |
| nat.manage | 管理 NAT |
| dns.manage | 管理 DNS 配置 |
| scan.run | 执行扫描 |
| request.approve | 审批 IP 申请 |
| user.manage | 管理用户 |
| role.manage | 管理角色权限 |
| audit.read | 查看审计日志 |
| system.manage | 系统设置 |

### 7.4 验收标准

1. 前端根据权限隐藏菜单和按钮。
2. 后端必须再次校验权限。
3. API 不得因前端隐藏而放松后端校验。
4. 权限变更后用户重新请求 API 即生效。

---

## 8. REST API 设计

### 8.1 API 基础规范

1. API 前缀：`/api/v1`。
2. 请求格式：JSON。
3. 响应格式：JSON。
4. 认证：`Authorization: Bearer <token>`。
5. 分页参数：`page`、`page_size`。
6. 排序参数：`sort`、`order`。
7. 过滤参数：按资源定义。
8. 错误响应统一包含 `code`、`message`、`request_id`。

### 8.2 统一响应结构

```json
{
  "success": true,
  "data": {},
  "request_id": "req_xxx"
}
```

失败：

```json
{
  "success": false,
  "code": "SUBNET_OVERLAP",
  "message": "subnet overlaps with existing subnet",
  "request_id": "req_xxx"
}
```

### 8.3 API 清单

#### 认证

| 方法 | 路径 | 说明 |
|---|---|---|
| POST | `/api/v1/auth/login` | 登录 |
| POST | `/api/v1/auth/logout` | 退出 |
| POST | `/api/v1/auth/refresh` | 刷新 Token |
| GET | `/api/v1/auth/me` | 当前用户信息 |
| POST | `/api/v1/auth/change-password` | 修改密码 |

#### Section

| 方法 | 路径 | 说明 |
|---|---|---|
| GET | `/api/v1/sections` | Section 列表 |
| POST | `/api/v1/sections` | 创建 Section |
| GET | `/api/v1/sections/{id}` | Section 详情 |
| PUT | `/api/v1/sections/{id}` | 更新 Section |
| DELETE | `/api/v1/sections/{id}` | 删除 Section |
| GET | `/api/v1/sections/{id}/tree` | Section 下资源树 |
| GET | `/api/v1/sections/{id}/stats` | Section 统计 |

#### Folder

| 方法 | 路径 | 说明 |
|---|---|---|
| GET | `/api/v1/folders` | Folder 列表 |
| POST | `/api/v1/folders` | 创建 Folder |
| GET | `/api/v1/folders/{id}` | Folder 详情 |
| PUT | `/api/v1/folders/{id}` | 更新 Folder |
| DELETE | `/api/v1/folders/{id}` | 删除 Folder |

#### Subnet

| 方法 | 路径 | 说明 |
|---|---|---|
| GET | `/api/v1/subnets` | Subnet 列表 |
| POST | `/api/v1/subnets` | 创建 Subnet |
| GET | `/api/v1/subnets/{id}` | Subnet 详情 |
| PUT | `/api/v1/subnets/{id}` | 更新 Subnet |
| DELETE | `/api/v1/subnets/{id}` | 删除 Subnet |
| GET | `/api/v1/subnets/{id}/children` | 子网段 |
| GET | `/api/v1/subnets/{id}/addresses` | 地址列表 |
| GET | `/api/v1/subnets/{id}/first-free-ip` | 第一个可用 IP |
| POST | `/api/v1/subnets/{id}/split` | 子网拆分 |
| GET | `/api/v1/subnets/{id}/available-ranges` | 可用地址范围 |
| GET | `/api/v1/subnets/{id}/stats` | 使用率统计 |
| POST | `/api/v1/subnets/check-overlap` | 重叠检测 |

#### Address

| 方法 | 路径 | 说明 |
|---|---|---|
| GET | `/api/v1/addresses` | 地址列表 |
| POST | `/api/v1/addresses` | 创建地址 |
| POST | `/api/v1/addresses/first-free` | 创建第一个可用地址 |
| GET | `/api/v1/addresses/{id}` | 地址详情 |
| PUT | `/api/v1/addresses/{id}` | 更新地址 |
| DELETE | `/api/v1/addresses/{id}` | 删除地址 |
| POST | `/api/v1/addresses/{id}/ping` | 单地址探活 |
| POST | `/api/v1/addresses/{id}/dns-resolve` | DNS 解析 |

#### VLAN / VRF / Device / Rack / NAT

| 方法 | 路径 | 说明 |
|---|---|---|
| GET/POST | `/api/v1/vlans` | VLAN 列表/创建 |
| GET/PUT/DELETE | `/api/v1/vlans/{id}` | VLAN 详情/更新/删除 |
| GET/POST | `/api/v1/l2-domains` | L2 Domain 列表/创建 |
| GET/POST | `/api/v1/vrfs` | VRF 列表/创建 |
| GET/PUT/DELETE | `/api/v1/vrfs/{id}` | VRF 详情/更新/删除 |
| GET/POST | `/api/v1/devices` | 设备列表/创建 |
| GET/PUT/DELETE | `/api/v1/devices/{id}` | 设备详情/更新/删除 |
| GET/POST | `/api/v1/locations` | 位置列表/创建 |
| GET/POST | `/api/v1/racks` | 机柜列表/创建 |
| GET/PUT/DELETE | `/api/v1/racks/{id}` | 机柜详情/更新/删除 |
| POST | `/api/v1/racks/{id}/mount-device` | 设备上架 |
| POST | `/api/v1/racks/{id}/unmount-device` | 设备下架 |
| GET/POST | `/api/v1/nats` | NAT 列表/创建 |
| GET/PUT/DELETE | `/api/v1/nats/{id}` | NAT 详情/更新/删除 |

#### 扫描与申请

| 方法 | 路径 | 说明 |
|---|---|---|
| POST | `/api/v1/scans/subnets/{id}` | 扫描指定 Subnet |
| GET | `/api/v1/scans/tasks` | 扫描任务列表 |
| GET | `/api/v1/scans/tasks/{id}` | 扫描任务详情 |
| POST | `/api/v1/scans/tasks/{id}/cancel` | 取消扫描 |
| GET | `/api/v1/requests` | IP 申请列表 |
| POST | `/api/v1/requests` | 提交 IP 申请 |
| POST | `/api/v1/requests/{id}/approve` | 审批通过 |
| POST | `/api/v1/requests/{id}/reject` | 拒绝申请 |

#### 系统管理

| 方法 | 路径 | 说明 |
|---|---|---|
| GET | `/api/v1/users` | 用户列表 |
| POST | `/api/v1/users` | 创建用户 |
| GET/PUT/DELETE | `/api/v1/users/{id}` | 用户详情/更新/删除 |
| GET/POST | `/api/v1/roles` | 角色列表/创建 |
| GET/POST | `/api/v1/groups` | 用户组列表/创建 |
| GET | `/api/v1/audit-logs` | 审计日志 |
| GET | `/api/v1/system/health` | 健康检查 |
| GET | `/api/v1/system/metrics` | Prometheus 指标 |
| POST | `/api/v1/import/{type}` | 导入数据 |
| GET | `/api/v1/export/{type}` | 导出数据 |

---

## 9. 数据库设计

### 9.1 表清单

| 表名 | 说明 |
|---|---|
| users | 用户 |
| groups | 用户组 |
| roles | 角色 |
| permissions | 权限码 |
| user_groups | 用户与组关系 |
| role_permissions | 角色权限关系 |
| resource_permissions | Section/Subnet 资源授权 |
| api_tokens | API Token |
| sections | Section |
| folders | Folder |
| subnets | Subnet |
| addresses | IP 地址 |
| vlans | VLAN |
| l2_domains | 二层域 |
| vrfs | VRF |
| devices | 设备 |
| device_types | 设备类型 |
| locations | 位置 |
| racks | 机柜 |
| nats | NAT 映射 |
| nameservers | DNS 服务器 |
| scan_tasks | 扫描任务 |
| scan_results | 扫描结果 |
| ip_requests | IP 申请 |
| tags | 标签 |
| custom_fields | 自定义字段定义 |
| custom_field_values | 自定义字段值 |
| audit_logs | 审计日志 |
| system_settings | 系统设置 |

### 9.2 关键建模要求

1. IP 地址统一使用 `VARBINARY(16)` 或 SQLite BLOB 保存，兼容 IPv4 和 IPv6。
2. 同时保存 `ip_text` 方便展示和搜索。
3. Subnet 保存 network 二进制和 prefix_len，避免纯字符串计算。
4. SQLite 与 MySQL 必须分别提供迁移脚本。
5. SQLC 查询文件必须区分 SQLite / MySQL 方言。
6. 所有数据库访问必须通过 SQLC 生成代码或 Repository 接口。
7. 所有表必须包含 `created_at`、`updated_at`，关键业务表增加 `created_by`、`updated_by`。
8. 删除策略默认软删除可选；首版建议核心资源硬删除 + 审计日志保留。

### 9.3 关键唯一约束

| 表 | 约束 |
|---|---|
| sections | name 唯一 |
| subnets | section_id + vrf_id + network + prefix_len 唯一 |
| addresses | subnet_id + ip 唯一 |
| vlans | l2_domain_id + vlan_id 唯一 |
| vrfs | name 唯一 |
| devices | name 唯一 |
| users | username 唯一 |
| roles | name 唯一 |
| groups | name 唯一 |

---

## 10. 前端需求

### 10.1 技术栈

1. Vue 3。
2. TypeScript。
3. Naive UI。
4. Vue Router。
5. Pinia。
6. Axios。
7. Vite。
8. vue-i18n 预留。
9. ECharts 用于仪表盘和统计图。

### 10.2 页面清单

| 页面 | 功能 |
|---|---|
| 登录页 | 用户登录 |
| 仪表盘 | 系统概览、容量、申请、扫描 |
| Section 页面 | Section 列表、树形导航、详情 |
| Subnet 页面 | 网段列表、树形结构、可视化地址占用 |
| Address 页面 | 地址列表、详情、编辑、探活 |
| VLAN 页面 | VLAN 与 L2 Domain 管理 |
| VRF 页面 | VRF 管理 |
| Device 页面 | 设备管理 |
| Rack 页面 | 机柜可视化与设备上架 |
| Location 页面 | 位置管理 |
| NAT 页面 | NAT 映射管理 |
| DNS 页面 | Nameserver 与解析测试 |
| Scan 页面 | 扫描任务、结果、新主机发现 |
| Request 页面 | IP 申请、审批 |
| Search 页面 | 全局搜索 |
| Import/Export 页面 | 批量导入导出 |
| User/Role 页面 | 用户、用户组、角色、权限 |
| Audit 页面 | 审计日志 |
| Settings 页面 | 系统配置 |

### 10.3 UI 规则

1. 左侧菜单按权限展示。
2. 所有列表支持分页、筛选、排序。
3. 表单统一使用抽屉或弹窗。
4. 危险操作必须二次确认。
5. 删除操作展示影响范围。
6. Subnet 树和 Address 列表联动。
7. IP 地址状态使用标签颜色区分。
8. 使用率超过阈值时高亮预警。

---

## 11. 安全需求

1. 密码使用 Argon2id 哈希。
2. Token 使用安全随机数或 JWT + 签名密钥。
3. API Token 只展示一次，数据库只保存哈希。
4. 所有写操作必须校验 CSRF 或使用 Bearer Token 模式规避 Cookie CSRF。
5. 后端必须做输入校验和权限校验。
6. 所有 SQL 必须参数化，由 SQLC 生成或封装。
7. 支持安全响应头：HSTS、X-Frame-Options、Content-Security-Policy、X-Content-Type-Options。
8. 支持反向代理部署时可信 Header 配置。
9. 审计日志不可被普通管理员删除。
10. 敏感配置不得通过 API 明文返回。

---

## 12. 日志与审计

### 12.1 操作日志

记录以下操作：

1. 登录、退出、失败登录。
2. 用户、角色、权限变更。
3. Section / Folder / Subnet 增删改。
4. Address 增删改、状态变更。
5. VLAN / VRF / Device / Rack / NAT 变更。
6. IP 申请提交、审批、拒绝。
7. 扫描任务创建、完成、失败。
8. 导入导出。
9. 系统设置变更。
10. API Token 创建、禁用、删除。

### 12.2 审计字段

| 字段 | 说明 |
|---|---|
| id | 主键 |
| actor_id | 操作人 |
| action | 操作类型 |
| resource_type | 资源类型 |
| resource_id | 资源 ID |
| before_json | 变更前 |
| after_json | 变更后 |
| ip | 客户端 IP |
| user_agent | User-Agent |
| request_id | 请求 ID |
| created_at | 时间 |

---

## 13. 运维与部署

### 13.1 配置文件

配置文件支持 YAML：

```yaml
server:
  listen: "0.0.0.0:8080"
  public_url: "http://localhost:8080"

database:
  driver: "sqlite"
  dsn: "./data/goipam.db"

security:
  jwt_secret: "change-me"
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

### 13.2 systemd 服务

```ini
[Unit]
Description=GoIPAM Service
After=network.target

[Service]
User=goipam
Group=goipam
ExecStart=/opt/goipam/goipam --config /etc/goipam/config.yaml
Restart=always
RestartSec=3
WorkingDirectory=/opt/goipam

[Install]
WantedBy=multi-user.target
```

### 13.3 健康检查

| 端点 | 说明 |
|---|---|
| `/healthz` | 进程存活 |
| `/readyz` | 数据库可用、迁移完成 |
| `/metrics` | Prometheus 指标 |

### 13.4 备份与恢复

1. SQLite 部署支持备份数据库文件。
2. MySQL 部署使用 mysqldump 或物理备份。
3. 系统提供导出关键数据能力。
4. 升级前必须自动检测数据库版本。
5. Goose 迁移必须可重复执行且可追踪。

---

## 14. 非功能需求

### 14.1 性能

1. 支持至少 10 万 IP 地址记录的常规查询。
2. 地址列表必须分页。
3. 全局搜索响应目标小于 2 秒，具体取决于数据库规模和索引。
4. 扫描任务必须异步执行。
5. 大网段扫描必须限制最大并发和最大地址数。

### 14.2 兼容性

| 项目 | 要求 |
|---|---|
| 操作系统 | Linux amd64/arm64 优先 |
| Go | 1.22+ |
| SQLite | 3.40+ 推荐 |
| MySQL | 8.0+ |
| 浏览器 | Chrome、Edge、Firefox、Safari 近两年版本 |

### 14.3 可维护性

1. 代码按领域模块组织。
2. API DTO 与数据库模型分离。
3. 业务规则放在 Service 层，不写在 Handler 中。
4. 单元测试覆盖 IP 计算、重叠校验、权限校验。
5. 集成测试覆盖 SQLite 和 MySQL。

---

## 15. 开发验收清单

### 15.1 核心 IPAM

1. 可创建 Section。
2. 可创建 Folder。
3. 可创建 IPv4 Subnet。
4. 可创建 IPv6 Subnet。
5. 可创建 IP 地址。
6. 可查询第一个可用 IP。
7. 可识别重叠网段。
8. 可计算使用率。
9. 可展示 Subnet 树。
10. 可导入导出地址。

### 15.2 网络资源

1. 可管理 VLAN。
2. 可管理 L2 Domain。
3. 可管理 VRF。
4. 可管理设备。
5. 可管理 Location。
6. 可管理 Rack。
7. 可完成设备上架。
8. 可管理 NAT。
9. 可配置 Nameserver。

### 15.3 扫描与申请

1. 可扫描单个 Subnet。
2. 可记录在线状态。
3. 可发现新主机。
4. 可提交 IP 申请。
5. 可审批通过并生成地址。
6. 可拒绝申请并记录原因。

### 15.4 权限安全

1. 用户登录正常。
2. 角色权限生效。
3. Section 权限生效。
4. Subnet 权限生效。
5. API Token 权限生效。
6. 无权限资源不可见。
7. 审计日志完整。

### 15.5 运维部署

1. SQLite 单文件部署成功。
2. MySQL 部署成功。
3. systemd 启停正常。
4. Docker Compose 启动正常。
5. 数据库迁移成功。
6. 健康检查正常。
7. Prometheus 指标可访问。

---

## 16. 推荐开发阶段

虽然最终可整体联调验收，但研发实现建议按以下阶段推进：

### 阶段一：基础框架

1. Go 服务框架。
2. Vue 前端框架。
3. 登录认证。
4. 用户、角色、权限。
5. SQLite / MySQL 迁移。
6. 审计日志。

### 阶段二：核心 IPAM

1. Section。
2. Folder。
3. Subnet。
4. Address。
5. IP 计算、重叠检测、可用地址。
6. Subnet 树和地址列表。

### 阶段三：网络资源

1. VLAN / L2 Domain。
2. VRF。
3. Device / Device Type。
4. Rack / Location。
5. NAT。
6. DNS / Nameserver。

### 阶段四：流程和自动化

1. IP 申请审批。
2. 扫描探活。
3. 新主机发现。
4. 导入导出。
5. 全局搜索。

### 阶段五：生产化

1. API 完整化。
2. 安全加固。
3. 指标监控。
4. 备份恢复。
5. Docker / systemd。
6. 文档和测试。

---

## 17. 方案取舍

| 原系统能力 | 目标系统处理 |
|---|---|
| IPv4/IPv6 | 保留并增强 IP 二进制建模 |
| Section/Subnet | 保留，作为核心模型 |
| Visual subnet display | 保留，前端用 Vue 组件重做 |
| 自动扫描/IP 状态 | 保留，Go 异步任务实现 |
| PowerDNS | 首版保留 DNS 查询，PowerDNS 写入后续增强 |
| NAT | 保留，并扩展端口映射字段 |
| Rack | 保留核心 Rack 和上架能力 |
| AD/LDAP/Radius | 首版不做，后续扩展 |
| Per-group permissions | 保留，改为 RBAC + 资源权限 |
| RIPE 导入 | 首版不做 |
| XLS/CSV 导入 | 保留 CSV/XLSX |
| IP request | 保留 |
| REST API | 保留并作为前后端唯一边界 |
| Locations | 保留 |
| Circuits/Customers/PSTN/BGP | 本版本不纳入 |

---

## 18. 结论

本 PRD 建议将核心 IP 地址管理能力建设为现代化 Go + Vue 前后端分离系统。核心策略是：

1. 不逐行迁移旧代码，而是重建领域模型和 API。
2. 保留原系统最有价值的 Section / Subnet / Address / VLAN / VRF / NAT / Rack / DNS / Scan / Request / Permission 能力。
3. 删除或延后复杂辅助模块，降低首版实现复杂度。
4. 使用 Go 单二进制 + SQLite 默认部署，降低落地成本。
5. 支持 MySQL 8.0+、SQLC、Goose，满足生产环境可维护性。
6. 使用 Vue 3 + TypeScript + Naive UI 重做 Web 交互体验。

该文档可作为 Go 语言版 IPAM 系统的开发、评审、排期、测试和验收基线。
