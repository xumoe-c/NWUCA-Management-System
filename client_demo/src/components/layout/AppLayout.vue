<template>
  <t-layout class="h100">
    <t-aside width="232px">
      <div class="logo">NWUCA 管理系统</div>
      <t-menu theme="light" :value="active" @change="onMenuChange">
        <t-menu-item value="dashboard">
          <t-icon name="dashboard" /> 仪表盘
        </t-menu-item>
        <t-submenu value="org" title="组织管理">
          <t-menu-item value="departments">部门管理</t-menu-item>
          <t-menu-item value="positions">职务管理</t-menu-item>
          <t-menu-item value="members">成员管理</t-menu-item>
          <t-menu-item value="assignments">任期分配</t-menu-item>
        </t-submenu>
        <t-menu-item value="profile">
          <t-icon name="user-circle" /> 个人中心
        </t-menu-item>
      </t-menu>
    </t-aside>

    <t-layout>
      <t-header>
        <div class="header-left">
          <t-breadcrumb>
            <t-breadcrumbItem>NWUCA</t-breadcrumbItem>
            <t-breadcrumbItem>{{ breadcrumb }}</t-breadcrumbItem>
          </t-breadcrumb>
        </div>
        <div class="header-right">
          <t-button theme="default" variant="outline" size="small" @click="toggleDark">主题</t-button>
          <t-dropdown :options="userMenu" @click="onUserAction">
            <t-button shape="round" variant="text">
              <t-icon name="user-circle" />
            </t-button>
          </t-dropdown>
        </div>
      </t-header>

      <t-content>
        <router-view />
      </t-content>
    </t-layout>
  </t-layout>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { MenuValue } from 'tdesign-vue-next';

const router = useRouter();
const route = useRoute();

const active = ref<MenuValue>('dashboard');

const routeMap: Record<string, string> = {
  dashboard: '/dashboard',
  departments: '/departments',
  positions: '/positions',
  members: '/members',
  assignments: '/assignments',
  profile: '/profile',
};

const breadcrumb = computed(() => {
  const map: Record<string, string> = {
    dashboard: '仪表盘',
    departments: '部门管理',
    positions: '职务管理',
    members: '成员管理',
    assignments: '任期分配',
    profile: '个人中心',
  };
  const key = Object.keys(routeMap).find((k) => route.path.startsWith(routeMap[k])) || 'dashboard';
  active.value = key as MenuValue;
  return map[key];
});

function onMenuChange(v: MenuValue) {
  const path = routeMap[String(v)];
  if (path) router.push(path);
}

const userMenu = [
  { content: '个人中心', value: 'profile' },
  { content: '退出登录', value: 'logout' },
];

function onUserAction({ value }: { value: string }) {
  if (value === 'profile') router.push('/profile');
  if (value === 'logout') {
    localStorage.removeItem('token');
    router.push('/login');
  }
}

function toggleDark() {
  document.documentElement.classList.toggle('tdesign-theme--dark');
}
</script>

<style scoped>
.h100 { height: 100vh; }
.logo {
  height: 56px;
  display: flex;
  align-items: center;
  padding: 0 16px;
  font-weight: 600;
}
.t-layout__header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.t-layout__content {
  padding: 16px;
}
</style>
