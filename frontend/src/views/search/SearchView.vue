<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import {
  NPageHeader,
  NCard,
  NInput,
  NTag,
  NSpin,
  NEmpty,
  useMessage,
} from 'naive-ui'
import { search } from '@/api/ipam'
import type { SearchResult, SearchRequest } from '@/api/types'
import { useI18n } from 'vue-i18n'

const message = useMessage()
const { t } = useI18n()

const searchQuery = ref('')
const loading = ref(false)
const results = ref<SearchResult[]>([])
let debounceTimer: ReturnType<typeof setTimeout> | null = null

// Type color mapping
const typeColorMap: Record<string, string> = {
  Section: 'info',
  Subnet: 'success',
  Address: 'warning',
  VLAN: 'default',
  VRF: 'default',
  Device: 'default',
}

// Group results by type
function groupByType(items: SearchResult[]): { type: string; items: SearchResult[] }[] {
  const groups: Record<string, SearchResult[]> = {}
  for (const item of items) {
    if (!groups[item.type]) {
      groups[item.type] = []
    }
    groups[item.type].push(item)
  }
  const typeOrder = ['Section', 'Subnet', 'Address', 'VLAN', 'VRF', 'Device']
  return typeOrder
    .filter(t => groups[t])
    .map(t => ({ type: t, items: groups[t] }))
    .concat(
      Object.keys(groups)
        .filter(t => !typeOrder.includes(t))
        .map(t => ({ type: t, items: groups[t] })),
    )
}

async function doSearch() {
  const q = searchQuery.value.trim()
  if (!q) {
    results.value = []
    return
  }
  loading.value = true
  try {
    const res = await search({ q } as SearchRequest)
    if (res.success && res.data) {
      results.value = res.data
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('search.failed'))
  } finally {
    loading.value = false
  }
}

// Debounced search (300ms)
function handleInput() {
  if (debounceTimer) {
    clearTimeout(debounceTimer)
  }
  debounceTimer = setTimeout(() => {
    doSearch()
  }, 300)
}

onMounted(() => {
  // Focus the search input on mount
})

onUnmounted(() => {
  if (debounceTimer) {
    clearTimeout(debounceTimer)
  }
})
</script>

<template>
  <div>
    <NPageHeader :title="t('search.title')" :subtitle="t('search.subtitle')" />

    <NCard style="margin-top: 16px;">
      <NInput
        v-model:value="searchQuery"
        :placeholder="t('search.placeholder')"
        size="large"
        clearable
        @input="handleInput"
        @clear="results = []"
      />
    </NCard>

    <NSpin :show="loading" style="margin-top: 16px;">
      <div v-if="!searchQuery.trim()" style="text-align: center; padding: 48px; color: #999;">
        {{ t('search.typeToSearch') }}
      </div>
      <div v-else-if="!results.length && !loading" style="text-align: center; padding: 48px;">
        <NEmpty :description="t('search.noResults')" />
      </div>
      <div v-else>
        <div v-for="group in groupByType(results)" :key="group.type" style="margin-top: 16px;">
          <NCard :title="group.type" size="small">
            <div v-for="item in group.items" :key="`${item.type}-${item.id}`" style="padding: 8px 0; border-bottom: 1px solid #f0f0f0;">
              <div style="display: flex; align-items: center; gap: 8px;">
                <NTag :type="typeColorMap[item.type] as any || 'default'" size="small">{{ item.type }}</NTag>
                <span style="font-weight: 500;">{{ item.name }}</span>
              </div>
              <div v-if="item.description" style="margin-top: 4px; margin-left: 60px; font-size: 13px; color: #666;">
                {{ item.description }}
              </div>
              <div v-if="item.path" style="margin-top: 2px; margin-left: 60px; font-size: 12px; color: #999;">
                {{ item.path }}
              </div>
            </div>
          </NCard>
        </div>
      </div>
    </NSpin>
  </div>
</template>
