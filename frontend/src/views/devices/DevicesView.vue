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
import { listDevices, createDevice, updateDevice, deleteDevice, listDeviceTypes, listLocations, listRacks } from '@/api/ipam'
import type { Device, DeviceType, Location, Rack, CreateDeviceRequest } from '@/api/types'
import { useI18n } from 'vue-i18n'

const message = useMessage()
const { t } = useI18n()

// Data
const loading = ref(false)
const devices = ref<Device[]>([])
const deviceTypes = ref<DeviceType[]>([])
const locations = ref<Location[]>([])
const racks = ref<Rack[]>([])
const deviceTypeOptions = computed(() =>
  deviceTypes.value.map(t => ({ label: t.name, value: t.id })),
)
const locationOptions = computed(() =>
  locations.value.map(l => ({ label: l.name, value: l.id })),
)
const rackOptions = computed(() =>
  racks.value.map(r => ({ label: r.name, value: r.id })),
)
const pagination = ref({ page: 1, pageSize: 20, itemCount: 0 })
const showModal = ref(false)
const formRef = ref<FormInst | null>(null)
const saving = ref(false)
const editingId = ref<number | null>(null)

const formValue = ref<CreateDeviceRequest>({
  name: '',
  hostname: undefined,
  mgmt_ip: undefined,
  type_id: undefined,
  location_id: undefined,
  rack_id: undefined,
  rack_start_u: undefined,
  rack_size_u: undefined,
  manufacturer: undefined,
  model: undefined,
  serial_number: undefined,
  description: '',
})

const formRules: FormRules = {
  name: { required: true, message: () => t('common.required'), trigger: 'blur' },
}

