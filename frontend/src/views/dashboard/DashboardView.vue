<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import {
  NPageHeader,
  NCard,
  NGrid,
  NGi,
  NStatistic,
  NDataTable,
  NProgress,
  NSpin,
  useMessage,
} from 'naive-ui'
import { getDashboard } from '@/api/ipam'
import type { DashboardData } from '@/api/types'
import { formatDate, formatNumber } from '@/utils/format'

const { t } = useI18n()
const message = useMessage()
const loading = ref(false)
const dashboard = ref<DashboardData | null>(null)

// Recent subnets columns
const recentSubnetColumns = computed(() => [
  { title: 'CIDR', key: 'cidr', width: 160 },
  { title: t('common.name'), key: 'name', ellipsis: { tooltip: true } },
  { title: t('common.createdAt'), key: 'created_at', width: 180, render(row: any) { return formatDate(row.created_at) } },
])

// Recent changes columns
const recentChangeColumns = computed(() => [
  { title: 'IP', key: 'ip_text', width: 140 },
  { title: t('address.hostname'), key: 'hostname', width: 140, render(row: any) { return row.hostname || '-' } },
  { title: t('common.status'), key: 'status', width: 100 },
  { title: t('common.updatedAt'), key: 'updated_at', width: 180, render(row: any) { return formatDate(row.updated_at) } },
])

// Threshold subnets columns
const thresholdColumns = computed(() => [
  { title: 'CIDR', key: 'cidr', width: 160 },
  { title: t('common.name'), key: 'name', ellipsis: { tooltip: true } },
  { title: t('dashboard.addressUsage'), key: 'threshold_percent', width: 100 },
])

async function fetchData() {
  loading.value = true
  try {
    const res = await getDashboard()
    if (res.success && res.data) {
      dashboard.value = res.data
    }
  } catch (e: any) {
    message.error(e?.response?.data?.message || t('dashboard.loadFailed'))
  } finally {
    loading.value = false
  }
}

function getUsagePercent(): number {
  if (!dashboard.value) return 0
  const { used, available, free } = dashboard.value.address_usage
  const total = used + available + free
  return total > 0 ? (used / total) * 100 : 0
}

onMounted(fetchData)
</script>

<template>
  <div>
    <NPageHeader :title="t('dashboard.title')" :subtitle="t('dashboard.subtitle')" />

    <NSpin :show="loading">
      <!-- Stats Cards -->
      <NGrid :cols="4" :x-gap="12" :y-gap="12" style="margin-top: 16px;">
        <NGi>
          <NCard>
            <NStatistic :label="t('dashboard.totalSections')" :value="formatNumber(dashboard?.total_sections ?? 0)" />
          </NCard>
        </NGi>
        <NGi>
          <NCard>
            <NStatistic :label="t('dashboard.totalSubnets')" :value="formatNumber(dashboard?.total_subnets ?? 0)" />
          </NCard>
        </NGi>
        <NGi>
          <NCard>
            <NStatistic :label="t('dashboard.totalAddresses')" :value="formatNumber(dashboard?.total_addresses ?? 0)" />
          </NCard>
        </NGi>
        <NGi>
          <NCard>
            <NStatistic :label="t('dashboard.addressUsage')">
              <template #default>
                <div>
                  <div style="font-size: 14px; color: #666; margin-bottom: 8px;">
                    {{ t('dashboard.used') }}: {{ formatNumber(dashboard?.address_usage?.used ?? 0) }} /
                    {{ t('dashboard.available') }}: {{ formatNumber(dashboard?.address_usage?.available ?? 0) }} /
                    {{ t('dashboard.free') }}: {{ formatNumber(dashboard?.address_usage?.free ?? 0) }}
                  </div>
                  <NProgress
                    type="line"
                    :percentage="getUsagePercent()"
                    :status="getUsagePercent() > 90 ? 'error' : getUsagePercent() > 70 ? 'warning' : 'success'"
                  />
                </div>
              </template>
            </NStatistic>
          </NCard>
        </NGi>
      </NGrid>

      <!-- Recent Subnets & Recent Address Changes -->
      <NGrid :cols="2" :x-gap="12" style="margin-top: 16px;">
        <NGi>
          <NCard :title="t('dashboard.recentSubnets')">
            <NDataTable
              :columns="recentSubnetColumns"
              :data="dashboard?.recent_subnets ?? []"
              :bordered="false"
              size="small"
              :pagination="{ pageSize: 5 }"
            />
          </NCard>
        </NGi>
        <NGi>
          <NCard :title="t('dashboard.recentChanges')">
            <NDataTable
              :columns="recentChangeColumns"
              :data="dashboard?.recent_address_changes ?? []"
              :bordered="false"
              size="small"
              :pagination="{ pageSize: 5 }"
            />
          </NCard>
        </NGi>
      </NGrid>

      <!-- Threshold Subnets -->
      <NCard :title="t('dashboard.thresholdSubnets')" style="margin-top: 16px;">
        <div v-if="!dashboard?.threshold_subnets?.length" style="text-align: center; padding: 24px; color: #999;">
          {{ t('dashboard.noThreshold') }}
        </div>
        <NDataTable
          v-else
          :columns="thresholdColumns"
          :data="dashboard.threshold_subnets"
          :bordered="false"
          size="small"
        />
      </NCard>
    </NSpin>
  </div>
</template>
