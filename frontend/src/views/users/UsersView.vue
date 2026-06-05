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
  NSwitch,
  NTag,
  NSpace,
  NPopconfirm,
  NTabs,
  NTabPane,
  useMessage,
} from 'naive-ui'
import type { DataTableColumns, FormInst, FormRules } from 'naive-ui'
import { listUsers, createUser, updateUser, deleteUser } from '@/api/ipam'
import type { UserInfo, CreateUserRequest } from '@/api/types'
import { useI18n } from 'vue-i18n'

const message = useMessage()
const { t } = useI18n()

// Data
const loading = ref(false)
const users = ref<UserInfo[]>([])
const pagination = ref({ page: 1, pageSize: 20, itemCount: 0 })
const showModal = ref(false)
const formRef = ref<FormInst | null>(null)
const saving = ref(false)
const editingId = ref<number | null>(null)
const activeTab = ref('users')

const roleOptions = [
  { label: 'Admin', value: 'admin' },
  { label: 'Operator', value: 'operator' },
  { label: 'Viewer', value: 'viewer' },
]

const formValue = ref<Record<string, any>>({
  username: '',
  password: '',
  email: '',
  is_active: true,
  is_superuser: false,
  roles: [],
})

const formRules: FormRules = {
  username: { required: true, message: () => t('common.required'), trigger: 'blur' },
  password: {
    required: true,
    message: () => t('common.required'),
    trigger: 'blur',
    validator(_rule, value) {
      if (!editingId.value && !value) {
        return new Error(t('common.required'))
      }
      return true
    },
  },
}

