<script setup lang="ts">
import { ref, h, onMounted, computed } from 'vue'
import {
  NPageHeader,
  NCard,
  NButton,
  NDataTable,
  NModal,
  NForm,
  NFormItem,
  NInput,
  NInputNumber,
  NSelect,
  NSpace,
  NPopconfirm,
  useMessage,
} from 'naive-ui'
import type { DataTableColumns, FormInst, FormRules } from 'naive-ui'
import { listNameservers, createNameserver, updateNameserver, deleteNameserver } from '@/api/ipam'
import type { Nameserver, CreateNameserverRequest } from '@/api/types'
import { useI18n } from 'vue-i18n'

const message = useMessage()
const { t } = useI18n()

// Data
const loading = ref(false)
const nameservers = ref<Nameserver[]>([])
const pagination = ref({ page: 1, pageSize: 20, itemCount: 0 })
const showModal = ref(false)
const formRef = ref<FormInst | null>(null)
const saving = ref(false)
const editingId = ref<number | null>(null)

const formValue = ref<CreateNameserverRequest>({
  name: '',
  type: undefined,
  server: '',
  port: 53,
  description: '',
})

const typeOptions = [
  { label: 'DNS', value: 'dns' },
  { label: 'NTP', value: 'ntp' },
]

const formRules: FormRules = {
  name: { required: true, message: () => t('common.required'), trigger: 'blur' },
  server: { required: true, message: () => t('common.required'), trigger: 'blur' },
}

// Table columns
const columns = computed<DataTableColumns<Nameserver>>(() => [
  { title: t('common.id'), key: 'id', width: 80 },
  { title: t('common.name'), key: 'name', ellipsis: { tooltip: true } },
  {
    title: t('dns.type'),
    key: 'type',
    width: 80,
    render(row) {
      return row.type ? row.type.toUpperCase() : '-'
    },
  },
  { title: t('dns.server'), key: 'server', width: 180 },
  {
    title: t('dns.port'),
    key: 'port',
    width: 80,
    render(row) {
      return row.port ?? 53
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
    title: t('common.actions'),
    key: 'actions',
    width: 160,
    render(row) {
      return h(NSpace, { size: 'small' }, {
        default: () => [
          h(NButton, { size: 'small', quaternary: true, onClick: () => handleEdit(row) }, { default: () => t('common.edit') }),
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

// Fetch nameservers
async function fetchData() {
  loading.value = true
  try {
    const res = await listNameservers(pagination.value.page, pagination.value.pageSize)
    if (res.success && res.data) {
      nameservers.value = res.data.items
      pagination.value.itemCount = res.data.total
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('dns.loadFailed'))
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

// Create nameserver
function handleCreate() {
  editingId.value = null
  formValue.value = {
    name: '',
    type: undefined,
    server: '',
    port: 53,
    description: '',
  }
  showModal.value = true
}

// Edit nameserver
function handleEdit(row: Nameserver) {
  editingId.value = row.id
  formValue.value = {
    name: row.name,
    type: row.type || undefined,
    server: row.server,
    port: row.port ?? 53,
    description: row.description || '',
  }
  showModal.value = true
}

// Save nameserver
async function handleSave() {
  try {
    await formRef.value?.validate()
  } catch {
    return
  }
  saving.value = true
  try {
    let res
    if (editingId.value) {
      res = await updateNameserver(editingId.value, formValue.value)
    } else {
      res = await createNameserver(formValue.value)
    }
    if (res.success) {
      message.success(editingId.value ? t('dns.updated') : t('dns.created'))
      showModal.value = false
      fetchData()
    } else {
      message.error(res.message || t('dns.saveFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('dns.saveFailed'))
  } finally {
    saving.value = false
  }
}

// Delete nameserver
async function handleDelete(id: number) {
  try {
    const res = await deleteNameserver(id)
    if (res.success) {
      message.success(t('dns.deleted'))
      fetchData()
    } else {
      message.error(res.message || t('dns.deleteFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('dns.deleteFailed'))
  }
}

onMounted(fetchData)
</script>

<template>
  <div>
    <NPageHeader :title="t('dns.title')" :subtitle="t('dns.subtitle')">
      <template #extra>
        <NButton type="primary" @click="handleCreate">{{ t('dns.newNameserver') }}</NButton>
      </template>
    </NPageHeader>

    <NCard style="margin-top: 16px;">
      <NDataTable
        :columns="columns"
        :data="nameservers"
        :loading="loading"
        :pagination="{
          page: pagination.page,
          pageSize: pagination.pageSize,
          itemCount: pagination.itemCount,
          showSizePicker: true,
          pageSizes: [10, 20, 50],
        }"
        :row-key="(row: Nameserver) => row.id"
        @update:page="handlePageChange"
        @update:page-size="handlePageSizeChange"
      />
    </NCard>

    <!-- Create/Edit Nameserver Modal -->
    <NModal
      v-model:show="showModal"
      preset="dialog"
      :title="editingId ? t('dns.editNameserver') : t('dns.newNameserver')"
      :positive-text="t('common.save')"
      :negative-text="t('common.cancel')"
      :loading="saving"
      @positive-click="handleSave"
    >
      <NForm ref="formRef" :model="formValue" :rules="formRules" label-placement="left" label-width="120">
        <NFormItem :label="t('common.name')" path="name">
          <NInput v-model:value="formValue.name" :placeholder="t('common.name')" />
        </NFormItem>
        <NFormItem :label="t('dns.type')" path="type">
          <NSelect
            v-model:value="formValue.type"
            :options="typeOptions"
            :placeholder="t('dns.type')"
            clearable
          />
        </NFormItem>
        <NFormItem :label="t('dns.server')" path="server">
          <NInput v-model:value="formValue.server" placeholder="e.g. 8.8.8.8" />
        </NFormItem>
        <NFormItem :label="t('dns.port')" path="port">
          <NInputNumber v-model:value="formValue.port" :placeholder="t('dns.port')" :min="1" :max="65535" style="width: 100%;" />
        </NFormItem>
        <NFormItem :label="t('common.description')" path="description">
          <NInput v-model:value="formValue.description" type="textarea" :placeholder="t('common.description')" />
        </NFormItem>
      </NForm>
    </NModal>
  </div>
</template>
