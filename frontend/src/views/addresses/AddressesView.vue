<script setup lang="ts">
import { ref, h, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import {
  NPageHeader,
  NCard,
  NButton,
  NDataTable,
  NInput,
  NTag,
  NSpace,
  NSelect,
  useMessage,
} from 'naive-ui'
import type { DataTableColumns } from 'naive-ui'
import { listAddresses, listSubnets, listSections } from '@/api/ipam'
import type { Address, Subnet, Section } from '@/api/types'
import { useI18n } from 'vue-i18n'

const router = useRouter()
const message = useMessage()
const { t } = useI18n()

// Data
const loading = ref(false)
const addresses = ref<Address[]>([])
const subnets = ref<Subnet[]>([])
const sections = ref<Section[]>([])
const searchQuery = ref('')
const selectedSectionId = ref<number | null>(null)
const selectedSubnetId = ref<number | null>(null)
const pagination = ref({ page: 1, pageSize: 50, itemCount: 0 })

// Computed
const sectionOptions = computed(() =>
  sections.value.map(s => ({ label: s.name, value: s.id })),
)

const subnetOptions = computed(() => {
  let filtered = subnets.value
  if (selectedSectionId.value) {
    filtered = filtered.filter(s => s.section_id === selectedSectionId.value)
  }
  return filtered.map(s => ({ label: `${s.cidr} - ${s.name}`, value: s.id }))
})

// Status tag color mapping
const statusColorMap: Record<string, string> = {
  active: 'success',
  inactive: 'default',
  reserved: 'warning',
  pending_confirm: 'info',
  pending: 'info',
  disabled: 'error',
}

// Filtered addresses (client-side search)
const filteredAddresses = computed(() => {
  if (!searchQuery.value) return addresses.value
  const q = searchQuery.value.toLowerCase()
  return addresses.value.filter(a =>
    a.ip_text?.toLowerCase().includes(q) ||
    a.hostname?.toLowerCase().includes(q) ||
    a.description?.toLowerCase().includes(q)
  )
})

// Table columns
const columns = computed<DataTableColumns<Address>>(() => [
  {
    title: t('address.ip'),
    key: 'ip_text',
    width: 140,
    render(row) {
      return h(
        NButton,
        { text: true, type: 'primary', onClick: () => router.push(`/subnets/${row.subnet_id}`) },
        { default: () => row.ip_text },
      )
    },
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
    title: t('address.subnet'),
    key: 'subnet_id',
    width: 160,
    render(row) {
      const subnet = subnets.value.find(s => s.id === row.subnet_id)
      return subnet ? subnet.cidr : row.subnet_id
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
])

// Fetch sections for filters
async function fetchFilters() {
  try {
    const sectionsRes = await listSections(1, 100)
    if (sectionsRes.success && sectionsRes.data) {
      sections.value = sectionsRes.data.items
    }
  } catch {
    // ignore
  }
}

// Fetch subnets for a given section
async function fetchSubnetsForSection(sectionId: number) {
  try {
    const res = await listSubnets(sectionId, 1, 1000)
    if (res.success && res.data) {
      subnets.value = res.data.items
    }
  } catch {
    // ignore
  }
}

// Fetch addresses
async function fetchData() {
  loading.value = true
  try {
    if (!selectedSubnetId.value) {
      addresses.value = []
      loading.value = false
      return
    }
    const res = await listAddresses(selectedSubnetId.value, pagination.value.page, pagination.value.pageSize)
    if (res.success && res.data) {
      addresses.value = res.data.items
      pagination.value.itemCount = res.data.total
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('address.loadFailed'))
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

// Handle section filter
function handleSectionChange(val: number | null) {
  selectedSectionId.value = val
  selectedSubnetId.value = null
  addresses.value = []
  if (val) {
    fetchSubnetsForSection(val)
  } else {
    subnets.value = []
  }
}

// Handle subnet filter
function handleSubnetChange(val: number | null) {
  selectedSubnetId.value = val
  pagination.value.page = 1
  fetchData()
}

onMounted(() => {
  fetchFilters()
})
</script>

<template>
  <div>
    <NPageHeader :title="t('address.title')" :subtitle="t('address.subtitle')">
      <template #extra>
        <NSpace>
          <NSelect
            :value="selectedSectionId"
            :options="sectionOptions"
            :placeholder="t('subnet.filterSection')"
            clearable
            style="width: 200px;"
            @update:value="handleSectionChange"
          />
          <NSelect
            :value="selectedSubnetId"
            :options="subnetOptions"
            :placeholder="t('address.filterSubnet')"
            clearable
            style="width: 240px;"
            @update:value="handleSubnetChange"
          />
          <NInput
            v-model:value="searchQuery"
            :placeholder="t('common.search')"
            clearable
            style="width: 200px;"
            disabled
          />
        </NSpace>
      </template>
    </NPageHeader>

    <NCard style="margin-top: 16px;">
      <div v-if="!selectedSubnetId" style="text-align: center; padding: 48px; color: #999;">
        {{ t('address.selectSubnet') }}
      </div>
      <NDataTable
        v-else
        :columns="columns"
        :data="filteredAddresses"
        :loading="loading"
        :pagination="{
          page: pagination.page,
          pageSize: pagination.pageSize,
          itemCount: pagination.itemCount,
          showSizePicker: true,
          pageSizes: [20, 50, 100],
        }"
        :row-key="(row: Address) => row.id"
        @update:page="handlePageChange"
        @update:page-size="handlePageSizeChange"
      />
    </NCard>
  </div>
</template>
