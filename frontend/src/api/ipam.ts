import { get, post, put, del } from './index'
import type {
  ApiResponse,
  PagedData,
  Section,
  SectionStats,
  SectionTree,
  Folder,
  Subnet,
  SubnetStats,
  Address,
  Vlan,
  L2Domain,
  Vrf,
  Device,
  DeviceType,
  Location,
  Rack,
  Nat,
  Nameserver,
  ScanTask,
  IpRequest,
  AuditLog,
  DashboardData,
  SearchResult,
  UserInfo,
  CreateSectionRequest,
  CreateSubnetRequest,
  CreateAddressRequest,
  FirstFreeAddressRequest,
  CheckOverlapRequest,
  AvailableRange,
  CreateVlanRequest,
  CreateVrfRequest,
  CreateDeviceRequest,
  CreateLocationRequest,
  CreateRackRequest,
  CreateNatRequest,
  CreateNameserverRequest,
  CreateScanRequest,
  CreateIpRequestRequest,
  SearchRequest,
  CreateFolderRequest,
  CreateUserRequest,
} from './types'

// ---- Sections ----

export function listSections(page = 1, pageSize = 20): Promise<ApiResponse<PagedData<Section>>> {
  return get('/sections', { page, page_size: pageSize })
}

export function getSection(id: number): Promise<ApiResponse<Section>> {
  return get(`/sections/${id}`)
}

export function createSection(data: CreateSectionRequest): Promise<ApiResponse<{ id: number }>> {
  return post('/sections', data)
}

export function updateSection(id: number, data: Partial<CreateSectionRequest>): Promise<ApiResponse<{ message: string }>> {
  return put(`/sections/${id}`, data)
}

export function deleteSection(id: number): Promise<ApiResponse<{ message: string }>> {
  return del(`/sections/${id}`)
}

export function getSectionTree(id: number): Promise<ApiResponse<SectionTree>> {
  return get(`/sections/${id}/tree`)
}

export function getSectionStats(id: number): Promise<ApiResponse<SectionStats>> {
  return get(`/sections/${id}/stats`)
}

// ---- Subnets ----

export function listSubnets(sectionId: number, page = 1, pageSize = 20): Promise<ApiResponse<PagedData<Subnet>>> {
  return get('/subnets', { section_id: sectionId, page, page_size: pageSize })
}

export function getSubnet(id: number): Promise<ApiResponse<Subnet>> {
  return get(`/subnets/${id}`)
}

export function createSubnet(data: CreateSubnetRequest): Promise<ApiResponse<{ id: number }>> {
  return post('/subnets', data)
}

export function updateSubnet(id: number, data: Partial<CreateSubnetRequest>): Promise<ApiResponse<{ message: string }>> {
  return put(`/subnets/${id}`, data)
}

export function deleteSubnet(id: number): Promise<ApiResponse<{ message: string }>> {
  return del(`/subnets/${id}`)
}

export function getSubnetChildren(id: number): Promise<ApiResponse<Subnet[]>> {
  return get(`/subnets/${id}/children`)
}

export function getSubnetAddresses(id: number, page = 1, pageSize = 50): Promise<ApiResponse<PagedData<Address>>> {
  return get(`/subnets/${id}/addresses`, { page, page_size: pageSize })
}

export function getSubnetFirstFreeIP(id: number): Promise<ApiResponse<{ ip: string }>> {
  return get(`/subnets/${id}/first-free-ip`)
}

export function getSubnetAvailableRanges(id: number): Promise<ApiResponse<AvailableRange[]>> {
  return get(`/subnets/${id}/available-ranges`)
}

export function getSubnetStats(id: number): Promise<ApiResponse<SubnetStats>> {
  return get(`/subnets/${id}/stats`)
}

export function checkSubnetOverlap(data: CheckOverlapRequest): Promise<ApiResponse<{ overlaps: boolean; overlapping: Subnet[] }>> {
  return post('/subnets/check-overlap', data)
}

// ---- Addresses ----

export function listAddresses(subnetId: number, page = 1, pageSize = 50): Promise<ApiResponse<PagedData<Address>>> {
  return get('/addresses', { subnet_id: subnetId, page, page_size: pageSize })
}

export function getAddress(id: number): Promise<ApiResponse<Address>> {
  return get(`/addresses/${id}`)
}

export function createAddress(data: CreateAddressRequest): Promise<ApiResponse<{ id: number }>> {
  return post('/addresses', data)
}

export function updateAddress(id: number, data: Partial<CreateAddressRequest>): Promise<ApiResponse<{ message: string }>> {
  return put(`/addresses/${id}`, data)
}

export function deleteAddress(id: number): Promise<ApiResponse<{ message: string }>> {
  return del(`/addresses/${id}`)
}

export function createFirstFreeAddress(data: FirstFreeAddressRequest): Promise<ApiResponse<{ id: number; ip: string }>> {
  return post('/addresses/first-free', data)
}

// ---- Folders ----

