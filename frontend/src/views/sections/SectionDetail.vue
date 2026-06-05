<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  NPageHeader,
  NCard,
  NDescriptions,
  NDescriptionsItem,
  NTag,
  NButton,
  NTree,
  NModal,
  NForm,
  NFormItem,
  NInput,
  NSwitch,
  NProgress,
  NStatistic,
  NGrid,
  NGi,
  NSpin,
  useMessage,
} from 'naive-ui'
import type { TreeOption, FormInst, FormRules } from 'naive-ui'
import { getSection, getSectionStats, getSectionTree, createSubnet } from '@/api/ipam'
import type { Section, SectionStats, SectionTree, CreateSubnetRequest } from '@/api/types'
import { formatPercent, formatNumber } from '@/utils/format'
import { useI18n } from 'vue-i18n'

const route = useRoute()
const router = useRouter()
const message = useMessage()
const { t } = useI18n()

const sectionId = computed(() => Number(route.params.id))

// Data
const loading = ref(false)
const section = ref<Section | null>(null)
const stats = ref<SectionStats | null>(null)
const treeData = ref<TreeOption[]>([])
const showCreateSubnet = ref(false)
const formRef = ref<FormInst | null>(null)
const saving = ref(false)

const formValue = ref<CreateSubnetRequest>({
  section_id: sectionId.value,
  cidr: '',
  name: '',
  description: '',
  scan_enabled: false,
  discover_enabled: false,
})

const formRules: FormRules = {
  cidr: { required: true, message: () => t('common.required'), trigger: 'blur' },
}

// Build tree data from section tree response
function buildTreeOptions(data: SectionTree): TreeOption[] {
  const options: TreeOption[] = []

  // Add folders as tree nodes
  if (data.folders && data.folders.length > 0) {
    const folderMap = new Map<number, TreeOption>()
    for (const folder of data.folders) {
      folderMap.set(folder.id, {
        key: `folder-${folder.id}`,
        label: folder.name,
        prefix: () => '📁',
        children: [],
      })
    }
    // Build folder hierarchy
    for (const folder of data.folders) {
      const node = folderMap.get(folder.id)!
      if (folder.parent_id && folderMap.has(folder.parent_id)) {
        folderMap.get(folder.parent_id)!.children!.push(node)
      } else {
        options.push(node)
      }
    }
  }

  // Add subnets as tree nodes
  if (data.subnets && data.subnets.length > 0) {
    for (const subnet of data.subnets) {
      const node: TreeOption = {
        key: `subnet-${subnet.id}`,
        label: `${subnet.cidr}${subnet.name ? ' - ' + subnet.name : ''}`,
        prefix: () => '🌐',
      }
      // If subnet has a parent_id that matches a folder, add it there
      if (subnet.parent_id) {
        const folderNode = findFolderNode(options, subnet.parent_id)
        if (folderNode && folderNode.children) {
          folderNode.children.push(node)
          continue
        }
      }
      options.push(node)
    }
  }

  return options
}

function findFolderNode(nodes: TreeOption[], folderId: number): TreeOption | null {
  for (const node of nodes) {
    if (node.key === `folder-${folderId}`) return node
    if (node.children) {
      const found = findFolderNode(node.children, folderId)
      if (found) return found
    }
  }
  return null
}

// Fetch data
async function fetchData() {
  loading.value = true
  try {
    const [sectionRes, statsRes, treeRes] = await Promise.all([
      getSection(sectionId.value),
      getSectionStats(sectionId.value),
      getSectionTree(sectionId.value),
    ])
    if (sectionRes.success && sectionRes.data) {
      section.value = sectionRes.data
    }
    if (statsRes.success && statsRes.data) {
      stats.value = statsRes.data
    }
    if (treeRes.success && treeRes.data) {
      treeData.value = buildTreeOptions(treeRes.data)
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('section.loadFailed'))
  } finally {
    loading.value = false
  }
}

// Handle tree node click
function handleTreeSelect(keys: string[]) {
  if (keys.length === 0) return
  const key = keys[0]
  if (typeof key === 'string' && key.startsWith('subnet-')) {
    const subnetId = key.replace('subnet-', '')
    router.push(`/subnets/${subnetId}`)
  }
}

// Create subnet
function handleCreateSubnet() {
  formValue.value = {
    section_id: sectionId.value,
    cidr: '',
    name: '',
    description: '',
    scan_enabled: false,
    discover_enabled: false,
  }
  showCreateSubnet.value = true
}

async function handleSaveSubnet() {
  try {
    await formRef.value?.validate()
  } catch {
    return
  }
  saving.value = true
  try {
    const res = await createSubnet(formValue.value)
    if (res.success) {
      message.success(t('subnet.created'))
      showCreateSubnet.value = false
      fetchData()
    } else {
      message.error(res.message || t('subnet.createFailed'))
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('subnet.createFailed'))
  } finally {
    saving.value = false
  }
}

