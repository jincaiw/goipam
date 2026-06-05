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
import { listRacks, createRack, updateRack, deleteRack, listLocations } from '@/api/ipam'
import type { Rack, Location, CreateRackRequest } from '@/api/types'
import { useI18n } from 'vue-i18n'

const message = useMessage()
const { t } = useI18n()

// Data
const loading = ref(false)
const racks = ref<Rack[]>([])
const locations = ref<Location[]>([])
const locationOptions = computed(() =>
  locations.value.map(l => ({ label: l.name, value: l.id })),
)
const pagination = ref({ page: 1, pageSize: 20, itemCount: 0 })
const showModal = ref(false)
const formRef = ref<FormInst | null>(null)
const saving = ref(false)
const editingId = ref<number | null>(null)

const formValue = ref<CreateRackRequest>({
  name: '',
  location_id: 0,
  size_u: 42,
  description: '',
})

const formRules: FormRules = {
  name: { required: true, message: () => t('common.required'), trigger: 'blur' },
  location_id: { required: true, type: 'number', message: () => t('common.required'), trigger: 'change' },
  size_u: { required: true, type: 'number', message: () => t('common.required'), trigger: 'blur' },
}

// Table columns
const columns = computed<DataTableColumns<Rack>>(() => [
  { title: t('common.id'), key: 'id', width: 80 },
  { title: t('common.name'), key: 'name', ellipsis: { tooltip: true } },
  {
    title: t('rack.location'),
    key: 'location_name',
    width: 140,
    render(row) {
      return row.location_name || '-'
    },
  },
  {
    title: t('rack.sizeU'),
    key: 'size',
    width: 80,
    render(row) {
      return `${row.size_u}U`
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

// Fetch locations for dropdown
async function fetchLocations() {
  try {
    const res = await listLocations(1, 100)
    if (res.success && res.data) {
      locations.value = res.data.items
    }
  } catch {
    // ignore
  }
}

// Fetch racks
async function fetchData() {
  loading.value = true
  try {
    const res = await listRacks(pagination.value.page, pagination.value.pageSize)
    if (res.success && res.data) {
      racks.value = res.data.items
      pagination.value.itemCount = res.data.total
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('rack.loadFailed'))
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

// Create rack
function handleCreate() {
  editingId.value = null
  formValue.value = {
    name: '',
    location_id: 0,
    size_u: 42,
    description: '',
  }
  showModal.value = true
}

// Edit rack
function handleEdit(row: Rack) {
  editingId.value = row.id
  formValue.value = {
    name: row.name,
    location_id: row.location_id,
    size_u: row.size_u,
    description: row.description || '',
  }
  showModal.value = true
}

// Save rack
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
      res = await updateRack(editingId.value, formValue.value)
    } else {
      res = await createRack(formValue.value)
    }
    if (res.success) {
      message.success(editingId.value ? t('rack.updated') : t('rack.created'))
      showModal.value = false
      fetchData()
    } else {
      message.error(res.message || t('rack.saveFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('rack.saveFailed'))
  } finally {
    saving.value = false
  }
}

// Delete rack
async function handleDelete(id: number) {
  try {
    const res = await deleteRack(id)
    if (res.success) {
      message.success(t('rack.deleted'))
      fetchData()
    } else {
      message.error(res.message || t('rack.deleteFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('rack.deleteFailed'))
  }
}

onMounted(() => {
  fetchLocations()
  fetchData()
})
</script>

<template>
  <div>
    <NPageHeader :title="t('rack.title')" :subtitle="t('rack.subtitle')">
      <template #extra>
        <NButton type="primary" @click="handleCreate">{{ t('rack.newRack') }}</NButton>
      </template>
    </NPageHeader>

    <NCard style="margin-top: 16px;">
      <NDataTable
        :columns="columns"
        :data="racks"
        :loading="loading"
        :pagination="{
          page: pagination.page,
          pageSize: pagination.pageSize,
          itemCount: pagination.itemCount,
          showSizePicker: true,
          pageSizes: [10, 20, 50],
        }"
        :row-key="(row: Rack) => row.id"
        @update:page="handlePageChange"
        @update:page-size="handlePageSizeChange"
      />
    </NCard>

    <!-- Create/Edit Rack Modal -->
    <NModal
      v-model:show="showModal"
      preset="dialog"
      :title="editingId ? t('rack.editRack') : t('rack.newRack')"
      :positive-text="t('common.save')"
      :negative-text="t('common.cancel')"
      :loading="saving"
      @positive-click="handleSave"
    >
      <NForm ref="formRef" :model="formValue" :rules="formRules" label-placement="left" label-width="120">
        <NFormItem :label="t('common.name')" path="name">
          <NInput v-model:value="formValue.name" :placeholder="t('common.name')" />
        </NFormItem>
        <NFormItem :label="t('rack.location')" path="location_id">
          <NSelect
            v-model:value="formValue.location_id"
            :options="locationOptions"
            :placeholder="t('rack.location')"
          />
        </NFormItem>
        <NFormItem :label="t('rack.sizeU')" path="size_u">
          <NInputNumber v-model:value="formValue.size_u" :placeholder="t('rack.sizeU')" :min="1" style="width: 100%;" />
        </NFormItem>
        <NFormItem :label="t('common.description')" path="description">
          <NInput v-model:value="formValue.description" type="textarea" :placeholder="t('common.description')" />
        </NFormItem>
      </NForm>
    </NModal>
  </div>
</template>
