<script setup lang="ts">
import { computed } from 'vue'
import { NConfigProvider, NMessageProvider, NDialogProvider, zhCN, dateZhCN, enUS, dateEnUS } from 'naive-ui'
import { useAppStore } from '@/stores/app'

const appStore = useAppStore()

const naiveLocale = computed(() => appStore.locale === 'zh' ? zhCN : enUS)
const naiveDateLocale = computed(() => appStore.locale === 'zh' ? dateZhCN : dateEnUS)
</script>

<template>
  <NConfigProvider :locale="naiveLocale" :date-locale="naiveDateLocale">
    <NDialogProvider>
      <NMessageProvider>
        <RouterView />
      </NMessageProvider>
    </NDialogProvider>
  </NConfigProvider>
</template>
