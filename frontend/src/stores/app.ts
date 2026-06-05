import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAppStore = defineStore('app', () => {
  // State
  const sidebarCollapsed = ref(false)
  const locale = ref<string>(localStorage.getItem('goipam_locale') || 'zh')
  const timezone = ref<string>(
    localStorage.getItem('goipam_timezone') ||
      Intl.DateTimeFormat().resolvedOptions().timeZone ||
      'Asia/Shanghai'
  )

  // Actions
  function toggleSidebar(): void {
    sidebarCollapsed.value = !sidebarCollapsed.value
  }

  function setSidebarCollapsed(collapsed: boolean): void {
    sidebarCollapsed.value = collapsed
  }

  function setLocale(val: string): void {
    locale.value = val
    localStorage.setItem('goipam_locale', val)
  }

  function setTimezone(val: string): void {
    timezone.value = val
    localStorage.setItem('goipam_timezone', val)
  }

  return {
    sidebarCollapsed,
    locale,
    timezone,
    toggleSidebar,
    setSidebarCollapsed,
    setLocale,
    setTimezone,
  }
})
