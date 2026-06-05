<script setup lang="ts">
import { ref, h, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  NPageHeader,
  NCard,
  NDescriptions,
  NDescriptionsItem,
  NTag,
  NButton,
  NDataTable,
  NModal,
  NForm,
  NFormItem,
  NInput,
  NSwitch,
  NSpace,
  NPopconfirm,
  NProgress,
  NStatistic,
  NGrid,
  NGi,
  NSpin,
  NTooltip,
  useMessage,
} from 'naive-ui'
import type { DataTableColumns, FormInst, FormRules } from 'naive-ui'
import {
  getSubnet,
  getSubnetStats,
  getSubnetAddresses,
  createAddress,
  createFirstFreeAddress,
  deleteAddress,
  getSubnetFirstFreeIP,
} from '@/api/ipam'
import type { Subnet, SubnetStats, Address, CreateAddressRequest, FirstFreeAddressRequest } from '@/api/types'
import { formatPercent, formatNumber } from '@/utils/format'
import { useI18n } from 'vue-i18n'

const route = useRoute()
const router = useRouter()
const message = useMessage()
const { t } = useI18n()

const subnetId = computed(() => Number(route.params.id))

// Data
const loading = ref(false)
const subnet = ref<Subnet | null>(null)
const stats = ref<SubnetStats | null>(null)
const addresses = ref<Address[]>([])
const addressPagination = ref({ page: 1, pageSize: 50, itemCount: 0 })
const firstFreeIP = ref<string | null>(null)

// Modals
const showAddAddress = ref(false)
const showAutoAssign = ref(false)
const addFormRef = ref<FormInst | null>(null)
const saving = ref(false)

const addFormValue = ref<CreateAddressRequest>({
  subnet_id: subnetId.value,
  ip: '',
  hostname: '',
  mac: '',
  description: '',
  status: 'active',
  is_gateway: false,
  owner: '',
})

const autoFormValue = ref<FirstFreeAddressRequest>({
  subnet_id: subnetId.value,
  hostname: '',
  mac: '',
  description: '',
  status: 'active',
  is_gateway: false,
  owner: '',
})

const addFormRules: FormRules = {
  ip: { required: true, message: () => t('common.required'), trigger: 'blur' },
}

// Status tag color mapping
const statusColorMap: Record<string, string> = {
  active: 'success',
  inactive: 'default',
  reserved: 'warning',
  pending_confirm: 'info',
  pending: 'info',
  disabled: 'error',
}

// ---- Address Map ----

// Parse CIDR to get network info for the visual map
interface IPBlock {
  ip: string
  status: 'free' | 'used' | 'gateway' | 'pending_confirm'
  hostname: string
}

const addressMap = computed<IPBlock[]>(() => {
  if (!subnet.value || !stats.value) return []

  const cidr = subnet.value.cidr
  const parts = cidr.split('/')
  if (parts.length !== 2) return []

  const prefixLen = parseInt(parts[1], 10)

  // For large subnets (>/22 = 1024+ addresses), limit the map display
  if (prefixLen < 22) return []

  const networkIP = parts[0]
  const ipParts = networkIP.split('.').map(Number)
  const totalAddresses = Math.pow(2, 32 - prefixLen)

  // Build a set of used IPs for quick lookup
  const usedIPs = new Map<string, Address>()
  for (const addr of addresses.value) {
    usedIPs.set(addr.ip_text, addr)
  }

  const blocks: IPBlock[] = []
  for (let i = 0; i < totalAddresses; i++) {
    const ip0 = ipParts[0]
    const ip1 = ipParts[1]
    const ip2 = ipParts[2]
    const ip3 = ipParts[3] + i
    const finalIp3 = (ip3 % 256)
    const carry1 = Math.floor(ip3 / 256)
    const finalIp2 = (ip2 + carry1) % 256
    const carry2 = Math.floor((ip2 + carry1) / 256)
    const finalIp1 = (ip1 + carry2) % 256

    const ip = `${ip0}.${finalIp1}.${finalIp2}.${finalIp3}`
    const addr = usedIPs.get(ip)

    if (addr) {
      if (addr.is_gateway) {
        blocks.push({ ip, status: 'gateway', hostname: addr.hostname })
      } else if (addr.status === 'pending_confirm') {
        blocks.push({ ip, status: 'pending_confirm', hostname: addr.hostname })
      } else {
        blocks.push({ ip, status: 'used', hostname: addr.hostname })
      }
    } else {
      blocks.push({ ip, status: 'free', hostname: '' })
    }
  }

  return blocks
})

const showAddressMap = computed(() => {
  if (!subnet.value) return false
  const parts = subnet.value.cidr.split('/')
  if (parts.length !== 2) return false
  return parseInt(parts[1], 10) >= 22
})

