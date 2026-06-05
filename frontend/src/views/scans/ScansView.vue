<script setup lang="ts">
import { ref, h, onMounted, onUnmounted, computed } from 'vue'
import {
  NPageHeader,
  NCard,
  NButton,
  NDataTable,
  NForm,
  NFormItem,
  NSelect,
  NTag,
  NProgress,
  useMessage,
} from 'naive-ui'
import type { DataTableColumns, FormInst, FormRules } from 'naive-ui'
import { listScanTasks, createScan, cancelScanTask, listSubnets, listSections } from '@/api/ipam'
import type { ScanTask, Subnet, Section, CreateScanRequest } from '@/api/types'
import { formatDate } from '@/utils/format'
import { useI18n } from 'vue-i18n'

const message = useMessage()
const { t } = useI18n()

// Data
const loading = ref(false)
const scanTasks = ref<ScanTask[]>([])
const sections = ref<Section[]>([])
const subnets = ref<Subnet[]>([])
const sectionOptions = computed(() =>
  sections.value.map(s => ({ label: s.name, value: s.id })),
)
const subnetOptions = computed(() =>
  subnets.value.map(s => ({ label: `${s.cidr} - ${s.name}`, value: s.id })),
)
const pagination = ref({ page: 1, pageSize: 20, itemCount: 0 })
const creating = ref(false)
const selectedSectionId = ref<number | null>(null)

const scanTypeOptions = [
  { label: 'Ping', value: 'ping' },
  { label: 'ARP', value: 'arp' },
  { label: 'TCP', value: 'tcp' },
]

const formRef = ref<FormInst | null>(null)
const formValue = ref<CreateScanRequest>({
  subnet_id: 0,
  type: 'ping',
})

const formRules: FormRules = {
  subnet_id: { required: true, type: 'number', message: () => t('common.required'), trigger: 'change' },
  type: { required: true, message: () => t('common.required'), trigger: 'change' },
}

// Status color mapping
const statusColorMap: Record<string, string> = {
  pending: 'default',
  running: 'info',
  completed: 'success',
  failed: 'error',
  cancelled: 'warning',
}

// Auto-refresh interval
let refreshInterval: ReturnType<typeof setInterval> | null = null

// Table columns
const columns = computed<DataTableColumns<ScanTask>>(() => [
  { title: t('common.id'), key: 'id', width: 80 },
  {
    title: t('scan.subnet'),
    key: 'subnet_name',
    width: 180,
    render(row) {
      return row.subnet_name || String(row.subnet_id)
    },
  },
  {
    title: t('scan.scanType'),
    key: 'type',
    width: 90,
    render(row) {
      return (row.type || '-').toUpperCase()
    },
  },
  {
    title: t('common.status'),
    key: 'status',
    width: 110,
    render(row) {
      return h(NTag, { type: statusColorMap[row.status] as any || 'default', size: 'small' }, { default: () => row.status })
    },
  },
  {
    title: t('scan.progress'),
    key: 'progress',
    width: 160,
    render(row) {
      if (row.status === 'completed') return h(NProgress, { type: 'line', percentage: 100, status: 'success', showIndicator: true })
      if (row.status === 'failed') return h(NProgress, { type: 'line', percentage: row.progress ?? 0, status: 'error', showIndicator: true })
      if (row.status === 'cancelled') return h(NProgress, { type: 'line', percentage: row.progress ?? 0, status: 'warning', showIndicator: true })
      return h(NProgress, { type: 'line', percentage: row.progress ?? 0, showIndicator: true })
    },
  },
  {
    title: t('common.createdAt'),
    key: 'created_at',
    width: 180,
    render(row) {
      return formatDate(row.created_at)
    },
  },
  {
    title: t('common.actions'),
    key: 'actions',
    width: 100,
    render(row) {
      if (row.status === 'pending' || row.status === 'running') {
        return h(
          NButton,
          { size: 'small', type: 'warning', quaternary: true, onClick: () => handleCancel(row.id) },
          { default: () => t('scan.cancel') },
        )
      }
      return null
    },
  },
])

// Fetch sections for dropdown
async function fetchSections() {
  try {
    const res = await listSections(1, 100)
    if (res.success && res.data) {
      sections.value = res.data.items
    }
  } catch {
    // ignore
  }
}

