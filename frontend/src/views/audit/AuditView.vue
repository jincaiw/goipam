<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import {
  NPageHeader,
  NCard,
  NButton,
  NDataTable,
  NSelect,
  NInput,
  NDatePicker,
  NSpace,
  useMessage,
} from 'naive-ui'
import type { DataTableColumns } from 'naive-ui'
import { listAuditLogs } from '@/api/ipam'
import type { AuditLog, AuditLogFilters } from '@/api/types'
import { formatDate } from '@/utils/format'
import { useI18n } from 'vue-i18n'

const message = useMessage()
const { t } = useI18n()

// Data
const loading = ref(false)
const logs = ref<AuditLog[]>([])
const pagination = ref({ page: 1, pageSize: 20, itemCount: 0 })

// Filters
const filterUser = ref<string | null>(null)
const filterAction = ref<string | null>(null)
const filterResourceType = ref<string | null>(null)
const filterDateRange = ref<[number, number] | null>(null)

const actionOptions = [
  { label: 'Create', value: 'create' },
  { label: 'Update', value: 'update' },
  { label: 'Delete', value: 'delete' },
  { label: 'Login', value: 'login' },
  { label: 'Logout', value: 'logout' },
]

const resourceTypeOptions = [
  { label: 'Section', value: 'section' },
  { label: 'Subnet', value: 'subnet' },
  { label: 'Address', value: 'address' },
  { label: 'VLAN', value: 'vlan' },
  { label: 'VRF', value: 'vrf' },
  { label: 'Device', value: 'device' },
  { label: 'User', value: 'user' },
]

// Table columns
const columns = computed<DataTableColumns<AuditLog>>(() => [
  { title: t('common.id'), key: 'id', width: 80 },
  {
    title: t('audit.user'),
    key: 'username',
    width: 120,
    render(row) {
      return row.username || '-'
    },
  },
  {
    title: t('audit.action'),
    key: 'action',
    width: 100,
  },
  {
    title: t('audit.resource'),
    key: 'resource_type',
    width: 130,
  },
  {
    title: t('audit.resourceId'),
    key: 'resource_id',
    width: 110,
    render(row) {
      return row.resource_id != null ? String(row.resource_id) : '-'
    },
  },
  {
    title: t('audit.details'),
    key: 'details',
    ellipsis: { tooltip: true },
    render(row) {
      return row.details || '-'
    },
  },
  {
    title: t('audit.ip'),
    key: 'ip_address',
    width: 130,
    render(row) {
      return row.ip_address || '-'
    },
  },
  {
    title: t('audit.time'),
    key: 'created_at',
    width: 180,
    render(row) {
      return formatDate(row.created_at)
    },
  },
])

// Fetch audit logs
async function fetchData() {
  loading.value = true
  try {
    const filters: AuditLogFilters = {}
    if (filterUser.value) filters.user = filterUser.value
    if (filterAction.value) filters.action = filterAction.value
    if (filterResourceType.value) filters.resource_type = filterResourceType.value
    if (filterDateRange.value) {
      filters.start_date = new Date(filterDateRange.value[0]).toISOString()
      filters.end_date = new Date(filterDateRange.value[1]).toISOString()
    }
    const res = await listAuditLogs(pagination.value.page, pagination.value.pageSize, filters)
    if (res.success && res.data) {
      logs.value = res.data.items
      pagination.value.itemCount = res.data.total
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('audit.loadFailed'))
  } finally {
    loading.value = false
  }
}

// Handle page change
function handlePageChange(page: number) {
  pagination.value.page = page
  fetchData()
}

function handlePageSizeChange(pageSize: number) {
  pagination.value.pageSize = pageSize
  pagination.value.page = 1
  fetchData()
}

// Apply filters
function applyFilters() {
  pagination.value.page = 1
  fetchData()
}

// Reset filters
function resetFilters() {
  filterUser.value = null
  filterAction.value = null
  filterResourceType.value = null
  filterDateRange.value = null
  pagination.value.page = 1
  fetchData()
}

onMounted(fetchData)
</script>

<template>
  <div>
    <NPageHeader :title="t('audit.title')" :subtitle="t('audit.subtitle')" />

    <NCard style="margin-top: 16px;">
      <!-- Filters -->
      <NSpace style="margin-bottom: 16px;" align="center">
        <NInput
          v-model:value="filterUser"
          :placeholder="t('audit.filterUser')"
          clearable
          style="width: 160px;"
          @update:value="applyFilters"
        />
        <NSelect
          v-model:value="filterAction"
          :options="actionOptions"
          :placeholder="t('audit.filterAction')"
          clearable
          style="width: 160px;"
          @update:value="applyFilters"
        />
        <NSelect
          v-model:value="filterResourceType"
          :options="resourceTypeOptions"
          :placeholder="t('audit.filterResource')"
          clearable
          style="width: 180px;"
          @update:value="applyFilters"
        />
        <NDatePicker
          v-model:value="filterDateRange"
          type="daterange"
          clearable
          @update:value="applyFilters"
        />
        <NButton @click="resetFilters">{{ t('common.cancel') }}</NButton>
      </NSpace>

      <NDataTable
        :columns="columns"
        :data="logs"
        :loading="loading"
        :pagination="{
          page: pagination.page,
          pageSize: pagination.pageSize,
          itemCount: pagination.itemCount,
          showSizePicker: true,
          pageSizes: [20, 50, 100],
        }"
        :row-key="(row: AuditLog) => row.id"
        @update:page="handlePageChange"
        @update:page-size="handlePageSizeChange"
      />
    </NCard>
  </div>
</template>
