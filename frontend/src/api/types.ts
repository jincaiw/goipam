// Common API response types for GoIPAM

/** Standard API response wrapper (matches backend Response struct) */
export interface ApiResponse<T = unknown> {
  success: boolean
  data: T
  code?: string
  message?: string
  request_id?: string
}

/** Paginated data wrapper (matches backend PagedData struct) */
export interface PagedData<T = unknown> {
  items: T[]
  total: number
  page: number
  page_size: number
}

/** Login request */
export interface LoginRequest {
  username: string
  password: string
}

/** Login response */
export interface LoginResponse {
  access_token: string
  refresh_token: string
  user: UserInfo
}

/** User information */
export interface UserInfo {
  id: number
  username: string
  display_name: string
  email: string
  status: string
  is_active: boolean
  is_superuser: boolean
  roles: string[]
  failed_login_count: number
  locked_until: string | null
  last_login_at: string | null
  created_at: string
  updated_at: string
}

/** Pagination query parameters */
export interface PaginationParams {
  page?: number
  page_size?: number
  search?: string
  ordering?: string
}

// ---- IPAM Domain Types ----

/** Section */
export interface Section {
  id: number
  name: string
  description: string | null
  parent_id: number | null
  strict_mode: boolean
  show_vlan: boolean
  show_vrf: boolean
  subnet_order: string
  sort_order: number
  created_at: string
  updated_at: string
}

/** Section statistics */
export interface SectionStats {
  subnet_count: number
  address_count: number
  total_addresses: number
  usage_percent: number
}

/** Section tree data */
export interface SectionTree {
  section: Section
  folders: Folder[]
  subnets: Subnet[]
}

/** Folder */
export interface Folder {
  id: number
  section_id: number
  parent_id: number | null
  name: string
  description: string | null
  sort_order: number
  created_at: string
  updated_at: string
}

/** Subnet */
export interface Subnet {
  id: number
  section_id: number
  parent_id: number | null
  cidr: string
  ip_version: number
  prefix_len: number
  name: string
  description: string | null
  vlan_id: number | null
  vrf_id: number | null
  location_id: number | null
  nameserver_id: number | null
  scan_enabled: boolean
  discover_enabled: boolean
  threshold_percent: number
  allow_request: boolean
  status: string
  created_at: string
  updated_at: string
}

/** Subnet statistics */
export interface SubnetStats {
  cidr: string
  total_addresses: number
  usable_addresses: number
  used_addresses: number
  free_addresses: number
  usage_percent: number
}

/** Address */
export interface Address {
  id: number
  subnet_id: number
  ip_text: string
  hostname: string
  mac: string
  description: string | null
  status: string
  tag_id: number | null
  is_gateway: boolean
  device_id: number | null
  device_port: string
  owner: string
  last_seen_at: string | null
  dns_checked_at: string | null
  created_at: string
  updated_at: string
}

/** Create section request */
export interface CreateSectionRequest {
  name: string
  description?: string
  parent_id?: number | null
  strict_mode?: boolean
  show_vlan?: boolean
  show_vrf?: boolean
  subnet_order?: string
  sort_order?: number
}

/** Create subnet request */
export interface CreateSubnetRequest {
  section_id: number
  parent_id?: number | null
  cidr: string
  name?: string
  description?: string
  vlan_id?: number | null
  vrf_id?: number | null
  location_id?: number | null
  nameserver_id?: number | null
  scan_enabled?: boolean
  discover_enabled?: boolean
  threshold_percent?: number
  allow_request?: boolean
  status?: string
}

/** Create address request */
export interface CreateAddressRequest {
  subnet_id: number
  ip: string
  hostname?: string
  mac?: string
  description?: string
  status?: string
  tag_id?: number | null
  is_gateway?: boolean
  device_id?: number | null
  device_port?: string
  owner?: string
}

/** First-free address request */
export interface FirstFreeAddressRequest {
  subnet_id: number
  hostname?: string
  mac?: string
  description?: string
  status?: string
  tag_id?: number | null
  is_gateway?: boolean
  device_id?: number | null
  device_port?: string
  owner?: string
}

/** Check overlap request */
export interface CheckOverlapRequest {
  section_id: number
  cidr: string
  vrf_id?: number | null
  exclude_id?: number
}

/** Available range */
export interface AvailableRange {
  start: string
  end: string
  count: number
}

// ---- VLAN ----

export interface Vlan {
  id: number
  name: string
  vlan_id: number
  description: string
  l2_domain_id?: number
  l2_domain_name?: string
  created_at: string
  updated_at: string
}