// Address table columns
const addressColumns = computed<DataTableColumns<Address>>(() => [
  {
    title: t('address.ip'),
    key: 'ip_text',
    width: 140,
  },
  {
    title: t('address.hostname'),
    key: 'hostname',
    ellipsis: { tooltip: true },
    render(row) {
      return row.hostname || '-'
    },
  },
  {
    title: t('address.mac'),
    key: 'mac',
    width: 150,
    render(row) {
      return row.mac || '-'
    },
  },
  {
    title: t('common.status'),
    key: 'status',
    width: 110,
    render(row) {
      const type = statusColorMap[row.status] || 'default'
      return h(NTag, { type: type as any, size: 'small' }, { default: () => row.status })
    },
  },
  {
    title: t('common.description'),
    key: 'description',
    ellipsis: { tooltip: true },
    render(row) {
      return row.description || '-'
    },
  },
  {
    title: t('address.owner'),
    key: 'owner',
    width: 120,
    render(row) {
      return row.owner || '-'
    },
  },
  {
    title: t('common.actions'),
    key: 'actions',
    width: 100,
    render(row) {
      return h(
        NPopconfirm,
        { onPositiveClick: () => handleDeleteAddress(row.id) },
        {
          trigger: () => h(NButton, { size: 'small', type: 'error', quaternary: true }, { default: () => t('common.delete') }),
          default: () => t('common.deleteConfirm'),
        },
      )
    },
  },
])

// Fetch data
async function fetchSubnetInfo() {
  loading.value = true
  try {
    const [subnetRes, statsRes] = await Promise.all([
      getSubnet(subnetId.value),
      getSubnetStats(subnetId.value),
    ])
    if (subnetRes.success && subnetRes.data) {
      subnet.value = subnetRes.data
    }
    if (statsRes.success && statsRes.data) {
      stats.value = statsRes.data
    }
    // Fetch first free IP
    try {
      const freeRes = await getSubnetFirstFreeIP(subnetId.value)
      if (freeRes.success && freeRes.data) {
        firstFreeIP.value = freeRes.data.ip
      }
    } catch {
      firstFreeIP.value = null
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('subnet.loadFailed'))
  } finally {
    loading.value = false
  }
}

async function fetchAddresses() {
  try {
    const res = await getSubnetAddresses(subnetId.value, addressPagination.value.page, addressPagination.value.pageSize)
    if (res.success && res.data) {
      addresses.value = res.data.items
      addressPagination.value.itemCount = res.data.total
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('address.loadFailed'))
  }
}

// Address pagination
function handleAddressPageChange(page: number) {
  addressPagination.value.page = page
  fetchAddresses()
}

function handleAddressPageSizeChange(pageSize: number) {
  addressPagination.value.pageSize = pageSize
  addressPagination.value.page = 1
  fetchAddresses()
}

// Add address (manual)
function handleAddAddress() {
  addFormValue.value = {
    subnet_id: subnetId.value,
    ip: firstFreeIP.value || '',
    hostname: '',
    mac: '',
    description: '',
    status: 'active',
    is_gateway: false,
    owner: '',
  }
  showAddAddress.value = true
}