// Table columns
const columns = computed<DataTableColumns<UserInfo>>(() => [
  { title: t('common.id'), key: 'id', width: 80 },
  { title: t('user.username'), key: 'username', width: 140 },
  {
    title: t('user.email'),
    key: 'email',
    ellipsis: { tooltip: true },
    render(row) {
      return row.email || '-'
    },
  },
  {
    title: t('user.role'),
    key: 'roles',
    width: 200,
    render(row) {
      if (row.is_superuser) {
        return h(NTag, { type: 'error', size: 'small' }, { default: () => 'Superuser' })
      }
      return h(NSpace, { size: 4 }, {
        default: () => (row.roles || []).map((r: string) =>
          h(NTag, { type: 'info', size: 'small' }, { default: () => r }),
        ),
      })
    },
  },
  {
    title: t('common.status'),
    key: 'is_active',
    width: 100,
    render(row) {
      return h(NTag, { type: row.is_active ? 'success' : 'error', size: 'small' }, { default: () => row.is_active ? t('user.active') : t('user.inactive') })
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

// Roles tab columns
const rolesColumns = computed<DataTableColumns<{ name: string; description: string }>>(() => [
  { title: t('user.role'), key: 'name', width: 160 },
  { title: t('common.description'), key: 'description' },
])

const rolesData = [
  { name: 'admin', description: 'Full access to all resources including user management' },
  { name: 'operator', description: 'Can create, edit, and delete IPAM resources' },
  { name: 'viewer', description: 'Read-only access to all resources' },
]

// Fetch users
async function fetchData() {
  loading.value = true
  try {
    const res = await listUsers(pagination.value.page, pagination.value.pageSize)
    if (res.success && res.data) {
      users.value = res.data.items
      pagination.value.itemCount = res.data.total
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('user.loadFailed'))
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

// Create user
function handleCreate() {
  editingId.value = null
  formValue.value = {
    username: '',
    password: '',
    email: '',
    is_active: true,
    is_superuser: false,
    roles: [],
  }
  showModal.value = true
}

// Edit user
function handleEdit(row: UserInfo) {
  editingId.value = row.id
  formValue.value = {
    username: row.username,
    password: '',
    email: row.email,
    is_active: row.is_active,
    is_superuser: row.is_superuser,
    roles: row.roles || [],
  }
  showModal.value = true
}

// Save user
async function handleSave() {
  try {
    await formRef.value?.validate()
  } catch {
    return
  }
  saving.value = true
  try {
    const data: Partial<CreateUserRequest> = {
      username: formValue.value.username,
      email: formValue.value.email || undefined,
      is_active: formValue.value.is_active,
      is_superuser: formValue.value.is_superuser,
      roles: formValue.value.roles,
    }
    // Only send password for new users or when changed
    if (!editingId.value || formValue.value.password) {
      data.password = formValue.value.password
    }
    let res
    if (editingId.value) {
      res = await updateUser(editingId.value, data)
    } else {
      res = await createUser(data as CreateUserRequest)
    }
    if (res.success) {
      message.success(editingId.value ? t('user.updated') : t('user.created'))
      showModal.value = false
      fetchData()
    } else {
      message.error(res.message || t('user.saveFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('user.saveFailed'))
  } finally {
    saving.value = false
  }
}

// Delete user
async function handleDelete(id: number) {
  try {
    const res = await deleteUser(id)
    if (res.success) {
      message.success(t('user.deleted'))
      fetchData()
    } else {
      message.error(res.message || t('user.deleteFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('user.deleteFailed'))
  }
}

onMounted(fetchData)
</script>

<template>
  <div>
    <NPageHeader :title="t('user.title')" :subtitle="t('user.subtitle')">
      <template #extra>
        <NButton type="primary" @click="handleCreate">{{ t('user.newUser') }}</NButton>
      </template>
    </NPageHeader>

    <NCard style="margin-top: 16px;">
      <NTabs v-model:value="activeTab" type="line">
        <NTabPane name="users" :tab="t('user.tabUsers')">
          <NDataTable
            :columns="columns"
            :data="users"
            :loading="loading"
            :pagination="{
              page: pagination.page,
              pageSize: pagination.pageSize,
              itemCount: pagination.itemCount,
              showSizePicker: true,
              pageSizes: [10, 20, 50],
            }"
            :row-key="(row: UserInfo) => row.id"
            @update:page="handlePageChange"
            @update:page-size="handlePageSizeChange"
          />
        </NTabPane>
        <NTabPane name="roles" :tab="t('user.tabRoles')">
          <NDataTable
            :columns="rolesColumns"
            :data="rolesData"
            :bordered="false"
            :row-key="(row) => row.name"
          />
        </NTabPane>
      </NTabs>
    </NCard>

    <!-- Create/Edit User Modal -->
    <NModal
      v-model:show="showModal"
      preset="dialog"
      :title="editingId ? t('user.editUser') : t('user.newUser')"
      :positive-text="t('common.save')"
      :negative-text="t('common.cancel')"
      :loading="saving"
      @positive-click="handleSave"
    >
      <NForm ref="formRef" :model="formValue" :rules="formRules" label-placement="left" label-width="120">
        <NFormItem :label="t('user.username')" path="username">
          <NInput v-model:value="formValue.username" :placeholder="t('user.username')" :disabled="!!editingId" />
        </NFormItem>
        <NFormItem v-if="!editingId" :label="t('user.password')" path="password">
          <NInput v-model:value="formValue.password" type="password" show-password-on="click" :placeholder="t('user.password')" />
        </NFormItem>
        <NFormItem v-else :label="t('user.password')" path="password">
          <NInput v-model:value="formValue.password" type="password" show-password-on="click" :placeholder="t('user.keepPassword')" />
        </NFormItem>
        <NFormItem :label="t('user.email')" path="email">
          <NInput v-model:value="formValue.email" :placeholder="t('user.email')" />
        </NFormItem>
        <NFormItem :label="t('user.role')" path="roles">
          <NSelect
            v-model:value="formValue.roles"
            :options="roleOptions"
            multiple
            :placeholder="t('user.role')"
          />
        </NFormItem>
        <NFormItem :label="t('user.isSuperuser')" path="is_superuser">
          <NSwitch v-model:value="formValue.is_superuser" />
        </NFormItem>
        <NFormItem :label="t('user.isActive')" path="is_active">
          <NSwitch v-model:value="formValue.is_active" />
        </NFormItem>
      </NForm>
    </NModal>
  </div>
</template>
