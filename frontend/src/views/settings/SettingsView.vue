<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import {
  NPageHeader,
  NCard,
  NForm,
  NFormItem,
  NInput,
  NInputNumber,
  NButton,
  NSelect,
  NSpace,
  useMessage,
} from 'naive-ui'
import { useAppStore } from '@/stores/app'
import { getSettings, saveSettings } from '@/api/ipam'

const { t, locale } = useI18n()
const appStore = useAppStore()
const message = useMessage()

// Language options
const languageOptions = [
  { label: '中文', value: 'zh' },
  { label: 'English', value: 'en' },
]

// Timezone options
const timezoneOptions = [
  { label: 'UTC', value: 'UTC' },
  { label: 'Asia/Shanghai (CST +8)', value: 'Asia/Shanghai' },
  { label: 'Asia/Tokyo (JST +9)', value: 'Asia/Tokyo' },
  { label: 'Asia/Singapore (SGT +8)', value: 'Asia/Singapore' },
  { label: 'Asia/Kolkata (IST +5:30)', value: 'Asia/Kolkata' },
  { label: 'America/New_York (EST -5)', value: 'America/New_York' },
  { label: 'America/Chicago (CST -6)', value: 'America/Chicago' },
  { label: 'America/Los_Angeles (PST -8)', value: 'America/Los_Angeles' },
  { label: 'Europe/London (GMT +0)', value: 'Europe/London' },
  { label: 'Europe/Berlin (CET +1)', value: 'Europe/Berlin' },
  { label: 'Europe/Moscow (MSK +3)', value: 'Europe/Moscow' },
  { label: 'Australia/Sydney (AEST +10)', value: 'Australia/Sydney' },
]

// General settings
const generalSettings = ref({
  siteName: 'GoIPAM',
  defaultPrefix: '24',
})

// Scan settings
const scanSettings = ref({
  defaultPorts: '22,80,443',
  timeout: 3000,
})

// API Token settings
const tokenSettings = ref({
  tokenName: '',
  tokenExpiry: 365,
})

// Load settings from backend
async function loadSettings() {
  try {
    const resp = await getSettings('general')
    if (resp.data) {
      generalSettings.value.siteName = resp.data.site_name || 'GoIPAM'
      generalSettings.value.defaultPrefix = resp.data.default_prefix || '24'
    }
  } catch {
    // Use defaults
  }
  try {
    const resp = await getSettings('scan')
    if (resp.data) {
      scanSettings.value.defaultPorts = resp.data.default_ports || '22,80,443'
      scanSettings.value.timeout = resp.data.timeout ? parseInt(resp.data.timeout) : 3000
    }
  } catch {
    // Use defaults
  }
}

onMounted(loadSettings)

// Handlers
function handleLanguageChange(val: string) {
  appStore.setLocale(val)
  locale.value = val
}

function handleTimezoneChange(val: string) {
  appStore.setTimezone(val)
}

async function handleSaveGeneral() {
  try {
    await saveSettings('general', {
      site_name: generalSettings.value.siteName,
      default_prefix: generalSettings.value.defaultPrefix,
    })
    message.success(t('settings.saved'))
  } catch {
    message.error(t('settings.saveFailed'))
  }
}

async function handleSaveScan() {
  try {
    await saveSettings('scan', {
      default_ports: scanSettings.value.defaultPorts,
      timeout: String(scanSettings.value.timeout),
    })
    message.success(t('settings.saved'))
  } catch {
    message.error(t('settings.saveFailed'))
  }
}

function handleGenerateToken() {
  message.success(t('settings.tokenGenerated'))
}
</script>

<template>
  <div>
    <NPageHeader :title="t('settings.title')" :subtitle="t('settings.subtitle')" />

    <!-- Language & Timezone Settings -->
    <NCard :title="t('settings.general')" style="margin-top: 16px;">
      <NForm label-placement="left" label-width="160">
        <NFormItem :label="t('settings.language')">
          <NSelect
            :value="appStore.locale"
            :options="languageOptions"
            @update:value="handleLanguageChange"
            style="width: 240px;"
          />
        </NFormItem>
        <NFormItem :label="t('settings.timezone')">
          <NSelect
            :value="appStore.timezone"
            :options="timezoneOptions"
            @update:value="handleTimezoneChange"
            style="width: 320px;"
          />
        </NFormItem>
        <NFormItem :label="t('settings.siteName')">
          <NInput v-model:value="generalSettings.siteName" :placeholder="t('settings.siteName')" />
        </NFormItem>
        <NFormItem :label="t('settings.defaultPrefix')">
          <NInput v-model:value="generalSettings.defaultPrefix" :placeholder="t('settings.defaultPrefix')" />
        </NFormItem>
      </NForm>
      <template #action>
        <NButton type="primary" @click="handleSaveGeneral">{{ t('settings.saveSettings') }}</NButton>
      </template>
    </NCard>

    <!-- Scan Settings -->
    <NCard :title="t('settings.scanSettings')" style="margin-top: 16px;">
      <NForm label-placement="left" label-width="160">
        <NFormItem :label="t('settings.defaultPorts')">
          <NInput v-model:value="scanSettings.defaultPorts" placeholder="22,80,443" />
        </NFormItem>
        <NFormItem :label="t('settings.timeout')">
          <NInputNumber v-model:value="scanSettings.timeout" :min="100" :max="60000" style="width: 100%;" />
        </NFormItem>
      </NForm>
      <template #action>
        <NButton type="primary" @click="handleSaveScan">{{ t('settings.saveSettings') }}</NButton>
      </template>
    </NCard>

    <!-- API Token Management -->
    <NCard :title="t('settings.apiTokens')" style="margin-top: 16px;">
      <NForm label-placement="left" label-width="160">
        <NFormItem :label="t('settings.tokenName')">
          <NInput v-model:value="tokenSettings.tokenName" :placeholder="t('settings.tokenName')" />
        </NFormItem>
        <NFormItem :label="t('settings.expiry')">
          <NInputNumber v-model:value="tokenSettings.tokenExpiry" :min="1" :max="3650" style="width: 100%;" />
        </NFormItem>
      </NForm>
      <template #action>
        <NSpace>
          <NButton type="primary" @click="handleGenerateToken">{{ t('settings.generateToken') }}</NButton>
        </NSpace>
      </template>
    </NCard>
  </div>
</template>
