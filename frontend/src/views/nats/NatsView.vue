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
import { listNats, createNat, updateNat, deleteNat } from '@/api/ipam'
import type { Nat, CreateNatRequest } from '@/api/types'
import { useI18n } from 'vue-i18n'

const message = useMessage()
const { t } = useI18n()

// Data
const loading = ref(false)
const nats = ref<Nat[]>([])
const pagination = ref({ page: 1, pageSize: 20, itemCount: 0 })
const showModal = ref(false)
const formRef = ref<FormInst | null>(null)
const saving = ref(false)
const editingId = ref<number | null>(null)

const protocolOptions = [
  { label: 'TCP', value: 'tcp' },
  { label: 'UDP', value: 'udp' },
  { label: 'ICMP', value: 'icmp' },
]

const formValue = ref<CreateNatRequest>({
  name: '',
  external_ip: '',
  external_port: undefined,
  internal_ip: '',
  internal_port: undefined,
  protocol: 'tcp',
  description: '',
})

const formRules: FormRules = {
  name: { required: true, message: () => t('common.required'), trigger: 'blur' },
  external_ip: { required: true, message: () => t('common.required'), trigger: 'blur' },
  internal_ip: { required: true, message: () => t('common.required'), trigger: 'blur' },
}

// Table columns
const columns = computed<DataTableColumns<Nat>>(() => [
  { title: t('common.id'), key: 'id', width: 80 },
  { title: t('common.name'), key: 'name', ellipsis: { tooltip: true } },
  {
    title: t('nat.externalIp'),
    key: 'external',
    width: 160,
    render(row) {
      return row.external_port ? `${row.external_ip}:${row.external_port}` : row.external_ip
    },
  },
  {
    title: t('nat.internalIp'),
    key: 'internal',
    width: 160,
    render(row) {
      return row.internal_port ? `${row.internal_ip}:${row.internal_port}` : row.internal_ip
    },
  },
  {
    title: t('nat.protocol'),
    key: 'protocol',
    width: 90,
    render(row) {
      return (row.protocol || '-').toUpperCase()
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

// Fetch NATs
async function fetchData() {
  loading.value = true
  try {
    const res = await listNats(pagination.value.page, pagination.value.pageSize)
    if (res.success && res.data) {
      nats.value = res.data.items
      pagination.value.itemCount = res.data.total
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('nat.loadFailed'))
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

// Create NAT
function handleCreate() {
  editingId.value = null
  formValue.value = {
    name: '',
    external_ip: '',
    external_port: undefined,
    internal_ip: '',
    internal_port: undefined,
    protocol: 'tcp',
    description: '',
  }
  showModal.value = true
}

// Edit NAT
function handleEdit(row: Nat) {
  editingId.value = row.id
  formValue.value = {
    name: row.name,
    external_ip: row.external_ip,
    external_port: row.external_port,
    internal_ip: row.internal_ip,
    internal_port: row.internal_port,
    protocol: row.protocol || 'tcp',
    description: row.description || '',
  }
  showModal.value = true
}

// Save NAT
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
      res = await updateNat(editingId.value, formValue.value)
    } else {
      res = await createNat(formValue.value)
    }
    if (res.success) {
      message.success(editingId.value ? t('nat.updated') : t('nat.created'))
      showModal.value = false
      fetchData()
    } else {
      message.error(res.message || t('nat.saveFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('nat.saveFailed'))
  } finally {
    saving.value = false
  }
}

// Delete NAT
async function handleDelete(id: number) {
  try {
    const res = await deleteNat(id)
    if (res.success) {
      message.success(t('nat.deleted'))
      fetchData()
    } else {
      message.error(res.message || t('nat.deleteFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('nat.deleteFailed'))
  }
}

onMounted(fetchData)
</script>

<template>
  <div>
    <NPageHeader :title="t('nat.title')" :subtitle="t('nat.subtitle')">
      <template #extra>
        <NButton type="primary" @click="handleCreate">{{ t('nat.newNat') }}</NButton>
      </template>
    </NPageHeader>

    <NCard style="margin-top: 16px;">
      <NDataTable
        :columns="columns"
        :data="nats"
        :loading="loading"
        :pagination="{
          page: pagination.page,
          pageSize: pagination.pageSize,
          itemCount: pagination.itemCount,
          showSizePicker: true,
          pageSizes: [10, 20, 50],
        }"
        :row-key="(row: Nat) => row.id"
        @update:page="handlePageChange"
        @update:page-size="handlePageSizeChange"
      />
    </NCard>

    <!-- Create/Edit NAT Modal -->
    <NModal
      v-model:show="showModal"
      preset="dialog"
      :title="editingId ? t('nat.editNat') : t('nat.newNat')"
      :positive-text="t('common.save')"
      :negative-text="t('common.cancel')"
      :loading="saving"
      @positive-click="handleSave"
    >
      <NForm ref="formRef" :model="formValue" :rules="formRules" label-placement="left" label-width="140">
        <NFormItem :label="t('common.name')" path="name">
          <NInput v-model:value="formValue.name" :placeholder="t('common.name')" />
        </NFormItem>
        <NFormItem :label="t('nat.externalIp')" path="external_ip">
          <NInput v-model:value="formValue.external_ip" :placeholder="t('nat.externalIp')" />
        </NFormItem>
        <NFormItem :label="t('nat.externalPort')" path="external_port">
          <NInputNumber v-model:value="formValue.external_port" :placeholder="t('nat.externalPort')" :min="1" :max="65535" style="width: 100%;" />
        </NFormItem>
        <NFormItem :label="t('nat.internalIp')" path="internal_ip">
          <NInput v-model:value="formValue.internal_ip" :placeholder="t('nat.internalIp')" />
        </NFormItem>
        <NFormItem :label="t('nat.internalPort')" path="internal_port">
          <NInputNumber v-model:value="formValue.internal_port" :placeholder="t('nat.internalPort')" :min="1" :max="65535" style="width: 100%;" />
        </NFormItem>
        <NFormItem :label="t('nat.protocol')" path="protocol">
          <NSelect v-model:value="formValue.protocol" :options="protocolOptions" />
        </NFormItem>
        <NFormItem :label="t('common.description')" path="description">
          <NInput v-model:value="formValue.description" type="textarea" :placeholder="t('common.description')" />
        </NFormItem>
      </NForm>
    </NModal>
  </div>
</template>
