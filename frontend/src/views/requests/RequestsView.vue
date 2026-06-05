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
  NTag,
  NSpace,
  useMessage,
} from 'naive-ui'
import type { DataTableColumns, FormInst, FormRules } from 'naive-ui'
import { listRequests, createRequest, approveRequest, rejectRequest, listSections, listSubnets } from '@/api/ipam'
import type { IpRequest, Section, Subnet, CreateIpRequestRequest } from '@/api/types'
import { formatDate } from '@/utils/format'
import { useI18n } from 'vue-i18n'

const message = useMessage()
const { t } = useI18n()

// Data
const loading = ref(false)
const requests = ref<IpRequest[]>([])
const sections = ref<Section[]>([])
const subnets = ref<Subnet[]>([])
const sectionOptions = computed(() =>
  sections.value.map(s => ({ label: s.name, value: s.id })),
)
const subnetOptions = computed(() =>
  subnets.value.map(s => ({ label: `${s.cidr} - ${s.name}`, value: s.id })),
)
const pagination = ref({ page: 1, pageSize: 20, itemCount: 0 })
const showModal = ref(false)
const formRef = ref<FormInst | null>(null)
const saving = ref(false)
const selectedSectionId = ref<number | null>(null)

const formValue = ref<CreateIpRequestRequest>({
  requester: '',
  section_id: 0,
  subnet_id: undefined,
  description: '',
})

const formRules: FormRules = {
  requester: { required: true, message: () => t('common.required'), trigger: 'blur' },
  section_id: { required: true, type: 'number', message: () => t('common.required'), trigger: 'change' },
  description: { required: true, message: () => t('common.required'), trigger: 'blur' },
}

// Status color mapping
const statusColorMap: Record<string, string> = {
  pending: 'warning',
  approved: 'success',
  rejected: 'error',
}

// Table columns
const columns = computed<DataTableColumns<IpRequest>>(() => [
  { title: t('common.id'), key: 'id', width: 80 },
  { title: t('request.requester'), key: 'requester', width: 120 },
  {
    title: t('request.section'),
    key: 'section_name',
    width: 140,
    render(row) {
      return row.section_name || String(row.section_id)
    },
  },
  {
    title: t('request.subnet'),
    key: 'subnet_name',
    width: 160,
    render(row) {
      return row.subnet_name || '-'
    },
  },
  {
    title: t('address.ip'),
    key: 'address',
    width: 140,
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
    title: t('common.status'),
    key: 'status',
    width: 100,
    render(row) {
      return h(NTag, { type: statusColorMap[row.status] as any || 'default', size: 'small' }, { default: () => row.status })
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
    width: 200,
    render(row) {
      const actions: any[] = []
      if (row.status === 'pending') {
        actions.push(
          h(NButton, { size: 'small', type: 'success', quaternary: true, onClick: () => handleApprove(row.id) }, { default: () => t('request.approve') }),
          h(NButton, { size: 'small', type: 'error', quaternary: true, onClick: () => handleReject(row.id) }, { default: () => t('request.reject') }),
        )
      }
      return h(NSpace, { size: 'small' }, { default: () => actions })
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

// Handle section change in form
function handleFormSectionChange(val: number | null) {
  selectedSectionId.value = val
  formValue.value.subnet_id = undefined
  if (val) {
    fetchSubnets(val)
  } else {
    subnets.value = []
  }
}

// Fetch requests
async function fetchData() {
  loading.value = true
  try {
    const res = await listRequests(pagination.value.page, pagination.value.pageSize)
    if (res.success && res.data) {
      requests.value = res.data.items
      pagination.value.itemCount = res.data.total
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('request.loadFailed'))
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

// Create request
function handleCreate() {
  formValue.value = {
    requester: '',
    section_id: 0,
    subnet_id: undefined,
    description: '',
  }
  selectedSectionId.value = null
  subnets.value = []
  showModal.value = true
}

// Save request
async function handleSave() {
  try {
    await formRef.value?.validate()
  } catch {
    return
  }
  saving.value = true
  try {
    const res = await createRequest(formValue.value)
    if (res.success) {
      message.success(t('request.created'))
      showModal.value = false
      fetchData()
    } else {
      message.error(res.message || t('request.createFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('request.createFailed'))
  } finally {
    saving.value = false
  }
}

// Approve request
async function handleApprove(id: number) {
  try {
    const res = await approveRequest(id)
    if (res.success) {
      message.success(t('request.approved'))
      fetchData()
    } else {
      message.error(res.message || t('request.approveFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('request.approveFailed'))
  }
}

// Reject request
async function handleReject(id: number) {
  try {
    const res = await rejectRequest(id)
    if (res.success) {
      message.success(t('request.rejected'))
      fetchData()
    } else {
      message.error(res.message || t('request.rejectFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('request.rejectFailed'))
  }
}

onMounted(() => {
  fetchSections()
  fetchData()
})
</script>

<template>
  <div>
    <NPageHeader :title="t('request.title')" :subtitle="t('request.subtitle')">
      <template #extra>
        <NButton type="primary" @click="handleCreate">{{ t('request.newRequest') }}</NButton>
      </template>
    </NPageHeader>

    <NCard style="margin-top: 16px;">
      <NDataTable
        :columns="columns"
        :data="requests"
        :loading="loading"
        :pagination="{
          page: pagination.page,
          pageSize: pagination.pageSize,
          itemCount: pagination.itemCount,
          showSizePicker: true,
          pageSizes: [10, 20, 50],
        }"
        :row-key="(row: IpRequest) => row.id"
        @update:page="handlePageChange"
        @update:page-size="handlePageSizeChange"
      />
    </NCard>

    <!-- Create Request Modal -->
    <NModal
      v-model:show="showModal"
      preset="dialog"
      :title="t('request.newRequest')"
      positive-text="Submit"
      :negative-text="t('common.cancel')"
      :loading="saving"
      @positive-click="handleSave"
    >
      <NForm ref="formRef" :model="formValue" :rules="formRules" label-placement="left" label-width="120">
        <NFormItem :label="t('request.requester')" path="requester">
          <NInput v-model:value="formValue.requester" :placeholder="t('request.requester')" />
        </NFormItem>
        <NFormItem :label="t('request.section')" path="section_id">
          <NSelect
            v-model:value="formValue.section_id"
            :options="sectionOptions"
            :placeholder="t('request.section')"
            @update:value="handleFormSectionChange"
          />
        </NFormItem>
        <NFormItem :label="t('request.subnet')" path="subnet_id">
          <NSelect
            v-model:value="formValue.subnet_id"
            :options="subnetOptions"
            :placeholder="t('request.subnet')"
            clearable
          />
        </NFormItem>
        <NFormItem :label="t('common.description')" path="description">
          <NInput v-model:value="formValue.description" type="textarea" :placeholder="t('request.describeRequest')" />
        </NFormItem>
      </NForm>
    </NModal>
  </div>
</template>
