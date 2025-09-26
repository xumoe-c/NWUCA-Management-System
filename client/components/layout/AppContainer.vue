<template>
  <view class="app-container">
    <slot />
  </view>
</template>

<script>
export default { name: 'AppContainer' }
</script>

<style>
.app-container { padding: 16px; min-height: 100vh; background:#f5f6f8; }
</style>