export interface L2Domain {
  id: number
  name: string
  description: string
  created_at: string
  updated_at: string
}

// ---- VRF ----

export interface Vrf {
  id: number
  name: string
  rd?: string
  description: string
  created_at: string
  updated_at: string
}

// ---- Device ----

export interface Device {
  id: number
  name: string
  hostname?: string
  mgmt_ip?: string
  type_id?: number
  type_name?: string
  location_id?: number
  rack_id?: number
  rack_start_u?: number
  rack_size_u?: number
  manufacturer?: string
  model?: string
  serial_number?: string
  description: string
  created_at: string
  updated_at: string
}

export interface DeviceType {
  id: number
  name: string
  description: string
  created_at: string
  updated_at: string
}

// ---- Location ----

export interface Location {
  id: number
  name: string
  description: string
  parent_id?: number
  address?: string
  created_at: string
  updated_at: string
}

export interface Rack {
  id: number
  name: string
  location_id: number
  location_name?: string
  size_u: number
  description: string
  created_at: string
  updated_at: string
}

// ---- NAT ----

export interface Nat {
  id: number
  name: string
  external_ip: string
  external_port?: number
  internal_ip: string
  internal_port?: number
  protocol?: string
  description: string
  created_at: string
  updated_at: string
}

// ---- Nameserver ----

export interface Nameserver {
  id: number
  name: string
  type?: string
  server: string
  port?: number
  description: string
  created_at: string
  updated_at: string
}

// ---- Scan ----

export interface ScanTask {
  id: number
  subnet_id: number
  subnet_name?: string
  type: string
  status: string
  progress?: number
  result?: string
  created_at: string
  updated_at: string
}

// ---- IP Request ----

export interface IpRequest {
  id: number
  requester: string
  section_id: number
  section_name?: string
  subnet_id?: number
  subnet_name?: string
  address?: string
  description: string
  status: string
  created_at: string
  updated_at: string
}

/** Audit log filter parameters */
export interface AuditLogFilters {
  user?: string
  action?: string
  resource_type?: string
  start_date?: string
  end_date?: string
}

/** Audit log entry */
export interface AuditLog {
  id: number
  user_id?: number
  username?: string
  action: string
  resource_type: string
  resource_id?: number
  details?: string
  ip_address?: string
  created_at: string
}

// ---- Dashboard ----

export interface DashboardData {
  total_sections: number
  total_subnets: number
  total_addresses: number
  address_usage: {
    used: number
    available: number
    free: number
  }
  recent_subnets: {
    id: number
    cidr: string
    name: string
    section_id: number
    created_at: string
  }[] | null
  recent_address_changes: {
    id: number
    ip_text: string
    hostname: string
    status: string
    subnet_id: number
    updated_at: string
  }[] | null
  threshold_subnets: Subnet[] | null
}

// ---- Search ----

export interface SearchResult {
  type: string
  id: number
  name: string
  description?: string
  path?: string
}

// ---- Request types ----

export interface CreateVlanRequest {
  name: string
  vlan_id: number
  description?: string
  l2_domain_id?: number
}

export interface CreateVrfRequest {
  name: string
  rd?: string
  description?: string
}

export interface CreateDeviceRequest {
  name: string
  hostname?: string
  mgmt_ip?: string
  type_id?: number
  location_id?: number
  rack_id?: number
  rack_start_u?: number
  rack_size_u?: number
  manufacturer?: string
  model?: string
  serial_number?: string
  description?: string
}

export interface CreateLocationRequest {
  name: string
  description?: string
  parent_id?: number
  address?: string
}

export interface CreateRackRequest {
  name: string
  location_id: number
  size_u: number
  row_name?: string
  position?: number
  direction?: string
  description?: string
}

export interface CreateNatRequest {
  name: string
  external_ip: string
  external_port?: number
  internal_ip: string
  internal_port?: number
  protocol?: string
  description?: string
}

export interface CreateNameserverRequest {
  name: string
  type?: string
  server: string
  port?: number
  description?: string
}

export interface CreateScanRequest {
  subnet_id: number
  type: string
}

export interface CreateIpRequestRequest {
  requester: string
  section_id: number
  subnet_id?: number
  description: string
}

export interface SearchRequest {
  q: string
  type?: string
}

export interface CreateFolderRequest {
  section_id: number
  parent_id?: number | null
  name: string
  description?: string
  sort_order?: number
}

export interface CreateUserRequest {
  username: string
  password: string
  email?: string
  is_active?: boolean
  is_superuser?: boolean
  roles?: string[]
}
