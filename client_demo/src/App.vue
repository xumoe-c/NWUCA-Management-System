<template>
  <t-config-provider :global-config="globalConfig">
    <AppLayout />
  </t-config-provider>
</template>

<script setup lang="ts">
import AppLayout from './components/layout/AppLayout.vue';

const globalConfig = {
  brand: {
    color: '#2A82E4',
  },
};
</script>

<style>
html, body, #app {
  height: 100%;
  margin: 0;
}
</style>