// Fetch subnets for a section
async function fetchSubnets(sectionId: number) {
  try {
    const res = await listSubnets(sectionId, 1, 1000)
    if (res.success && res.data) {
      subnets.value = res.data.items
    }
  } catch {
    // ignore
  }
}

// Handle section change
function handleSectionChange(val: number | null) {
  selectedSectionId.value = val
  formValue.value.subnet_id = 0
  if (val) {
    fetchSubnets(val)
  } else {
    subnets.value = []
  }
}

// Fetch scan tasks
async function fetchData() {
  loading.value = true
  try {
    const res = await listScanTasks(pagination.value.page, pagination.value.pageSize)
    if (res.success && res.data) {
      scanTasks.value = res.data.items
      pagination.value.itemCount = res.data.total
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('scan.loadFailed'))
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

// Create scan task
async function handleCreateScan() {
  try {
    await formRef.value?.validate()
  } catch {
    return
  }
  creating.value = true
  try {
    const subnetId = formValue.value.subnet_id
    const res = await createScan(subnetId, { type: formValue.value.type })
    if (res.success) {
      message.success(t('scan.created'))
      formValue.value = { subnet_id: 0, type: 'ping' }
      selectedSectionId.value = null
      subnets.value = []
      fetchData()
      // Start auto-refresh if not already running
      startAutoRefresh()
    } else {
      message.error(res.message || t('scan.createFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('scan.createFailed'))
  } finally {
    creating.value = false
  }
}

// Cancel scan task
async function handleCancel(id: number) {
  try {
    const res = await cancelScanTask(id)
    if (res.success) {
      message.success(t('scan.cancelled'))
      fetchData()
    } else {
      message.error(res.message || t('scan.cancelFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('scan.cancelFailed'))
  }
}

// Check if there are running tasks for auto-refresh
function hasRunningTasks(): boolean {
  return scanTasks.value.some(t => t.status === 'pending' || t.status === 'running')
}

// Start auto-refresh
function startAutoRefresh() {
  if (refreshInterval) return
  refreshInterval = setInterval(() => {
    if (hasRunningTasks()) {
      fetchData()
    } else {
      stopAutoRefresh()
    }
  }, 5000)
}

function stopAutoRefresh() {
  if (refreshInterval) {
    clearInterval(refreshInterval)
    refreshInterval = null
  }
}

onMounted(() => {
  fetchSections()
  fetchData().then(() => {
    if (hasRunningTasks()) {
      startAutoRefresh()
    }
  })
})

onUnmounted(() => {
  stopAutoRefresh()
})
</script>

<template>
  <div>
    <NPageHeader :title="t('scan.title')" :subtitle="t('scan.subtitle')" />

    <!-- Create Scan Form -->
    <NCard :title="t('scan.newScan')" style="margin-top: 16px;">
      <NForm ref="formRef" :model="formValue" :rules="formRules" label-placement="left" label-width="100" inline>
        <NFormItem :label="t('scan.section')" path="section_id">
          <NSelect
            :value="selectedSectionId"
            :options="sectionOptions"
            :placeholder="t('scan.section')"
            style="width: 200px;"
            @update:value="handleSectionChange"
          />
        </NFormItem>
        <NFormItem :label="t('scan.subnet')" path="subnet_id">
          <NSelect
            v-model:value="formValue.subnet_id"
            :options="subnetOptions"
            :placeholder="t('scan.subnet')"
            style="width: 240px;"
          />
        </NFormItem>
        <NFormItem :label="t('scan.scanType')" path="type">
          <NSelect
            v-model:value="formValue.type"
            :options="scanTypeOptions"
            style="width: 120px;"
          />
        </NFormItem>
        <NFormItem>
          <NButton type="primary" :loading="creating" @click="handleCreateScan">{{ t('scan.startScan') }}</NButton>
        </NFormItem>
      </NForm>
    </NCard>

    <!-- Scan Tasks Table -->
    <NCard :title="t('scan.scanTasks')" style="margin-top: 16px;">
      <NDataTable
        :columns="columns"
        :data="scanTasks"
        :loading="loading"
        :pagination="{
          page: pagination.page,
          pageSize: pagination.pageSize,
          itemCount: pagination.itemCount,
          showSizePicker: true,
          pageSizes: [10, 20, 50],
        }"
        :row-key="(row: ScanTask) => row.id"
        @update:page="handlePageChange"
        @update:page-size="handlePageSizeChange"
      />
    </NCard>
  </div>
</template>
