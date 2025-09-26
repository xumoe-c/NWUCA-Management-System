<template>
  <view class="page">
    <view class="header">
      <text class="h1">仪表盘</text>
    </view>
    <t-row gutter="12">
      <t-col :span="6" v-for="item in entries" :key="item.path">
        <view class="card" @click="go(item.path)">
          <view class="card-title">{{ item.title }}</view>
          <view class="card-desc">{{ item.desc }}</view>
        </view>
      </t-col>
    </t-row>
    <view class="section">
      <view class="section-title">概览 / 占位</view>
      <view class="placeholder">后续接入：最近创建成员 / 操作日志 / 待办事项</view>
    </view>
  </view>
</template>

<script>
import { isLoggedIn } from '../../store/auth'
export default {
  data(){
    return {
      entries: [
        { title: '成员管理', path: '/pages/members/index', desc: '成员档案与查询' },
        { title: '部门管理', path: '/pages/departments/index', desc: '组织结构树' },
        { title: '职务管理', path: '/pages/positions/index', desc: '职务与级别' },
        { title: '任期分配', path: '/pages/assignments/index', desc: '成员任期记录' },
        { title: '个人中心', path: '/pages/profile/index', desc: '账户与安全' },
      ]
    }
  },
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
.card { background: #fff; border-radius: 14px; padding: 16px; box-shadow: 0 2px 10px rgba(0,0,0,0.06); margin-bottom:12px; min-height: 96px; display:flex; flex-direction:column; justify-content:center; }
.card:active { background:#f5f7fa; }
.card-title { font-size:16px; font-weight:600; margin-bottom:4px; }
.card-desc { font-size:12px; color:#666; }
.section { margin-top: 24px; }
.section-title { font-size: 14px; font-weight: 600; margin-bottom:8px; }
.placeholder { background:#fff; padding:16px; border-radius:12px; color:#999; font-size:12px; }
</style>