async function handleSaveAddress() {
  try {
    await addFormRef.value?.validate()
  } catch {
    return
  }
  saving.value = true
  try {
    const res = await createAddress(addFormValue.value)
    if (res.success) {
      message.success(t('address.created'))
      showAddAddress.value = false
      fetchAddresses()
      fetchSubnetInfo()
    } else {
      message.error(res.message || t('address.createFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('address.createFailed'))
  } finally {
    saving.value = false
  }
}

// Auto assign (first-free)
function handleAutoAssign() {
  autoFormValue.value = {
    subnet_id: subnetId.value,
    hostname: '',
    mac: '',
    description: '',
    status: 'active',
    is_gateway: false,
    owner: '',
  }
  showAutoAssign.value = true
}

async function handleSaveAutoAssign() {
  saving.value = true
  try {
    const res = await createFirstFreeAddress(autoFormValue.value)
    if (res.success) {
      message.success(t('address.assigned', { ip: res.data?.ip }))
      showAutoAssign.value = false
      fetchAddresses()
      fetchSubnetInfo()
    } else {
      message.error(res.message || t('address.autoAssignFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('address.autoAssignFailed'))
  } finally {
    saving.value = false
  }
}

// Delete address
async function handleDeleteAddress(id: number) {
  try {
    const res = await deleteAddress(id)
    if (res.success) {
      message.success(t('address.deleted'))
      fetchAddresses()
      fetchSubnetInfo()
    } else {
      message.error(res.message || t('address.deleteFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('address.deleteFailed'))
  }
}

onMounted(() => {
  fetchSubnetInfo()
  fetchAddresses()
})
</script>

<template>
  <div>
    <NSpin :show="loading">
      <NPageHeader :title="subnet?.cidr || t('subnet.detailTitle')" :subtitle="subnet?.name || ''" @back="router.push('/subnets')">
        <template #extra>
          <NSpace>
            <NButton @click="handleAddAddress">{{ t('address.addAddress') }}</NButton>
            <NButton type="primary" @click="handleAutoAssign">{{ t('address.autoAssign') }}</NButton>
          </NSpace>
        </template>
      </NPageHeader>

      <!-- Stats -->
      <NGrid :cols="5" :x-gap="16" style="margin-top: 16px;">
        <NGi>
          <NCard>
            <NStatistic :label="t('subnet.cidr')" :value="subnet?.cidr || '-'" />
          </NCard>
        </NGi>
        <NGi>
          <NCard>
            <NStatistic :label="t('subnet.usableAddresses')" :value="formatNumber(stats?.usable_addresses ?? 0)" />
          </NCard>
        </NGi>
        <NGi>
          <NCard>
            <NStatistic :label="t('dashboard.used')" :value="formatNumber(stats?.used_addresses ?? 0)" />
          </NCard>
        </NGi>
        <NGi>
          <NCard>
            <NStatistic :label="t('dashboard.free')" :value="formatNumber(stats?.free_addresses ?? 0)" />
          </NCard>
        </NGi>
        <NGi>
          <NCard>
            <NStatistic :label="t('dashboard.addressUsage')">
              <template #default>
                <span>{{ formatPercent(stats?.usage_percent ?? 0) }}</span>
              </template>
            </NStatistic>
            <NProgress
              type="line"
              :percentage="stats?.usage_percent ?? 0"
              :status="(stats?.usage_percent ?? 0) > 90 ? 'error' : (stats?.usage_percent ?? 0) > 70 ? 'warning' : 'success'"
              style="margin-top: 8px;"
            />
          </NCard>
        </NGi>
      </NGrid>

      <!-- Subnet Info -->
      <NCard :title="t('subnet.subnetInfo')" style="margin-top: 16px;">
        <NDescriptions :column="2" bordered>
          <NDescriptionsItem :label="t('subnet.cidr')">{{ subnet?.cidr }}</NDescriptionsItem>
          <NDescriptionsItem :label="t('common.name')">{{ subnet?.name || '-' }}</NDescriptionsItem>
          <NDescriptionsItem :label="t('common.description')">{{ subnet?.description || '-' }}</NDescriptionsItem>
          <NDescriptionsItem :label="t('common.status')">
            <NTag :type="statusColorMap[subnet?.status || ''] as any || 'default'" size="small">
              {{ subnet?.status || '-' }}
            </NTag>
          </NDescriptionsItem>
          <NDescriptionsItem label="Scan Enabled">
            <NTag :type="subnet?.scan_enabled ? 'success' : 'default'" size="small">
              {{ subnet?.scan_enabled ? t('common.yes') : t('common.no') }}
            </NTag>
          </NDescriptionsItem>
          <NDescriptionsItem label="Discover Enabled">
            <NTag :type="subnet?.discover_enabled ? 'success' : 'default'" size="small">
              {{ subnet?.discover_enabled ? t('common.yes') : t('common.no') }}
            </NTag>
          </NDescriptionsItem>
          <NDescriptionsItem :label="t('subnet.firstFree')">
            <NTag v-if="firstFreeIP" type="success" size="small">{{ firstFreeIP }}</NTag>
            <span v-else style="color: #999;">{{ t('subnet.noAvailableIp') }}</span>
          </NDescriptionsItem>
          <NDescriptionsItem label="Threshold">{{ subnet?.threshold_percent ?? 0 }}%</NDescriptionsItem>
        </NDescriptions>
      </NCard>

      <!-- Address Map -->
      <NCard v-if="showAddressMap" :title="t('subnet.addressMap')" style="margin-top: 16px;">
        <div class="address-map-legend">
          <span class="legend-item"><span class="legend-block legend-free"></span> {{ t('dashboard.free') }}</span>
          <span class="legend-item"><span class="legend-block legend-used"></span> {{ t('dashboard.used') }}</span>
          <span class="legend-item"><span class="legend-block legend-gateway"></span> Gateway</span>
          <span class="legend-item"><span class="legend-block legend-pending"></span> Pending</span>
        </div>
        <div class="address-map">
          <NTooltip v-for="block in addressMap" :key="block.ip" trigger="hover">
            <template #trigger>
              <div
                :class="['address-block', `address-${block.status}`]"
              ></div>
            </template>
            <span>{{ block.ip }}{{ block.hostname ? ' - ' + block.hostname : '' }} ({{ block.status }})</span>
          </NTooltip>
        </div>
      </NCard>

      <!-- Address List -->
      <NCard :title="t('address.title')" style="margin-top: 16px;">
        <NDataTable
          :columns="addressColumns"
          :data="addresses"
          :pagination="{
            page: addressPagination.page,
            pageSize: addressPagination.pageSize,
            itemCount: addressPagination.itemCount,
            showSizePicker: true,
            pageSizes: [20, 50, 100],
          }"
          :row-key="(row: Address) => row.id"
          @update:page="handleAddressPageChange"
          @update:page-size="handleAddressPageSizeChange"
        />
      </NCard>
    </NSpin>

    <!-- Add Address Modal -->
    <NModal
      v-model:show="showAddAddress"
      preset="dialog"
      :title="t('address.addAddress')"
      :positive-text="t('common.create')"
      :negative-text="t('common.cancel')"
      :loading="saving"
      @positive-click="handleSaveAddress"
    >
      <NForm ref="addFormRef" :model="addFormValue" :rules="addFormRules" label-placement="left" label-width="120">
        <NFormItem :label="t('address.ip')" path="ip">
          <NInput v-model:value="addFormValue.ip" placeholder="e.g. 192.168.1.10" />
        </NFormItem>
        <NFormItem :label="t('address.hostname')" path="hostname">
          <NInput v-model:value="addFormValue.hostname" :placeholder="t('address.hostname')" />
        </NFormItem>
        <NFormItem :label="t('address.mac')" path="mac">
          <NInput v-model:value="addFormValue.mac" :placeholder="t('address.mac')" />
        </NFormItem>
        <NFormItem :label="t('common.description')" path="description">
          <NInput v-model:value="addFormValue.description" type="textarea" :placeholder="t('common.description')" />
        </NFormItem>
        <NFormItem :label="t('address.owner')" path="owner">
          <NInput v-model:value="addFormValue.owner" :placeholder="t('address.owner')" />
        </NFormItem>
        <NFormItem label="Is Gateway" path="is_gateway">
          <NSwitch v-model:value="addFormValue.is_gateway" />
        </NFormItem>
      </NForm>
    </NModal>

    <!-- Auto Assign Modal -->
    <NModal
      v-model:show="showAutoAssign"
      preset="dialog"
      :title="t('address.autoAssignTitle')"
      positive-text="Assign"
      :negative-text="t('common.cancel')"
      :loading="saving"
      @positive-click="handleSaveAutoAssign"
    >
      <NForm :model="autoFormValue" label-placement="left" label-width="120">
        <NFormItem :label="t('subnet.firstFree')">
          <NTag v-if="firstFreeIP" type="success">{{ firstFreeIP }}</NTag>
          <span v-else style="color: #999;">{{ t('subnet.noAvailableIp') }}</span>
        </NFormItem>
        <NFormItem :label="t('address.hostname')" path="hostname">
          <NInput v-model:value="autoFormValue.hostname" :placeholder="t('address.hostname')" />
        </NFormItem>
        <NFormItem :label="t('address.mac')" path="mac">
          <NInput v-model:value="autoFormValue.mac" :placeholder="t('address.mac')" />
        </NFormItem>
        <NFormItem :label="t('common.description')" path="description">
          <NInput v-model:value="autoFormValue.description" type="textarea" :placeholder="t('common.description')" />
        </NFormItem>
        <NFormItem :label="t('address.owner')" path="owner">
          <NInput v-model:value="autoFormValue.owner" :placeholder="t('address.owner')" />
        </NFormItem>
      </NForm>
    </NModal>
  </div>
</template>

<style scoped>
.address-map-legend {
  display: flex;
  gap: 16px;
  margin-bottom: 12px;
  font-size: 12px;
  color: #666;
}

.legend-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.legend-block {
  display: inline-block;
  width: 12px;
  height: 12px;
  border-radius: 2px;
}

.legend-free { background-color: #63e2b7; }
.legend-used { background-color: #e88080; }
.legend-gateway { background-color: #70c0e8; }
.legend-pending { background-color: #f2c97d; }

.address-map {
  display: flex;
  flex-wrap: wrap;
  gap: 2px;
}

.address-block {
  width: 8px;
  height: 8px;
  border-radius: 1px;
  cursor: pointer;
  transition: transform 0.1s;
}

.address-block:hover {
  transform: scale(2);
  z-index: 1;
}

.address-free { background-color: #63e2b7; }
.address-used { background-color: #e88080; }
.address-gateway { background-color: #70c0e8; }
.address-pending_confirm { background-color: #f2c97d; }
</style>
