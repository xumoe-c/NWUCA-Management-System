<template>
  <view class="page">
    <view class="header">
      <text class="h1">仪表盘</text>
    </view>
    <view class="grid">
      <view class="card" @click="go('/pages/members/index')">成员管理</view>
      <view class="card" @click="go('/pages/departments/index')">部门管理</view>
      <view class="card" @click="go('/pages/positions/index')">职务管理</view>
      <view class="card" @click="go('/pages/assignments/index')">任期分配</view>
      <view class="card" @click="go('/pages/profile/index')">个人中心</view>
    </view>
  </view>
</template>

<script>
import { isLoggedIn } from '../../store/auth'
export default {
  onShow() {
    if (!isLoggedIn()) {
      uni.reLaunch({ url: '/pages/login/index' })
    }
  },
  methods: {
    go(url){ uni.navigateTo({ url }) }
  }
}
</script>

<style>
.page { padding: 16px; }
.header { margin-bottom: 12px; }
.h1 { font-size: 20px; font-weight: 600; }
.grid { display: grid; grid-template-columns: repeat(2, 1fr); gap: 12px; }
.card { background: #fff; border-radius: 12px; padding: 18px; text-align: center; box-shadow: 0 2px 10px rgba(0,0,0,0.06); }
</style>
