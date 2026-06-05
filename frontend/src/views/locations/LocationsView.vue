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
  NSelect,
  NSpace,
  NPopconfirm,
  useMessage,
} from 'naive-ui'
import type { DataTableColumns, FormInst, FormRules } from 'naive-ui'
import { listLocations, createLocation, updateLocation, deleteLocation } from '@/api/ipam'
import type { Location, CreateLocationRequest } from '@/api/types'
import { useI18n } from 'vue-i18n'

const message = useMessage()
const { t } = useI18n()

// Data
const loading = ref(false)
const locations = ref<Location[]>([])
const locationOptions = computed(() =>
  locations.value.map(l => ({ label: l.name, value: l.id })),
)
const pagination = ref({ page: 1, pageSize: 20, itemCount: 0 })
const showModal = ref(false)
const formRef = ref<FormInst | null>(null)
const saving = ref(false)
const editingId = ref<number | null>(null)

const formValue = ref<CreateLocationRequest>({
  name: '',
  parent_id: undefined,
  address: '',
  description: '',
})

const formRules: FormRules = {
  name: { required: true, message: () => t('common.required'), trigger: 'blur' },
}

// Table columns
const columns = computed<DataTableColumns<Location>>(() => [
  { title: t('common.id'), key: 'id', width: 80 },
  { title: t('common.name'), key: 'name', ellipsis: { tooltip: true } },
  {
    title: t('location.parent'),
    key: 'parent_id',
    width: 140,
    render(row) {
      if (!row.parent_id) return '-'
      const parent = locations.value.find(l => l.id === row.parent_id)
      return parent ? parent.name : String(row.parent_id)
    },
  },
  {
    title: t('location.address'),
    key: 'address',
    ellipsis: { tooltip: true },
    render(row) {
      return row.address || '-'
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

// Fetch locations
async function fetchData() {
  loading.value = true
  try {
    const res = await listLocations(pagination.value.page, pagination.value.pageSize)
    if (res.success && res.data) {
      locations.value = res.data.items
      pagination.value.itemCount = res.data.total
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('location.loadFailed'))
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

// Create location
function handleCreate() {
  editingId.value = null
  formValue.value = {
    name: '',
    parent_id: undefined,
    address: '',
    description: '',
  }
  showModal.value = true
}

// Edit location
function handleEdit(row: Location) {
  editingId.value = row.id
  formValue.value = {
    name: row.name,
    parent_id: row.parent_id,
    address: row.address || '',
    description: row.description || '',
  }
  showModal.value = true
}

// Save location
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
      res = await updateLocation(editingId.value, formValue.value)
    } else {
      res = await createLocation(formValue.value)
    }
    if (res.success) {
      message.success(editingId.value ? t('location.updated') : t('location.created'))
      showModal.value = false
      fetchData()
    } else {
      message.error(res.message || t('location.saveFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('location.saveFailed'))
  } finally {
    saving.value = false
  }
}

// Delete location
async function handleDelete(id: number) {
  try {
    const res = await deleteLocation(id)
    if (res.success) {
      message.success(t('location.deleted'))
      fetchData()
    } else {
      message.error(res.message || t('location.deleteFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('location.deleteFailed'))
  }
}

onMounted(fetchData)
</script>

<template>
  <div>
    <NPageHeader :title="t('location.title')" :subtitle="t('location.subtitle')">
      <template #extra>
        <NButton type="primary" @click="handleCreate">{{ t('location.newLocation') }}</NButton>
      </template>
    </NPageHeader>

    <NCard style="margin-top: 16px;">
      <NDataTable
        :columns="columns"
        :data="locations"
        :loading="loading"
        :pagination="{
          page: pagination.page,
          pageSize: pagination.pageSize,
          itemCount: pagination.itemCount,
          showSizePicker: true,
          pageSizes: [10, 20, 50],
        }"
        :row-key="(row: Location) => row.id"
        @update:page="handlePageChange"
        @update:page-size="handlePageSizeChange"
      />
    </NCard>

    <!-- Create/Edit Location Modal -->
    <NModal
      v-model:show="showModal"
      preset="dialog"
      :title="editingId ? t('location.editLocation') : t('location.newLocation')"
      :positive-text="t('common.save')"
      :negative-text="t('common.cancel')"
      :loading="saving"
      @positive-click="handleSave"
    >
      <NForm ref="formRef" :model="formValue" :rules="formRules" label-placement="left" label-width="120">
        <NFormItem :label="t('common.name')" path="name">
          <NInput v-model:value="formValue.name" :placeholder="t('common.name')" />
        </NFormItem>
        <NFormItem :label="t('location.parent')" path="parent_id">
          <NSelect
            v-model:value="formValue.parent_id"
            :options="locationOptions.filter(o => o.value !== editingId)"
            :placeholder="t('location.parent')"
            clearable
          />
        </NFormItem>
        <NFormItem :label="t('location.address')" path="address">
          <NInput v-model:value="formValue.address" type="textarea" :placeholder="t('location.address')" />
        </NFormItem>
        <NFormItem :label="t('common.description')" path="description">
          <NInput v-model:value="formValue.description" type="textarea" :placeholder="t('common.description')" />
        </NFormItem>
      </NForm>
    </NModal>
  </div>
</template>
