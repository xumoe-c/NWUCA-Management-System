<template>
  <div class="login-page">
    <t-card class="login-card" :bordered="false">
      <h2 class="title">NWUCA 登录</h2>
      <t-form @submit.prevent="onSubmit" :data="form" :rules="rules" label-width="0">
        <t-form-item name="email">
          <t-input v-model="form.email" size="large" placeholder="邮箱" clearable />
        </t-form-item>
        <t-form-item name="password">
          <t-input v-model="form.password" size="large" type="password" placeholder="密码" clearable />
        </t-form-item>
        <t-form-item>
          <!-- 兼容某些场景下 submit 事件未触发，直接绑定点击事件 -->
          <t-button :loading="loading" theme="primary" type="submit" block @click="onSubmit">登录</t-button>
        </t-form-item>
        <div class="swap">
          还没有账号？<t-link theme="primary" @click="goRegister">去注册</t-link>
        </div>
      </t-form>
    </t-card>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import { MessagePlugin } from 'tdesign-vue-next';
import { useRouter, useRoute } from 'vue-router';
import { loginApi } from '../api/auth';
import { useAuthStore } from '../store/auth';

const router = useRouter();
const route = useRoute();
const auth = useAuthStore();

const form = reactive({ email: '', password: '' });
const loading = ref(false);

const rules = {
  email: [
    { required: true, message: '请输入邮箱' },
    { email: true, message: '邮箱格式不正确' },
  ],
  password: [{ required: true, message: '请输入密码' }],
};

async function onSubmit() {
  loading.value = true;
  try {
    const res = await loginApi({ email: form.email, password: form.password });
    const token = res.data?.token;
    if (token) {
      // 后端当前未返回角色，默认 member；后续可由 /users/me 获取
      auth.setAuth(token, 'member');
      MessagePlugin.success('登录成功');
      const redirect = (route.query.redirect as string) || '/dashboard';
      router.replace(redirect);
    } else {
      MessagePlugin.error(res.msg || '登录失败');
    }
  } catch (e: any) {
    const msg = e?.response?.data?.msg || '登录失败';
    MessagePlugin.error(msg);
  } finally {
    loading.value = false;
  }
}

function goRegister() {
  router.push('/register');
}
</script>

<style scoped>
.login-page {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--td-bg-color-page);
}
.login-card { width: 360px; }
.title { text-align: center; margin: 0 0 16px; }
.swap { text-align: center; }
</style>