export function listFolders(sectionId: number, page = 1, pageSize = 20): Promise<ApiResponse<PagedData<Folder>>> {
  return get('/folders', { section_id: sectionId, page, page_size: pageSize })
}

export function getFolder(id: number): Promise<ApiResponse<Folder>> {
  return get(`/folders/${id}`)
}

export function createFolder(data: CreateFolderRequest): Promise<ApiResponse<{ id: number }>> {
  return post('/folders', data)
}

export function updateFolder(id: number, data: Partial<CreateFolderRequest>): Promise<ApiResponse<{ message: string }>> {
  return put(`/folders/${id}`, data)
}

export function deleteFolder(id: number): Promise<ApiResponse<{ message: string }>> {
  return del(`/folders/${id}`)
}

// ---- VLANs ----

export function listVlans(page = 1, pageSize = 20): Promise<ApiResponse<PagedData<Vlan>>> {
  return get('/vlans', { page, page_size: pageSize })
}

export function getVlan(id: number): Promise<ApiResponse<Vlan>> {
  return get(`/vlans/${id}`)
}

export function createVlan(data: CreateVlanRequest): Promise<ApiResponse<{ id: number }>> {
  return post('/vlans', data)
}

export function updateVlan(id: number, data: Partial<CreateVlanRequest>): Promise<ApiResponse<{ message: string }>> {
  return put(`/vlans/${id}`, data)
}

export function deleteVlan(id: number): Promise<ApiResponse<{ message: string }>> {
  return del(`/vlans/${id}`)
}

export function listL2Domains(page = 1, pageSize = 20): Promise<ApiResponse<PagedData<L2Domain>>> {
  return get('/l2-domains', { page, page_size: pageSize })
}

export function createL2Domain(data: { name: string; description?: string }): Promise<ApiResponse<{ id: number }>> {
  return post('/l2-domains', data)
}

// ---- VRFs ----

export function listVrfs(page = 1, pageSize = 20): Promise<ApiResponse<PagedData<Vrf>>> {
  return get('/vrfs', { page, page_size: pageSize })
}

export function getVrf(id: number): Promise<ApiResponse<Vrf>> {
  return get(`/vrfs/${id}`)
}

export function createVrf(data: CreateVrfRequest): Promise<ApiResponse<{ id: number }>> {
  return post('/vrfs', data)
}

export function updateVrf(id: number, data: Partial<CreateVrfRequest>): Promise<ApiResponse<{ message: string }>> {
  return put(`/vrfs/${id}`, data)
}

export function deleteVrf(id: number): Promise<ApiResponse<{ message: string }>> {
  return del(`/vrfs/${id}`)
}

// ---- Devices ----

export function listDevices(page = 1, pageSize = 20): Promise<ApiResponse<PagedData<Device>>> {
  return get('/devices', { page, page_size: pageSize })
}

export function getDevice(id: number): Promise<ApiResponse<Device>> {
  return get(`/devices/${id}`)
}

export function createDevice(data: CreateDeviceRequest): Promise<ApiResponse<{ id: number }>> {
  return post('/devices', data)
}

export function updateDevice(id: number, data: Partial<CreateDeviceRequest>): Promise<ApiResponse<{ message: string }>> {
  return put(`/devices/${id}`, data)
}

export function deleteDevice(id: number): Promise<ApiResponse<{ message: string }>> {
  return del(`/devices/${id}`)
}

export function listDeviceTypes(page = 1, pageSize = 20): Promise<ApiResponse<PagedData<DeviceType>>> {
  return get('/device-types', { page, page_size: pageSize })
}

export function createDeviceType(data: { name: string; description?: string }): Promise<ApiResponse<{ id: number }>> {
  return post('/device-types', data)
}

// ---- Locations ----

export function listLocations(page = 1, pageSize = 20): Promise<ApiResponse<PagedData<Location>>> {
  return get('/locations', { page, page_size: pageSize })
}

export function getLocation(id: number): Promise<ApiResponse<Location>> {
  return get(`/locations/${id}`)
}

export function createLocation(data: CreateLocationRequest): Promise<ApiResponse<{ id: number }>> {
  return post('/locations', data)
}

export function updateLocation(id: number, data: Partial<CreateLocationRequest>): Promise<ApiResponse<{ message: string }>> {
  return put(`/locations/${id}`, data)
}

export function deleteLocation(id: number): Promise<ApiResponse<{ message: string }>> {
  return del(`/locations/${id}`)
}

// ---- Racks ----

export function listRacks(page = 1, pageSize = 20): Promise<ApiResponse<PagedData<Rack>>> {
  return get('/racks', { page, page_size: pageSize })
}

export function getRack(id: number): Promise<ApiResponse<Rack>> {
  return get(`/racks/${id}`)
}

export function createRack(data: CreateRackRequest): Promise<ApiResponse<{ id: number }>> {
  return post('/racks', data)
}

export function updateRack(id: number, data: Partial<CreateRackRequest>): Promise<ApiResponse<{ message: string }>> {
  return put(`/racks/${id}`, data)
}