onMounted(() => {
  fetchData()
})
</script>

<template>
  <div>
    <NSpin :show="loading">
      <NPageHeader :title="section?.name || t('section.detailTitle')" :subtitle="section?.description || ''" @back="router.push('/sections')">
        <template #extra>
          <NButton type="primary" @click="handleCreateSubnet">{{ t('subnet.newSubnet') }}</NButton>
        </template>
      </NPageHeader>

      <!-- Stats -->
      <NGrid :cols="4" :x-gap="16" style="margin-top: 16px;">
        <NGi>
          <NCard>
            <NStatistic :label="t('menu.subnets')" :value="stats?.subnet_count ?? 0" />
          </NCard>
        </NGi>
        <NGi>
          <NCard>
            <NStatistic :label="t('section.usedAddresses')" :value="stats?.address_count ?? 0" />
          </NCard>
        </NGi>
        <NGi>
          <NCard>
            <NStatistic :label="t('section.totalAddresses')" :value="formatNumber(stats?.total_addresses ?? 0)" />
          </NCard>
        </NGi>
        <NGi>
          <NCard>
            <NStatistic :label="t('dashboard.addressUsage')">
              <template #default>
                <span>{{ formatPercent(stats?.usage_percent ?? 0) }}</span>
              </template>
            </NStatistic>
            <NProgress
              type="line"
              :percentage="stats?.usage_percent ?? 0"
              :status="(stats?.usage_percent ?? 0) > 90 ? 'error' : (stats?.usage_percent ?? 0) > 70 ? 'warning' : 'success'"
              style="margin-top: 8px;"
            />
          </NCard>
        </NGi>
      </NGrid>

      <!-- Section Info -->
      <NCard :title="t('section.sectionInfo')" style="margin-top: 16px;">
        <NDescriptions :column="2" bordered>
          <NDescriptionsItem :label="t('common.name')">{{ section?.name }}</NDescriptionsItem>
          <NDescriptionsItem :label="t('common.description')">{{ section?.description || '-' }}</NDescriptionsItem>
          <NDescriptionsItem label="Strict Mode">
            <NTag :type="section?.strict_mode ? 'success' : 'default'" size="small">
              {{ section?.strict_mode ? t('common.yes') : t('common.no') }}
            </NTag>
          </NDescriptionsItem>
          <NDescriptionsItem label="Show VLAN">
            <NTag :type="section?.show_vlan ? 'success' : 'default'" size="small">
              {{ section?.show_vlan ? t('common.yes') : t('common.no') }}
            </NTag>
          </NDescriptionsItem>
          <NDescriptionsItem label="Show VRF">
            <NTag :type="section?.show_vrf ? 'success' : 'default'" size="small">
              {{ section?.show_vrf ? t('common.yes') : t('common.no') }}
            </NTag>
          </NDescriptionsItem>
          <NDescriptionsItem label="Subnet Order">{{ section?.subnet_order || '-' }}</NDescriptionsItem>
        </NDescriptions>
      </NCard>

      <!-- Tree View -->
      <NCard :title="t('section.sectionTree')" style="margin-top: 16px;">
        <NTree
          :data="treeData"
          block-line
          selectable
          :default-expand-all="true"
          @update:selected-keys="handleTreeSelect"
        />
        <div v-if="treeData.length === 0 && !loading" style="text-align: center; padding: 24px; color: #999;">
          {{ t('section.noSubnets') }}
        </div>
      </NCard>
    </NSpin>

    <!-- Create Subnet Modal -->
    <NModal
      v-model:show="showCreateSubnet"
      preset="dialog"
      :title="t('subnet.newSubnet')"
      :positive-text="t('common.create')"
      :negative-text="t('common.cancel')"
      :loading="saving"
      @positive-click="handleSaveSubnet"
    >
      <NForm ref="formRef" :model="formValue" :rules="formRules" label-placement="left" label-width="140">
        <NFormItem :label="t('subnet.cidr')" path="cidr">
          <NInput v-model:value="formValue.cidr" placeholder="e.g. 192.168.1.0/24" />
        </NFormItem>
        <NFormItem :label="t('common.name')" path="name">
          <NInput v-model:value="formValue.name" :placeholder="t('subnet.nameOptional')" />
        </NFormItem>
        <NFormItem :label="t('common.description')" path="description">
          <NInput v-model:value="formValue.description" type="textarea" :placeholder="t('common.description')" />
        </NFormItem>
        <NFormItem label="Scan Enabled" path="scan_enabled">
          <NSwitch v-model:value="formValue.scan_enabled" />
        </NFormItem>
        <NFormItem label="Discover Enabled" path="discover_enabled">
          <NSwitch v-model:value="formValue.discover_enabled" />
        </NFormItem>
      </NForm>
    </NModal>
  </div>
</template>
