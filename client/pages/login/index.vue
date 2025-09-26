<template>
  <view class="login-page">
    <view class="card">
      <view class="title">登录</view>
      <t-form :model="form" :rules="rules" ref="formRef" label-align="top">
        <t-form-item name="username" :label="''">
          <t-input v-model="form.username" placeholder="账号" clearable />
        </t-form-item>
        <t-form-item name="password" :label="''">
          <t-input v-model="form.password" type="password" placeholder="密码" clearable />
        </t-form-item>
        <view class="extra-row">
          <t-checkbox v-model="remember">记住我</t-checkbox>
        </view>
        <view class="submit-row">
          <t-button theme="primary" block :loading="loading" @click="onValidateAndSubmit">登录</t-button>
        </view>
      </t-form>
    </view>
  </view>
</template>

<script>
import { loginApi, getMeApi } from '../../api/auth'
import { setToken, setUser } from '../../store/auth'

// 说明：tdesign-uni 组件全局已按插件方式引入（见 uni_modules），无需单页重复注册。

export default {
  data() {
    return {
      form: { username: '', password: '' },
      remember: true,
      loading: false,
      rules: {
        username: [{ required: true, message: '请输入账号' }],
        password: [{ required: true, message: '请输入密码' }],
      }
    }
  },
  methods: {
    async onValidateAndSubmit(){
      const formRef = this.$refs.formRef
      if(!formRef || !formRef.validate){
        // 组件未就绪时直接兜底校验
        if(!this.form.username || !this.form.password){
          return uni.showToast({ title:'请输入账号和密码', icon:'none' })
        }
      } else {
        const { isValid } = await formRef.validate({ trigger:'all' })
        if(!isValid) return
      }
      this.loading = true
      try {
        const res = await loginApi({ username: this.form.username, password: this.form.password })
        const token = (res && (res.data && (res.data.token || res.data.access_token))) || res.token || res.access_token
        if (!token) throw new Error('无效的登录响应')
        setToken(token)
        if (this.remember) { /* 可扩展本地持久策略 */ }
        try {
          const me = await getMeApi()
          setUser(me.data || me)
        } catch (e) {}
        uni.showToast({ title: '登录成功', icon: 'success' })
        setTimeout(() => uni.reLaunch({ url: '/pages/dashboard/index' }), 250)
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
.extra-row { margin: 4px 0 12px; display:flex; justify-content: space-between; }
.submit-row { margin-top: 4px; }
</style>
