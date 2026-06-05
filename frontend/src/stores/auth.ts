import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { UserInfo } from '@/api/types'
import { post } from '@/api'
import type { LoginRequest, LoginResponse } from '@/api/types'

const TOKEN_KEY = 'goipam_token'
const USER_KEY = 'goipam_user'

export const useAuthStore = defineStore('auth', () => {
  // State
  const token = ref<string>(localStorage.getItem(TOKEN_KEY) || '')
  const user = ref<UserInfo | null>(
    JSON.parse(localStorage.getItem(USER_KEY) || 'null'),
  )

  // Getters
  const isAuthenticated = computed(() => !!token.value)
  const isSuperuser = computed(() => user.value?.is_superuser ?? false)
  const userRoles = computed(() => user.value?.roles ?? [])

  // Actions
  async function login(credentials: LoginRequest): Promise<void> {
    const response = await post<LoginResponse>('/auth/login', credentials)
    token.value = response.data.access_token
    user.value = response.data.user
    localStorage.setItem(TOKEN_KEY, response.data.access_token)
    localStorage.setItem(USER_KEY, JSON.stringify(response.data.user))
  }

  function logout(): void {
    token.value = ''
    user.value = null
    localStorage.removeItem(TOKEN_KEY)
    localStorage.removeItem(USER_KEY)
  }

  function hasRole(role: string): boolean {
    return userRoles.value.includes(role)
  }

  function hasPermission(permission: string): boolean {
    if (isSuperuser.value) return true
    return userRoles.value.includes(permission)
  }

  return {
    token,
    user,
    isAuthenticated,
    isSuperuser,
    userRoles,
    login,
    logout,
    hasRole,
    hasPermission,
  }
})
