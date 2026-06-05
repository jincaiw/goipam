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
import { listVlans, createVlan, updateVlan, deleteVlan, listL2Domains } from '@/api/ipam'
import type { Vlan, L2Domain, CreateVlanRequest } from '@/api/types'
import { useI18n } from 'vue-i18n'

const message = useMessage()
const { t } = useI18n()

// Data
const loading = ref(false)
const vlans = ref<Vlan[]>([])
const l2Domains = ref<L2Domain[]>([])
const l2DomainOptions = computed(() =>
  l2Domains.value.map(d => ({ label: d.name, value: d.id })),
)
const pagination = ref({ page: 1, pageSize: 20, itemCount: 0 })
const showModal = ref(false)
const formRef = ref<FormInst | null>(null)
const saving = ref(false)
const editingId = ref<number | null>(null)

const formValue = ref<CreateVlanRequest>({
  name: '',
  vlan_id: 1,
  description: '',
  l2_domain_id: undefined,
})

const formRules: FormRules = {
  name: { required: true, message: () => t('common.required'), trigger: 'blur' },
  vlan_id: { required: true, type: 'number', message: () => t('common.required'), trigger: 'blur' },
}

// Table columns
const columns = computed<DataTableColumns<Vlan>>(() => [
  { title: t('common.id'), key: 'id', width: 80 },
  { title: t('common.name'), key: 'name', ellipsis: { tooltip: true } },
  { title: t('vlan.vlanId'), key: 'vlan_id', width: 100 },
  {
    title: t('vlan.l2Domain'),
    key: 'l2_domain_name',
    width: 140,
    render(row) {
      return row.l2_domain_name || '-'
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

// Fetch L2 domains for dropdown
async function fetchL2Domains() {
  try {
    const res = await listL2Domains(1, 100)
    if (res.success && res.data) {
      l2Domains.value = res.data.items
    }
  } catch {
    // ignore
  }
}

// Fetch VLANs
async function fetchData() {
  loading.value = true
  try {
    const res = await listVlans(pagination.value.page, pagination.value.pageSize)
    if (res.success && res.data) {
      vlans.value = res.data.items
      pagination.value.itemCount = res.data.total
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('vlan.loadFailed'))
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

// Create VLAN
function handleCreate() {
  editingId.value = null
  formValue.value = {
    name: '',
    vlan_id: 1,
    description: '',
    l2_domain_id: undefined,
  }
  showModal.value = true
}

// Edit VLAN
function handleEdit(row: Vlan) {
  editingId.value = row.id
  formValue.value = {
    name: row.name,
    vlan_id: row.vlan_id,
    description: row.description || '',
    l2_domain_id: row.l2_domain_id,
  }
  showModal.value = true
}

// Save VLAN
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
      res = await updateVlan(editingId.value, formValue.value)
    } else {
      res = await createVlan(formValue.value)
    }
    if (res.success) {
      message.success(editingId.value ? t('vlan.updated') : t('vlan.created'))
      showModal.value = false
      fetchData()
    } else {
      message.error(res.message || t('vlan.saveFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('vlan.saveFailed'))
  } finally {
    saving.value = false
  }
}

// Delete VLAN
async function handleDelete(id: number) {
  try {
    const res = await deleteVlan(id)
    if (res.success) {
      message.success(t('vlan.deleted'))
      fetchData()
    } else {
      message.error(res.message || t('vlan.deleteFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('vlan.deleteFailed'))
  }
}

onMounted(() => {
  fetchL2Domains()
  fetchData()
})
</script>

<template>
  <div>
    <NPageHeader :title="t('vlan.title')" :subtitle="t('vlan.subtitle')">
      <template #extra>
        <NButton type="primary" @click="handleCreate">{{ t('vlan.newVlan') }}</NButton>
      </template>
    </NPageHeader>

    <NCard style="margin-top: 16px;">
      <NDataTable
        :columns="columns"
        :data="vlans"
        :loading="loading"
        :pagination="{
          page: pagination.page,
          pageSize: pagination.pageSize,
          itemCount: pagination.itemCount,
          showSizePicker: true,
          pageSizes: [10, 20, 50],
        }"
        :row-key="(row: Vlan) => row.id"
        @update:page="handlePageChange"
        @update:page-size="handlePageSizeChange"
      />
    </NCard>

    <!-- Create/Edit VLAN Modal -->
    <NModal
      v-model:show="showModal"
      preset="dialog"
      :title="editingId ? t('vlan.editVlan') : t('vlan.newVlan')"
      :positive-text="t('common.save')"
      :negative-text="t('common.cancel')"
      :loading="saving"
      @positive-click="handleSave"
    >
      <NForm ref="formRef" :model="formValue" :rules="formRules" label-placement="left" label-width="120">
        <NFormItem :label="t('common.name')" path="name">
          <NInput v-model:value="formValue.name" :placeholder="t('common.name')" />
        </NFormItem>
        <NFormItem :label="t('vlan.vlanId')" path="vlan_id">
          <NInputNumber v-model:value="formValue.vlan_id" :placeholder="t('vlan.vlanId')" :min="1" :max="4094" style="width: 100%;" />
        </NFormItem>
        <NFormItem :label="t('vlan.l2Domain')" path="l2_domain_id">
          <NSelect
            v-model:value="formValue.l2_domain_id"
            :options="l2DomainOptions"
            :placeholder="t('vlan.l2Domain')"
            clearable
          />
        </NFormItem>
        <NFormItem :label="t('common.description')" path="description">
          <NInput v-model:value="formValue.description" type="textarea" :placeholder="t('common.description')" />
        </NFormItem>
      </NForm>
    </NModal>
  </div>
</template>
