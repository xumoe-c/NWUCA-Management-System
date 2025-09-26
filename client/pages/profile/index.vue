<template>
  <view class="page">
    <view class="header">
      <text class="h1">个人中心</text>
    </view>
    <view class="card">
      <view class="row"><text class="label">姓名</text><text class="val">{{ user && (user.name || user.username) }}</text></view>
      <view class="row"><text class="label">角色</text><text class="val">{{ user && (user.role || user.roles && user.roles.join(',')) }}</text></view>
      <button class="btn" @click="logout">退出登录</button>
    </view>
  </view>
</template>

<script>
import { isLoggedIn, getUser, clearToken, clearUser } from '../../store/auth'
export default {
  data(){ return { user: null } },
  onShow(){
    if(!isLoggedIn()) return uni.reLaunch({ url:'/pages/login/index' })
    this.user = getUser()
  },
  methods:{
    logout(){
      clearToken(); clearUser();
      uni.reLaunch({ url: '/pages/login/index' })
    }
  }
}
</script>

<style>
.page { padding: 16px; }
.header { margin-bottom: 12px; }
.h1 { font-size: 20px; font-weight: 600; }
.card { background:#fff; border-radius:12px; padding:16px; }
.row { display:flex; justify-content:space-between; padding:8px 0; }
.label { color:#666; }
.val { color:#111; }
.btn { margin-top: 12px; background:#e34d59; color:#fff; border-radius:8px; }
</style>
