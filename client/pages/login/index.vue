<template>
  <view class="login-page">
    <view class="card">
      <view class="title">登录</view>
      <view class="form-item">
        <input class="input" type="text" v-model="username" placeholder="账号" />
      </view>
      <view class="form-item">
        <input class="input" type="password" v-model="password" placeholder="密码" />
      </view>
      <button class="btn" :disabled="loading" @click="onSubmit">{{ loading ? '登录中...' : '登录' }}</button>
    </view>
  </view>
  
</template>

<script>
import { loginApi, getMeApi } from '../../api/auth'
import { setToken, setUser } from '../../store/auth'

export default {
  data() {
    return { username: '', password: '', loading: false }
  },
  methods: {
    async onSubmit() {
      if (!this.username || !this.password) {
        uni.showToast({ title: '请输入账号和密码', icon: 'none' })
        return
      }
      this.loading = true
      try {
        const res = await loginApi({ username: this.username, password: this.password })
        //兼容统一响应或直接返回token
        const token = (res && (res.data && (res.data.token || res.data.access_token))) || res.token || res.access_token
        if (!token) throw new Error('无效的登录响应')
        setToken(token)
        // 拉取个人信息（容错）
        try {
          const me = await getMeApi()
          setUser(me.data || me)
        } catch (e) {}
        uni.showToast({ title: '登录成功', icon: 'success' })
        setTimeout(() => {
          uni.reLaunch({ url: '/pages/dashboard/index' })
        }, 300)
      } catch (e) {
        const msg = (e && e.data && e.data.message) || e.message || '登录失败，请检查账号或密码'
        uni.showToast({ title: msg, icon: 'none' })
      } finally {
        this.loading = false
      }
    }
  }
}
</script>

<style>
.login-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f5f6f8;
}
.card {
  width: 86vw;
  max-width: 420px;
  background: #fff;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 8px 24px rgba(0,0,0,0.08);
}
.title { font-size: 20px; font-weight: 600; margin-bottom: 16px; }
.form-item { margin-bottom: 12px; }
.input { width: 100%; border: 1px solid #e5e6eb; border-radius: 8px; padding: 10px 12px; }
.btn { width: 100%; background: #2A82E4; color: #fff; border: none; padding: 10px 12px; border-radius: 8px; }
</style>
