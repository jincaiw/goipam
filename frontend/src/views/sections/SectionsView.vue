<script setup lang="ts">
import { ref, h, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
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
  NProgress,
  useMessage,
} from 'naive-ui'
import type { DataTableColumns, FormInst, FormRules } from 'naive-ui'
import { listSections, createSection, deleteSection, getSectionStats } from '@/api/ipam'
import type { Section, SectionStats, CreateSectionRequest } from '@/api/types'
import { formatPercent } from '@/utils/format'
import { useI18n } from 'vue-i18n'

const router = useRouter()
const message = useMessage()
const { t } = useI18n()

// Data
const loading = ref(false)
const sections = ref<(Section & { stats?: SectionStats })[]>([])
const pagination = ref({ page: 1, pageSize: 20, itemCount: 0 })
const showModal = ref(false)
const formRef = ref<FormInst | null>(null)
const saving = ref(false)

const formValue = ref<CreateSectionRequest>({
  name: '',
  description: '',
  strict_mode: false,
  show_vlan: false,
  show_vrf: false,
})

const formRules: FormRules = {
  name: { required: true, message: () => t('common.required'), trigger: 'blur' },
}

// Table columns
const columns = computed<DataTableColumns<Section & { stats?: SectionStats }>>(() => [
  {
    title: t('common.name'),
    key: 'name',
    render(row) {
      return h(
        NButton,
        { text: true, type: 'primary', onClick: () => router.push(`/sections/${row.id}`) },
        { default: () => row.name },
      )
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
    title: t('menu.subnets'),
    key: 'subnet_count',
    width: 100,
    render(row) {
      return row.stats?.subnet_count ?? '-'
    },
  },
  {
    title: t('menu.addresses'),
    key: 'address_count',
    width: 100,
    render(row) {
      return row.stats?.address_count ?? '-'
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

// Fetch sections
async function fetchData() {
  loading.value = true
  try {
    const res = await listSections(pagination.value.page, pagination.value.pageSize)
    if (res.success && res.data) {
      sections.value = res.data.items as (Section & { stats?: SectionStats })[]
      pagination.value.itemCount = res.data.total
      // Fetch stats for each section in parallel
      await Promise.all(
        sections.value.map(async (section) => {
          try {
            const statsRes = await getSectionStats(section.id)
            if (statsRes.success && statsRes.data) {
              section.stats = statsRes.data
            }
          } catch {
            // ignore stats errors
          }
        }),
      )
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('section.loadFailed'))
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

// Create section
function handleCreate() {
  formValue.value = {
    name: '',
    description: '',
    strict_mode: false,
    show_vlan: false,
    show_vrf: false,
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
    const res = await createSection(formValue.value)
    if (res.success) {
      message.success(t('section.created'))
      showModal.value = false
      fetchData()
    } else {
      message.error(res.message || t('section.createFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('section.createFailed'))
  } finally {
    saving.value = false
  }
}

// Delete section
async function handleDelete(id: number) {
  try {
    const res = await deleteSection(id)
    if (res.success) {
      message.success(t('section.deleted'))
      fetchData()
    } else {
      message.error(res.message || t('section.deleteFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('section.deleteFailed'))
  }
}

onMounted(() => {
  fetchData()
})
</script>

<template>
  <div>
    <NPageHeader :title="t('section.title')" :subtitle="t('section.subtitle')">
      <template #extra>
        <NButton type="primary" @click="handleCreate">{{ t('section.newSection') }}</NButton>
      </template>
    </NPageHeader>

    <NCard style="margin-top: 16px;">
      <NDataTable
        :columns="columns"
        :data="sections"
        :loading="loading"
        :pagination="{
          page: pagination.page,
          pageSize: pagination.pageSize,
          itemCount: pagination.itemCount,
          showSizePicker: true,
          pageSizes: [10, 20, 50],
        }"
        :row-key="(row: Section) => row.id"
        @update:page="handlePageChange"
        @update:page-size="handlePageSizeChange"
      />
    </NCard>

    <!-- Create Section Modal -->
    <NModal
      v-model:show="showModal"
      preset="dialog"
      :title="t('section.newSection')"
      :positive-text="t('common.create')"
      :negative-text="t('common.cancel')"
      :loading="saving"
      @positive-click="handleSave"
    >
      <NForm ref="formRef" :model="formValue" :rules="formRules" label-placement="left" label-width="120">
        <NFormItem :label="t('common.name')" path="name">
          <NInput v-model:value="formValue.name" :placeholder="t('section.name')" />
        </NFormItem>
        <NFormItem :label="t('common.description')" path="description">
          <NInput v-model:value="formValue.description" type="textarea" :placeholder="t('common.description')" />
        </NFormItem>
        <NFormItem label="Strict Mode" path="strict_mode">
          <NSwitch v-model:value="formValue.strict_mode" />
        </NFormItem>
        <NFormItem label="Show VLAN" path="show_vlan">
          <NSwitch v-model:value="formValue.show_vlan" />
        </NFormItem>
        <NFormItem label="Show VRF" path="show_vrf">
          <NSwitch v-model:value="formValue.show_vrf" />
        </NFormItem>
      </NForm>
    </NModal>
  </div>
</template>
