<script setup lang="ts">
import { h, computed } from 'vue'
import type { Component } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import {
  NLayout,
  NLayoutSider,
  NLayoutHeader,
  NLayoutContent,
  NMenu,
  NIcon,
  NButton,
  NAvatar,
  NDropdown,
  NSpace,
  NBreadcrumb,
  NBreadcrumbItem,
} from 'naive-ui'
import {
  SpeedometerOutline,
  FolderOutline,
  GitNetworkOutline,
  LocationOutline,
  LayersOutline,
  SwapHorizontalOutline,
  ServerOutline,
  GridOutline,
  BusinessOutline,
  RepeatOutline,
  GlobeOutline,
  ScanOutline,
  DocumentTextOutline,
  SearchOutline,
  PeopleOutline,
  ClipboardOutline,
  SettingsOutline,
  MenuOutline,
  LogOutOutline,
  LanguageOutline,
} from '@vicons/ionicons5'
import { useAppStore } from '@/stores/app'
import { useAuthStore } from '@/stores/auth'
import type { MenuOption } from 'naive-ui'

const router = useRouter()
const route = useRoute()
const { t, locale } = useI18n()
const appStore = useAppStore()
const authStore = useAuthStore()

// Render icon helper
function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) })
}

// Menu items with i18n
const menuOptions = computed<MenuOption[]>(() => [
  { label: t('menu.dashboard'), key: 'dashboard', icon: renderIcon(SpeedometerOutline) },
  { label: t('menu.sections'), key: 'sections', icon: renderIcon(FolderOutline) },
  { label: t('menu.subnets'), key: 'subnets', icon: renderIcon(GitNetworkOutline) },
  { label: t('menu.addresses'), key: 'addresses', icon: renderIcon(LocationOutline) },
  { label: t('menu.vlans'), key: 'vlans', icon: renderIcon(LayersOutline) },
  { label: t('menu.vrfs'), key: 'vrfs', icon: renderIcon(SwapHorizontalOutline) },
  { label: t('menu.devices'), key: 'devices', icon: renderIcon(ServerOutline) },
  { label: t('menu.racks'), key: 'racks', icon: renderIcon(GridOutline) },
  { label: t('menu.locations'), key: 'locations', icon: renderIcon(BusinessOutline) },
  { label: t('menu.nats'), key: 'nats', icon: renderIcon(RepeatOutline) },
  { label: t('menu.dns'), key: 'dns', icon: renderIcon(GlobeOutline) },
  { label: t('menu.scans'), key: 'scans', icon: renderIcon(ScanOutline) },
  { label: t('menu.requests'), key: 'requests', icon: renderIcon(DocumentTextOutline) },
  { label: t('menu.search'), key: 'search', icon: renderIcon(SearchOutline) },
  { label: t('menu.users'), key: 'users', icon: renderIcon(PeopleOutline) },
  { label: t('menu.audit'), key: 'audit', icon: renderIcon(ClipboardOutline) },
  { label: t('menu.settings'), key: 'settings', icon: renderIcon(SettingsOutline) },
])

// Active menu key based on current route
const activeKey = computed(() => {
  const path = route.path
  if (path.startsWith('/sections')) return 'sections'
  if (path.startsWith('/subnets')) return 'subnets'
  if (path.startsWith('/addresses')) return 'addresses'
  return (route.name as string) || 'dashboard'
})

// Handle menu selection
function handleMenuUpdate(key: string) {
  router.push({ name: key })
}

// Language switch
function toggleLocale() {
  const newLocale = appStore.locale === 'zh' ? 'en' : 'zh'
  appStore.setLocale(newLocale)
  locale.value = newLocale
}

const localeLabel = computed(() => appStore.locale === 'zh' ? 'EN' : '中')

// User dropdown options
const userDropdownOptions = computed(() => [
  { label: t('auth.logout'), key: 'logout', icon: renderIcon(LogOutOutline) },
])

function handleUserDropdown(key: string) {
  if (key === 'logout') {
    authStore.logout()
    router.push({ name: 'login' })
  }
}

// Breadcrumb
const breadcrumbTitle = computed(() => {
  const name = (route.name as string) || 'dashboard'
  const menuKey = `menu.${name}` as const
  return t(menuKey) || (route.meta.title as string) || 'GoIPAM'
})
</script>

<template>
  <NLayout has-sider style="height: 100vh">
    <!-- Sidebar -->
    <NLayoutSider
      bordered
      collapse-mode="width"
      :collapsed-width="64"
      :width="240"
      :collapsed="appStore.sidebarCollapsed"
      show-trigger
      @collapse="appStore.setSidebarCollapsed(true)"
      @expand="appStore.setSidebarCollapsed(false)"
      :native-scrollbar="false"
      style="height: 100vh"
    >
      <div class="sider-logo">
        <h1 v-if="!appStore.sidebarCollapsed" class="logo-text">GoIPAM</h1>
        <h1 v-else class="logo-text logo-text-mini">G</h1>
      </div>
      <NMenu
        :collapsed="appStore.sidebarCollapsed"
        :collapsed-width="64"
        :collapsed-icon-size="22"
        :options="menuOptions"
        :value="activeKey"
        @update:value="handleMenuUpdate"
      />
    </NLayoutSider>

    <NLayout>
      <!-- Top bar -->
      <NLayoutHeader bordered style="height: 56px; display: flex; align-items: center; padding: 0 24px; justify-content: space-between;">
        <NSpace align="center">
          <NButton quaternary @click="appStore.toggleSidebar">
            <template #icon>
              <NIcon><MenuOutline /></NIcon>
            </template>
          </NButton>
          <NBreadcrumb>
            <NBreadcrumbItem>GoIPAM</NBreadcrumbItem>
            <NBreadcrumbItem>{{ breadcrumbTitle }}</NBreadcrumbItem>
          </NBreadcrumb>
        </NSpace>

        <NSpace align="center">
          <NButton quaternary size="small" @click="toggleLocale" :title="appStore.locale === 'zh' ? 'Switch to English' : '切换为中文'">
            <template #icon>
              <NIcon><LanguageOutline /></NIcon>
            </template>
            {{ localeLabel }}
          </NButton>
          <NDropdown :options="userDropdownOptions" @select="handleUserDropdown">
            <NSpace align="center" style="cursor: pointer;">
              <NAvatar round size="small" style="background-color: #2080f0;">
                {{ authStore.user?.username?.charAt(0)?.toUpperCase() || 'U' }}
              </NAvatar>
              <span v-if="authStore.user">{{ authStore.user.username }}</span>
            </NSpace>
          </NDropdown>
        </NSpace>
      </NLayoutHeader>

      <!-- Content area -->
      <NLayoutContent
        content-style="padding: 24px;"
        :native-scrollbar="false"
        style="height: calc(100vh - 56px);"
      >
        <RouterView />
      </NLayoutContent>
    </NLayout>
  </NLayout>
</template>

<style scoped>
.sider-logo {
  height: 56px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-bottom: 1px solid var(--n-border-color);
}

.logo-text {
  font-size: 20px;
  font-weight: 700;
  color: #2080f0;
  white-space: nowrap;
  overflow: hidden;
}

.logo-text-mini {
  font-size: 24px;
}
</style>
