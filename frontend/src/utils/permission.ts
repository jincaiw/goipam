import { useAuthStore } from '@/stores/auth'

/**
 * Check if the current user has a specific role
 */
export function hasRole(role: string): boolean {
  const authStore = useAuthStore()
  return authStore.hasRole(role)
}

/**
 * Check if the current user has a specific permission
 */
export function hasPermission(permission: string): boolean {
  const authStore = useAuthStore()
  return authStore.hasPermission(permission)
}

/**
 * Check if the current user is a superuser
 */
export function isSuperuser(): boolean {
  const authStore = useAuthStore()
  return authStore.isSuperuser
}

/**
 * Check if the user is authenticated
 */
export function isAuthenticated(): boolean {
  const authStore = useAuthStore()
  return authStore.isAuthenticated
}

/**
 * Check all permissions (user must have ALL of the specified permissions)
 */
export function hasAllPermissions(...permissions: string[]): boolean {
  return permissions.every(hasPermission)
}

/**
 * Check any permission (user must have at least ONE of the specified permissions)
 */
export function hasAnyPermission(...permissions: string[]): boolean {
  return permissions.some(hasPermission)
}
