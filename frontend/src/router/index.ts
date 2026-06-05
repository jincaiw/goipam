import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

// Layout
import MainLayout from '@/layouts/MainLayout.vue'

// Views
import LoginView from '@/views/auth/LoginView.vue'
import DashboardView from '@/views/dashboard/DashboardView.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'login',
    component: LoginView,
    meta: { requiresAuth: false },
  },
  {
    path: '/',
    component: MainLayout,
    meta: { requiresAuth: true },
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'dashboard',
        component: DashboardView,
        meta: { title: 'Dashboard', icon: 'SpeedometerOutline' },
      },
      {
        path: 'sections',
        name: 'sections',
        component: () => import('@/views/sections/SectionsView.vue'),
        meta: { title: 'Sections', icon: 'FolderOutline' },
      },
      {
        path: 'sections/:id',
        name: 'section-detail',
        component: () => import('@/views/sections/SectionDetail.vue'),
        meta: { title: 'Section Detail', icon: 'FolderOutline' },
      },
      {
        path: 'subnets',
        name: 'subnets',
        component: () => import('@/views/subnets/SubnetsView.vue'),
        meta: { title: 'Subnets', icon: 'GitNetworkOutline' },
      },
      {
        path: 'subnets/:id',
        name: 'subnet-detail',
        component: () => import('@/views/subnets/SubnetDetail.vue'),
        meta: { title: 'Subnet Detail', icon: 'GitNetworkOutline' },
      },
      {
        path: 'addresses',
        name: 'addresses',
        component: () => import('@/views/addresses/AddressesView.vue'),
        meta: { title: 'Addresses', icon: 'LocationOutline' },
      },
      {
        path: 'vlans',
        name: 'vlans',
        component: () => import('@/views/vlans/VlansView.vue'),
        meta: { title: 'VLANs', icon: 'LayersOutline' },
      },
      {
        path: 'vrfs',
        name: 'vrfs',
        component: () => import('@/views/vrfs/VrfsView.vue'),
        meta: { title: 'VRFs', icon: 'SwapHorizontalOutline' },
      },
      {
        path: 'devices',
        name: 'devices',
        component: () => import('@/views/devices/DevicesView.vue'),
        meta: { title: 'Devices', icon: 'ServerOutline' },
      },
      {
        path: 'racks',
        name: 'racks',
        component: () => import('@/views/racks/RacksView.vue'),
        meta: { title: 'Racks', icon: 'GridOutline' },
      },
      {
        path: 'locations',
        name: 'locations',
        component: () => import('@/views/locations/LocationsView.vue'),
        meta: { title: 'Locations', icon: 'BusinessOutline' },
      },
      {
        path: 'nats',
        name: 'nats',
        component: () => import('@/views/nats/NatsView.vue'),
        meta: { title: 'NATs', icon: 'RepeatOutline' },
      },
      {
        path: 'dns',
        name: 'dns',
        component: () => import('@/views/dns/DnsView.vue'),
        meta: { title: 'DNS', icon: 'GlobeOutline' },
      },
      {
        path: 'scans',
        name: 'scans',
        component: () => import('@/views/scans/ScansView.vue'),
        meta: { title: 'Scans', icon: 'ScanOutline' },
      },
      {
        path: 'requests',
        name: 'requests',
        component: () => import('@/views/requests/RequestsView.vue'),
        meta: { title: 'Requests', icon: 'DocumentTextOutline' },
      },
      {
        path: 'search',
        name: 'search',
        component: () => import('@/views/search/SearchView.vue'),
        meta: { title: 'Search', icon: 'SearchOutline' },
      },
      {
        path: 'users',
        name: 'users',
        component: () => import('@/views/users/UsersView.vue'),
        meta: { title: 'Users & Roles', icon: 'PeopleOutline' },
      },
      {
        path: 'audit',
        name: 'audit',
        component: () => import('@/views/audit/AuditView.vue'),
        meta: { title: 'Audit', icon: 'ClipboardOutline' },
      },
      {
        path: 'settings',
        name: 'settings',
        component: () => import('@/views/settings/SettingsView.vue'),
        meta: { title: 'Settings', icon: 'SettingsOutline' },
      },
    ],
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'not-found',
    redirect: '/dashboard',
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

// Navigation guard: check authentication
router.beforeEach((to, _from, next) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth !== false && !authStore.isAuthenticated) {
    next({ name: 'login', query: { redirect: to.fullPath } })
  } else if (to.name === 'login' && authStore.isAuthenticated) {
    next({ name: 'dashboard' })
  } else {
    next()
  }
})

export default router