export function deleteRack(id: number): Promise<ApiResponse<{ message: string }>> {
  return del(`/racks/${id}`)
}

// ---- NATs ----

export function listNats(page = 1, pageSize = 20): Promise<ApiResponse<PagedData<Nat>>> {
  return get('/nats', { page, page_size: pageSize })
}

export function getNat(id: number): Promise<ApiResponse<Nat>> {
  return get(`/nats/${id}`)
}

export function createNat(data: CreateNatRequest): Promise<ApiResponse<{ id: number }>> {
  return post('/nats', data)
}

export function updateNat(id: number, data: Partial<CreateNatRequest>): Promise<ApiResponse<{ message: string }>> {
  return put(`/nats/${id}`, data)
}

export function deleteNat(id: number): Promise<ApiResponse<{ message: string }>> {
  return del(`/nats/${id}`)
}

// ---- Nameservers ----

export function listNameservers(page = 1, pageSize = 20): Promise<ApiResponse<PagedData<Nameserver>>> {
  return get('/nameservers', { page, page_size: pageSize })
}

export function getNameserver(id: number): Promise<ApiResponse<Nameserver>> {
  return get(`/nameservers/${id}`)
}

export function createNameserver(data: CreateNameserverRequest): Promise<ApiResponse<{ id: number }>> {
  return post('/nameservers', data)
}

export function updateNameserver(id: number, data: Partial<CreateNameserverRequest>): Promise<ApiResponse<{ message: string }>> {
  return put(`/nameservers/${id}`, data)
}

export function deleteNameserver(id: number): Promise<ApiResponse<{ message: string }>> {
  return del(`/nameservers/${id}`)
}

// ---- Scans ----

export function createScan(subnetId: number, data: Omit<CreateScanRequest, 'subnet_id'>): Promise<ApiResponse<{ id: number }>> {
  return post(`/scans/subnets/${subnetId}`, data)
}

export function listScanTasks(page = 1, pageSize = 20): Promise<ApiResponse<PagedData<ScanTask>>> {
  return get('/scans/tasks', { page, page_size: pageSize })
}

export function getScanTask(id: number): Promise<ApiResponse<ScanTask>> {
  return get(`/scans/tasks/${id}`)
}

export function cancelScanTask(id: number): Promise<ApiResponse<{ message: string }>> {
  return post(`/scans/tasks/${id}/cancel`)
}

// ---- IP Requests ----

export function listRequests(page = 1, pageSize = 20): Promise<ApiResponse<PagedData<IpRequest>>> {
  return get('/requests', { page, page_size: pageSize })
}

export function createRequest(data: CreateIpRequestRequest): Promise<ApiResponse<{ id: number }>> {
  return post('/requests', data)
}

export function approveRequest(id: number): Promise<ApiResponse<{ message: string }>> {
  return post(`/requests/${id}/approve`)
}

export function rejectRequest(id: number): Promise<ApiResponse<{ message: string }>> {
  return post(`/requests/${id}/reject`)
}

// ---- Search ----

export function search(params: SearchRequest): Promise<ApiResponse<SearchResult[]>> {
  return get('/search', params as unknown as Record<string, unknown>)
}

// ---- Audit ----

export interface AuditLogFilters {
  user?: string
  action?: string
  resource_type?: string
  start_date?: string
  end_date?: string
}

export function listAuditLogs(page = 1, pageSize = 20, filters?: AuditLogFilters): Promise<ApiResponse<PagedData<AuditLog>>> {
  const params: Record<string, unknown> = { page, page_size: pageSize }
  if (filters) {
    if (filters.user) params.user = filters.user
    if (filters.action) params.action = filters.action
    if (filters.resource_type) params.resource_type = filters.resource_type
    if (filters.start_date) params.start_date = filters.start_date
    if (filters.end_date) params.end_date = filters.end_date
  }
  return get('/audit-logs', params)
}

// ---- Dashboard ----

export function getDashboard(): Promise<ApiResponse<DashboardData>> {
  return get('/dashboard')
}

// ---- Users ----

export function listUsers(page = 1, pageSize = 20): Promise<ApiResponse<PagedData<UserInfo>>> {
  return get('/users', { page, page_size: pageSize })
}

export function getUser(id: number): Promise<ApiResponse<UserInfo>> {
  return get(`/users/${id}`)
}

export function createUser(data: CreateUserRequest): Promise<ApiResponse<{ id: number }>> {
  return post('/users', data)
}

export function updateUser(id: number, data: Partial<CreateUserRequest>): Promise<ApiResponse<{ message: string }>> {
  return put(`/users/${id}`, data)
}

export function deleteUser(id: number): Promise<ApiResponse<{ message: string }>> {
  return del(`/users/${id}`)
}

// ---- Settings ----

export function getSettings(category: string): Promise<ApiResponse<Record<string, string>>> {
  return get(`/settings/${category}`)
}

export function saveSettings(category: string, data: Record<string, string>): Promise<ApiResponse<{ message: string }>> {
  return put(`/settings/${category}`, data)
}
