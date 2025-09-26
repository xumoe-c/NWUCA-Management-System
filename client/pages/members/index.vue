<template>
  <view class="page">
    <view class="header">
      <text class="h1">成员管理</text>
    </view>
    <view class="empty">查询区 + 表格 + 详情抽屉占位</view>
  </view>
</template>

<script>
import { isLoggedIn } from '../../store/auth'
export default { onShow(){ if(!isLoggedIn()) uni.reLaunch({ url:'/pages/login/index' }) } }
</script>

<style>
.page { padding: 16px; }
.header { margin-bottom: 12px; }
.h1 { font-size: 20px; font-weight: 600; }
.empty { padding: 24px; background: #fff; border-radius: 12px; color:#999; text-align:center; }
</style>
