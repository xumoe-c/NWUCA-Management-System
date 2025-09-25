<template>
  <div class="auth-page">
    <t-card class="auth-card" :bordered="false">
      <h2 class="title">注册账号</h2>
      <t-form ref="formRef" :data="form" :rules="rules" label-width="0" @submit.prevent="onSubmit">
        <t-form-item name="username">
          <t-input v-model="form.username" size="large" placeholder="用户名（必填）" clearable />
        </t-form-item>
        <t-form-item name="email">
          <t-input v-model="form.email" size="large" placeholder="邮箱（必填）" clearable />
        </t-form-item>
        <t-form-item name="password">
          <t-input v-model="form.password" size="large" type="password" placeholder="密码（至少 8 位）" clearable />
        </t-form-item>
        <t-form-item name="confirm">
          <t-input v-model="form.confirm" size="large" type="password" placeholder="确认密码" clearable />
        </t-form-item>
        <t-form-item>
          <t-button :loading="loading" theme="primary" type="submit" block>注册</t-button>
        </t-form-item>
        <div class="swap">
          已有账号？<t-link theme="primary" @click="goLogin">去登录</t-link>
        </div>
      </t-form>
    </t-card>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import { MessagePlugin } from 'tdesign-vue-next';
import { useRouter, useRoute } from 'vue-router';
import { registerApi } from '../api/auth';

const router = useRouter();
const route = useRoute();

const formRef = ref();
const form = reactive({ username: '', email: '', password: '', confirm: '' });
const loading = ref(false);

const rules = {
  username: [{ required: true, message: '请输入用户名' }],
  email: [
    { required: true, message: '请输入邮箱' },
    { email: true, message: '邮箱格式不正确' },
  ],
  password: [
    { required: true, message: '请输入密码' },
    { min: 8, message: '至少 8 位' },
  ],
  confirm: [
    {
      validator: (val: string) => val === form.password,
      message: '两次输入的密码不一致',
      trigger: 'change',
    },
  ],
};

async function onSubmit() {
  // 先运行表单校验，确保必填/格式/一致性都被校验
  const result = await formRef.value?.validate?.();
  if (result !== true) return;

  // 双重保险（即便未触发表单校验，这里也会拦截）
  if (form.confirm !== form.password) {
    MessagePlugin.error('两次输入的密码不一致');
    return;
  }

  loading.value = true;
  try {
    const res = await registerApi({ username: form.username, email: form.email, password: form.password });
    if (res.code === 201 || res.code === 200) {
      MessagePlugin.success('注册成功，请登录');
      const redirect = (route.query.redirect as string) || '/login';
      router.replace(redirect);
    } else {
      MessagePlugin.error(res.msg || '注册失败');
    }
  } catch (e: any) {
    const msg = e?.response?.data?.msg || '注册失败';
    MessagePlugin.error(msg);
  } finally {
    loading.value = false;
  }
}

function goLogin() {
  router.push('/login');
}
</script>

<style scoped>
.auth-page {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--td-bg-color-page);
}
.auth-card { width: 420px; }
.title { text-align: center; margin: 0 0 16px; }
.swap { text-align: center; color: var(--td-text-color-secondary); }
</style>
