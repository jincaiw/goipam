import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  build: {
    rollupOptions: {
      output: {
        manualChunks(id) {
          if (!id.includes('node_modules')) {
            return undefined
          }
          if (id.includes('/@css-render/') || id.includes('/date-fns/') || id.includes('/seemly/') || id.includes('/vooks/') || id.includes('/vueuc/')) {
            return 'vendor-naive-runtime'
          }
          if (id.includes('/naive-ui/')) {
            if (id.includes('/_internal/') || id.includes('/_utils/') || id.includes('/_styles/') || id.includes('/_mixins/') || id.includes('/config-provider/') || id.includes('/message/') || id.includes('/dialog/')) {
              return 'vendor-naive-core'
            }
            if (id.includes('/data-table/') || id.includes('/date-picker/') || id.includes('/modal/') || id.includes('/drawer/') || id.includes('/popconfirm/') || id.includes('/tooltip/') || id.includes('/popover/') || id.includes('/select/') || id.includes('/dropdown/') || id.includes('/menu/')) {
              return 'vendor-naive-data'
            }
            if (id.includes('/form/') || id.includes('/input/') || id.includes('/input-number/') || id.includes('/switch/') || id.includes('/checkbox/') || id.includes('/radio/') || id.includes('/button/') || id.includes('/button-group/')) {
              return 'vendor-naive-controls'
            }
            if (id.includes('/layout/') || id.includes('/card/') || id.includes('/grid/') || id.includes('/descriptions/') || id.includes('/statistic/') || id.includes('/progress/') || id.includes('/tag/') || id.includes('/avatar/') || id.includes('/breadcrumb/') || id.includes('/page-header/') || id.includes('/space/')) {
              return 'vendor-naive-display'
            }
            return 'vendor-naive-misc'
          }
          if (id.includes('/echarts/') || id.includes('/vue-echarts/') || id.includes('/zrender/')) {
            return 'vendor-charts'
          }
          if (id.includes('/vue/') || id.includes('/vue-router/') || id.includes('/pinia/') || id.includes('/vue-i18n/')) {
            return 'vendor-vue'
          }
          if (id.includes('/axios/')) {
            return 'vendor-axios'
          }
          return 'vendor'
        },
      },
    },
  },
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src'),
    },
  },
  server: {
    port: 5070,
    proxy: {
      '/api': {
        target: 'http://localhost:5050',
        changeOrigin: true,
      },
    },
  },
})