// Table columns
const columns = computed<DataTableColumns<Device>>(() => [
  { title: t('common.id'), key: 'id', width: 80 },
  { title: t('common.name'), key: 'name', ellipsis: { tooltip: true } },
  {
    title: t('device.hostname'),
    key: 'hostname',
    width: 140,
    render(row) {
      return row.hostname || '-'
    },
  },
  {
    title: t('device.mgmtIp'),
    key: 'mgmt_ip',
    width: 140,
    render(row) {
      return row.mgmt_ip || '-'
    },
  },
  {
    title: t('device.type'),
    key: 'type_name',
    width: 120,
    render(row) {
      return row.type_name || '-'
    },
  },
  {
    title: t('device.location'),
    key: 'location_id',
    width: 120,
    render(row) {
      const loc = locations.value.find(l => l.id === row.location_id)
      return loc ? loc.name : '-'
    },
  },
  {
    title: t('device.rack'),
    key: 'rack_id',
    width: 100,
    render(row) {
      const rack = racks.value.find(r => r.id === row.rack_id)
      return rack ? rack.name : '-'
    },
  },
  {
    title: t('device.rackStartU'),
    key: 'rack_start_u',
    width: 100,
    render(row) {
      return row.rack_start_u != null ? String(row.rack_start_u) : '-'
    },
  },
  {
    title: t('device.rackSizeU'),
    key: 'rack_size_u',
    width: 100,
    render(row) {
      return row.rack_size_u != null ? String(row.rack_size_u) : '-'
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

// Fetch related data for dropdowns
async function fetchRelated() {
  try {
    const [typesRes, locRes, rackRes] = await Promise.all([
      listDeviceTypes(1, 100),
      listLocations(1, 100),
      listRacks(1, 100),
    ])
    if (typesRes.success && typesRes.data) {
      deviceTypes.value = typesRes.data.items
    }
    if (locRes.success && locRes.data) {
      locations.value = locRes.data.items
    }
    if (rackRes.success && rackRes.data) {
      racks.value = rackRes.data.items
    }
  } catch {
    // ignore
  }
}

// Fetch devices
async function fetchData() {
  loading.value = true
  try {
    const res = await listDevices(pagination.value.page, pagination.value.pageSize)
    if (res.success && res.data) {
      devices.value = res.data.items
      pagination.value.itemCount = res.data.total
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('device.loadFailed'))
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

// Create device
function handleCreate() {
  editingId.value = null
  formValue.value = {
    name: '',
    hostname: undefined,
    mgmt_ip: undefined,
    type_id: undefined,
    location_id: undefined,
    rack_id: undefined,
    rack_start_u: undefined,
    rack_size_u: undefined,
    manufacturer: undefined,
    model: undefined,
    serial_number: undefined,
    description: '',
  }
  showModal.value = true
}

// Edit device
function handleEdit(row: Device) {
  editingId.value = row.id
  formValue.value = {
    name: row.name,
    hostname: row.hostname || undefined,
    mgmt_ip: row.mgmt_ip || undefined,
    type_id: row.type_id,
    location_id: row.location_id,
    rack_id: row.rack_id,
    rack_start_u: row.rack_start_u,
    rack_size_u: row.rack_size_u,
    manufacturer: row.manufacturer || undefined,
    model: row.model || undefined,
    serial_number: row.serial_number || undefined,
    description: row.description || '',
  }
  showModal.value = true
}

// Save device
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
      res = await updateDevice(editingId.value, formValue.value)
    } else {
      res = await createDevice(formValue.value)
    }
    if (res.success) {
      message.success(editingId.value ? t('device.updated') : t('device.created'))
      showModal.value = false
      fetchData()
    } else {
      message.error(res.message || t('device.saveFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('device.saveFailed'))
  } finally {
    saving.value = false
  }
}

// Delete device
async function handleDelete(id: number) {
  try {
    const res = await deleteDevice(id)
    if (res.success) {
      message.success(t('device.deleted'))
      fetchData()
    } else {
      message.error(res.message || t('device.deleteFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('device.deleteFailed'))
  }
}

onMounted(() => {
  fetchRelated()
  fetchData()
})
</script>

<template>
  <div>
    <NPageHeader :title="t('device.title')" :subtitle="t('device.subtitle')">
      <template #extra>
        <NButton type="primary" @click="handleCreate">{{ t('device.newDevice') }}</NButton>
      </template>
    </NPageHeader>

    <NCard style="margin-top: 16px;">
      <NDataTable
        :columns="columns"
        :data="devices"
        :loading="loading"
        :pagination="{
          page: pagination.page,
          pageSize: pagination.pageSize,
          itemCount: pagination.itemCount,
          showSizePicker: true,
          pageSizes: [10, 20, 50],
        }"
        :row-key="(row: Device) => row.id"
        @update:page="handlePageChange"
        @update:page-size="handlePageSizeChange"
      />
    </NCard>

    <!-- Create/Edit Device Modal -->
    <NModal
      v-model:show="showModal"
      preset="dialog"
      :title="editingId ? t('device.editDevice') : t('device.newDevice')"
      :positive-text="t('common.save')"
      :negative-text="t('common.cancel')"
      :loading="saving"
      @positive-click="handleSave"
    >
      <NForm ref="formRef" :model="formValue" :rules="formRules" label-placement="left" label-width="120">
        <NFormItem :label="t('common.name')" path="name">
          <NInput v-model:value="formValue.name" :placeholder="t('common.name')" />
        </NFormItem>
        <NFormItem :label="t('device.hostname')" path="hostname">
          <NInput v-model:value="formValue.hostname" :placeholder="t('device.hostname')" />
        </NFormItem>
        <NFormItem :label="t('device.mgmtIp')" path="mgmt_ip">
          <NInput v-model:value="formValue.mgmt_ip" :placeholder="t('device.mgmtIp')" />
        </NFormItem>
        <NFormItem :label="t('device.type')" path="type_id">
          <NSelect
            v-model:value="formValue.type_id"
            :options="deviceTypeOptions"
            :placeholder="t('device.type')"
            clearable
          />
        </NFormItem>
        <NFormItem :label="t('device.location')" path="location_id">
          <NSelect
            v-model:value="formValue.location_id"
            :options="locationOptions"
            :placeholder="t('device.location')"
            clearable
          />
        </NFormItem>
        <NFormItem :label="t('device.rack')" path="rack_id">
          <NSelect
            v-model:value="formValue.rack_id"
            :options="rackOptions"
            :placeholder="t('device.rack')"
            clearable
          />
        </NFormItem>
        <NFormItem :label="t('device.rackStartU')" path="rack_start_u">
          <NInputNumber v-model:value="formValue.rack_start_u" :placeholder="t('device.rackStartU')" :min="1" style="width: 100%;" />
        </NFormItem>
        <NFormItem :label="t('device.rackSizeU')" path="rack_size_u">
          <NInputNumber v-model:value="formValue.rack_size_u" :placeholder="t('device.rackSizeU')" :min="1" style="width: 100%;" />
        </NFormItem>
        <NFormItem :label="t('device.manufacturer')" path="manufacturer">
          <NInput v-model:value="formValue.manufacturer" :placeholder="t('device.manufacturer')" />
        </NFormItem>
        <NFormItem :label="t('device.model')" path="model">
          <NInput v-model:value="formValue.model" :placeholder="t('device.model')" />
        </NFormItem>
        <NFormItem :label="t('device.serialNumber')" path="serial_number">
          <NInput v-model:value="formValue.serial_number" :placeholder="t('device.serialNumber')" />
        </NFormItem>
        <NFormItem :label="t('common.description')" path="description">
          <NInput v-model:value="formValue.description" type="textarea" :placeholder="t('common.description')" />
        </NFormItem>
      </NForm>
    </NModal>
  </div>
</template>
