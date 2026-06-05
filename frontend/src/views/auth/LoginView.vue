<script setup lang="ts">
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { NCard, NForm, NFormItem, NInput, NButton, useMessage } from 'naive-ui'
import type { FormInst, FormRules } from 'naive-ui'
import { useAuthStore } from '@/stores/auth'
import { useI18n } from 'vue-i18n'

const router = useRouter()
const route = useRoute()
const message = useMessage()
const authStore = useAuthStore()
const { t } = useI18n()

const formRef = ref<FormInst | null>(null)
const loading = ref(false)

const formValue = ref({
  username: '',
  password: '',
})

const rules: FormRules = {
  username: [
    { required: true, message: () => t('common.required'), trigger: 'blur' },
  ],
  password: [
    { required: true, message: () => t('common.required'), trigger: 'blur' },
  ],
}

async function handleLogin() {
  try {
    await formRef.value?.validate()
  } catch {
    return
  }

  loading.value = true
  try {
    await authStore.login({
      username: formValue.value.username,
      password: formValue.value.password,
    })
    message.success(t('auth.loginSuccess'))
    const redirect = (route.query.redirect as string) || '/dashboard'
    router.push(redirect)
  } catch (error: unknown) {
    const err = error as { response?: { data?: { message?: string } } }
    message.error(err.response?.data?.message || t('auth.loginFailed'))
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="login-container">
    <NCard class="login-card" title="GoIPAM" :bordered="false">
      <template #header-extra>
        <span class="login-subtitle">{{ t('auth.subtitle') }}</span>
      </template>
      <NForm
        ref="formRef"
        :model="formValue"
        :rules="rules"
        label-placement="left"
        :show-label="false"
      >
        <NFormItem path="username" :label="t('auth.username')">
          <NInput
            v-model:value="formValue.username"
            :placeholder="t('auth.username')"
            size="large"
            @keyup.enter="handleLogin"
          />
        </NFormItem>
        <NFormItem path="password" :label="t('auth.password')">
          <NInput
            v-model:value="formValue.password"
            type="password"
            show-password-on="click"
            :placeholder="t('auth.password')"
            size="large"
            @keyup.enter="handleLogin"
          />
        </NFormItem>
        <NFormItem>
          <NButton
            type="primary"
            block
            size="large"
            :loading="loading"
            @click="handleLogin"
          >
            {{ t('auth.signIn') }}
          </NButton>
        </NFormItem>
      </NForm>
    </NCard>
  </div>
</template>

<style scoped>
.login-container {
  width: 100%;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.login-card {
  width: 400px;
  max-width: 90vw;
}

.login-subtitle {
  font-size: 14px;
  color: #999;
}
</style>
