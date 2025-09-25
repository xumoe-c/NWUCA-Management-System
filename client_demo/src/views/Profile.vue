<template>
  <t-row :gutter="16">
    <t-col :span="4">
      <t-card title="个人信息" hover-shadow>
        <t-space direction="vertical" size="small">
          <t-avatar size="72px">U</t-avatar>
          <div>用户名：demo</div>
          <div>角色：admin</div>
          <t-button size="small" variant="outline">修改密码</t-button>
        </t-space>
      </t-card>
    </t-col>
  </t-row>
</template>

<script setup lang="ts"></script>
