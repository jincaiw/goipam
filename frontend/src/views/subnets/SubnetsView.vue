<script setup lang="ts">
import { ref, h, onMounted, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {
  NPageHeader,
  NCard,
  NButton,
  NDataTable,
  NModal,
  NForm,
  NFormItem,
  NInput,
  NSwitch,
  NSpace,
  NPopconfirm,
  NTag,
  NProgress,
  NSelect,
  useMessage,
} from 'naive-ui'
import type { DataTableColumns, FormInst, FormRules } from 'naive-ui'
import { listSubnets, createSubnet, deleteSubnet, getSubnetStats } from '@/api/ipam'
import { listSections } from '@/api/ipam'
import type { Subnet, SubnetStats, Section, CreateSubnetRequest } from '@/api/types'
import { formatPercent } from '@/utils/format'
import { useI18n } from 'vue-i18n'

const router = useRouter()
const route = useRoute()
const message = useMessage()
const { t } = useI18n()

// Data
const loading = ref(false)
const subnets = ref<(Subnet & { stats?: SubnetStats })[]>([])
const sections = ref<Section[]>([])
const sectionOptions = computed(() =>
  sections.value.map(s => ({ label: s.name, value: s.id })),
)
const selectedSectionId = ref<number | null>(null)
const pagination = ref({ page: 1, pageSize: 20, itemCount: 0 })
const showModal = ref(false)
const formRef = ref<FormInst | null>(null)
const saving = ref(false)

const formValue = ref<CreateSubnetRequest>({
  section_id: 0,
  cidr: '',
  name: '',
  description: '',
  scan_enabled: false,
  discover_enabled: false,
})

const formRules: FormRules = {
  section_id: { required: true, type: 'number', message: () => t('common.required'), trigger: 'change' },
  cidr: { required: true, message: () => t('common.required'), trigger: 'blur' },
}

// Status tag color mapping
const statusColorMap: Record<string, string> = {
  active: 'success',
  inactive: 'default',
  reserved: 'warning',
  pending: 'info',
}

// Table columns
const columns = computed<DataTableColumns<Subnet & { stats?: SubnetStats }>>(() => [
  {
    title: t('subnet.cidr'),
    key: 'cidr',
    render(row) {
      return h(
        NButton,
        { text: true, type: 'primary', onClick: () => router.push(`/subnets/${row.id}`) },
        { default: () => row.cidr },
      )
    },
  },
  {
    title: t('common.name'),
    key: 'name',
    ellipsis: { tooltip: true },
    render(row) {
      return row.name || '-'
    },
  },
  {
    title: t('subnet.section'),
    key: 'section_id',
    width: 120,
    render(row) {
      const sec = sections.value.find(s => s.id === row.section_id)
      return sec ? sec.name : row.section_id
    },
  },
  {
    title: t('menu.addresses'),
    key: 'used_addresses',
    width: 120,
    render(row) {
      if (row.stats) {
        return `${row.stats.used_addresses} / ${row.stats.usable_addresses}`
      }
      return '-'
    },
  },
  {
    title: t('dashboard.addressUsage'),
    key: 'usage',
    width: 180,
    render(row) {
      const percent = row.stats?.usage_percent ?? 0
      const status = percent > 90 ? 'error' : percent > 70 ? 'warning' : 'success'
      return h('div', { style: 'display: flex; align-items: center; gap: 8px;' }, [
        h(NProgress, { type: 'line', percentage: percent, status, showIndicator: false, style: 'flex: 1' }),
        h('span', { style: 'font-size: 12px; white-space: nowrap;' }, formatPercent(percent)),
      ])
    },
  },
  {
    title: t('common.status'),
    key: 'status',
    width: 100,
    render(row) {
      return h(NTag, { type: statusColorMap[row.status] || 'default' as any, size: 'small' }, { default: () => row.status })
    },
  },
  {
    title: t('common.actions'),
    key: 'actions',
    width: 120,
    render(row) {
      return h(NSpace, { size: 'small' }, {
        default: () => [
          h(
            NPopconfirm,
            { onPositiveClick: () => handleDelete(row.id) },
            {
              trigger: () => h(NButton, { size: 'small', type: 'error', quaternary: true }, { default: () => t('common.delete') }),
              default: () => t('common.deleteConfirm'),
            },
          ),
        ],
      })
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

// Fetch subnets
async function fetchData() {
  loading.value = true
  try {
    // If no section selected, try from query param
    if (!selectedSectionId.value && route.query.section_id) {
      selectedSectionId.value = Number(route.query.section_id)
    }
    if (!selectedSectionId.value) {
      subnets.value = []
      loading.value = false
      return
    }
    const res = await listSubnets(selectedSectionId.value, pagination.value.page, pagination.value.pageSize)
    if (res.success && res.data) {
      subnets.value = res.data.items as (Subnet & { stats?: SubnetStats })[]
      pagination.value.itemCount = res.data.total
      // Fetch stats for each subnet in parallel
      await Promise.all(
        subnets.value.map(async (subnet) => {
          try {
            const statsRes = await getSubnetStats(subnet.id)
            if (statsRes.success && statsRes.data) {
              subnet.stats = statsRes.data
            }
          } catch {
            // ignore stats errors
          }
        }),
      )
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('subnet.loadFailed'))
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

// Handle section filter change
function handleSectionChange(val: number | null) {
  selectedSectionId.value = val
  pagination.value.page = 1
  fetchData()
}

// Create subnet
function handleCreate() {
  formValue.value = {
    section_id: selectedSectionId.value || 0,
    cidr: '',
    name: '',
    description: '',
    scan_enabled: false,
    discover_enabled: false,
  }
  showModal.value = true
}

async function handleSave() {
  try {
    await formRef.value?.validate()
  } catch {
    return
  }
  saving.value = true
  try {
    const res = await createSubnet(formValue.value)
    if (res.success) {
      message.success(t('subnet.created'))
      showModal.value = false
      fetchData()
    } else {
      message.error(res.message || t('subnet.createFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('subnet.createFailed'))
  } finally {
    saving.value = false
  }
}

// Delete subnet
async function handleDelete(id: number) {
  try {
    const res = await deleteSubnet(id)
    if (res.success) {
      message.success(t('subnet.deleted'))
      fetchData()
    } else {
      message.error(res.message || t('subnet.deleteFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('subnet.deleteFailed'))
  }
}

onMounted(() => {
  fetchSections()
  fetchData()
})
</script>

<template>
  <div>
    <NPageHeader :title="t('subnet.title')" :subtitle="t('subnet.subtitle')">
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
          <NButton type="primary" @click="handleCreate">{{ t('subnet.newSubnet') }}</NButton>
        </NSpace>
      </template>
    </NPageHeader>

    <NCard style="margin-top: 16px;">
      <div v-if="!selectedSectionId" style="text-align: center; padding: 48px; color: #999;">
        {{ t('subnet.selectSection') }}
      </div>
      <NDataTable
        v-else
        :columns="columns"
        :data="subnets"
        :loading="loading"
        :pagination="{
          page: pagination.page,
          pageSize: pagination.pageSize,
          itemCount: pagination.itemCount,
          showSizePicker: true,
          pageSizes: [10, 20, 50],
        }"
        :row-key="(row: Subnet) => row.id"
        @update:page="handlePageChange"
        @update:page-size="handlePageSizeChange"
      />
    </NCard>

    <!-- Create Subnet Modal -->
    <NModal
      v-model:show="showModal"
      preset="dialog"
      :title="t('subnet.newSubnet')"
      :positive-text="t('common.create')"
      :negative-text="t('common.cancel')"
      :loading="saving"
      @positive-click="handleSave"
    >
      <NForm ref="formRef" :model="formValue" :rules="formRules" label-placement="left" label-width="140">
        <NFormItem :label="t('subnet.section')" path="section_id">
          <NSelect
            v-model:value="formValue.section_id"
            :options="sectionOptions"
            :placeholder="t('subnet.selectSection')"
          />
        </NFormItem>
        <NFormItem :label="t('subnet.cidr')" path="cidr">
          <NInput v-model:value="formValue.cidr" placeholder="e.g. 192.168.1.0/24" />
        </NFormItem>
        <NFormItem :label="t('common.name')" path="name">
          <NInput v-model:value="formValue.name" :placeholder="t('subnet.nameOptional')" />
        </NFormItem>
        <NFormItem :label="t('common.description')" path="description">
          <NInput v-model:value="formValue.description" type="textarea" :placeholder="t('common.description')" />
        </NFormItem>
        <NFormItem label="Scan Enabled" path="scan_enabled">
          <NSwitch v-model:value="formValue.scan_enabled" />
        </NFormItem>
        <NFormItem label="Discover Enabled" path="discover_enabled">
          <NSwitch v-model:value="formValue.discover_enabled" />
        </NFormItem>
      </NForm>
    </NModal>
  </div>
</template>
