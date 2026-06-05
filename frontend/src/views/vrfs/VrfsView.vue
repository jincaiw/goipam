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
  NSpace,
  NPopconfirm,
  useMessage,
} from 'naive-ui'
import type { DataTableColumns, FormInst, FormRules } from 'naive-ui'
import { listVrfs, createVrf, updateVrf, deleteVrf } from '@/api/ipam'
import type { Vrf, CreateVrfRequest } from '@/api/types'
import { useI18n } from 'vue-i18n'

const message = useMessage()
const { t } = useI18n()

// Data
const loading = ref(false)
const vrfs = ref<Vrf[]>([])
const pagination = ref({ page: 1, pageSize: 20, itemCount: 0 })
const showModal = ref(false)
const formRef = ref<FormInst | null>(null)
const saving = ref(false)
const editingId = ref<number | null>(null)

const formValue = ref<CreateVrfRequest>({
  name: '',
  rd: '',
  description: '',
})

const formRules: FormRules = {
  name: { required: true, message: () => t('common.required'), trigger: 'blur' },
}

// Table columns
const columns = computed<DataTableColumns<Vrf>>(() => [
  { title: t('common.id'), key: 'id', width: 80 },
  { title: t('common.name'), key: 'name', ellipsis: { tooltip: true } },
  {
    title: t('vrf.rd'),
    key: 'rd',
    width: 160,
    render(row) {
      return row.rd || '-'
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

// Fetch VRFs
async function fetchData() {
  loading.value = true
  try {
    const res = await listVrfs(pagination.value.page, pagination.value.pageSize)
    if (res.success && res.data) {
      vrfs.value = res.data.items
      pagination.value.itemCount = res.data.total
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('vrf.loadFailed'))
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

// Create VRF
function handleCreate() {
  editingId.value = null
  formValue.value = {
    name: '',
    rd: '',
    description: '',
  }
  showModal.value = true
}

// Edit VRF
function handleEdit(row: Vrf) {
  editingId.value = row.id
  formValue.value = {
    name: row.name,
    rd: row.rd || '',
    description: row.description || '',
  }
  showModal.value = true
}

// Save VRF
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
      res = await updateVrf(editingId.value, formValue.value)
    } else {
      res = await createVrf(formValue.value)
    }
    if (res.success) {
      message.success(editingId.value ? t('vrf.updated') : t('vrf.created'))
      showModal.value = false
      fetchData()
    } else {
      message.error(res.message || t('vrf.saveFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('vrf.saveFailed'))
  } finally {
    saving.value = false
  }
}

// Delete VRF
async function handleDelete(id: number) {
  try {
    const res = await deleteVrf(id)
    if (res.success) {
      message.success(t('vrf.deleted'))
      fetchData()
    } else {
      message.error(res.message || t('vrf.deleteFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('vrf.deleteFailed'))
  }
}

onMounted(fetchData)
</script>

<template>
  <div>
    <NPageHeader :title="t('vrf.title')" :subtitle="t('vrf.subtitle')">
      <template #extra>
        <NButton type="primary" @click="handleCreate">{{ t('vrf.newVrf') }}</NButton>
      </template>
    </NPageHeader>

    <NCard style="margin-top: 16px;">
      <NDataTable
        :columns="columns"
        :data="vrfs"
        :loading="loading"
        :pagination="{
          page: pagination.page,
          pageSize: pagination.pageSize,
          itemCount: pagination.itemCount,
          showSizePicker: true,
          pageSizes: [10, 20, 50],
        }"
        :row-key="(row: Vrf) => row.id"
        @update:page="handlePageChange"
        @update:page-size="handlePageSizeChange"
      />
    </NCard>

    <!-- Create/Edit VRF Modal -->
    <NModal
      v-model:show="showModal"
      preset="dialog"
      :title="editingId ? t('vrf.editVrf') : t('vrf.newVrf')"
      :positive-text="t('common.save')"
      :negative-text="t('common.cancel')"
      :loading="saving"
      @positive-click="handleSave"
    >
      <NForm ref="formRef" :model="formValue" :rules="formRules" label-placement="left" label-width="120">
        <NFormItem :label="t('common.name')" path="name">
          <NInput v-model:value="formValue.name" :placeholder="t('common.name')" />
        </NFormItem>
        <NFormItem :label="t('vrf.rd')" path="rd">
          <NInput v-model:value="formValue.rd" :placeholder="t('vrf.rd')" />
        </NFormItem>
        <NFormItem :label="t('common.description')" path="description">
          <NInput v-model:value="formValue.description" type="textarea" :placeholder="t('common.description')" />
        </NFormItem>
      </NForm>
    </NModal>
  </div>
</template>
