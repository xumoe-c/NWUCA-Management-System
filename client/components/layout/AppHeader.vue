<template>
  <view class="app-header">
    <text class="title">{{ title }}</text>
    <slot />
  </view>
</template>

<script>
export default {
  name: 'AppHeader',
  props: { title: { type: String, default: '' } }
}
</script>

<style>
.app-header { display:flex; align-items:center; justify-content:space-between; padding:12px 16px; background:#fff; border-bottom:1px solid #f0f0f0; }
.title { font-size: 18px; font-weight: 600; }
</style>
