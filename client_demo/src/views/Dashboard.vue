<template>
  <t-space direction="vertical" size="large" style="width: 100%">
    <t-row :gutter="16">
      <t-col :span="3">
        <t-card hover-shadow>
          <div class="card-stat">
            <div class="label">成员</div>
            <div class="value">--</div>
          </div>
        </t-card>
      </t-col>
      <t-col :span="3">
        <t-card hover-shadow>
          <div class="card-stat">
            <div class="label">部门</div>
            <div class="value">--</div>
          </div>
        </t-card>
      </t-col>
      <t-col :span="3">
        <t-card hover-shadow>
          <div class="card-stat">
            <div class="label">职务</div>
            <div class="value">--</div>
          </div>
        </t-card>
      </t-col>
      <t-col :span="3">
        <t-card hover-shadow>
          <div class="card-stat">
            <div class="label">在任</div>
            <div class="value">--</div>
          </div>
        </t-card>
      </t-col>
    </t-row>
    <t-card title="最近动态" hover-shadow>
      <t-empty description="稍后接入数据" />
    </t-card>
  </t-space>
</template>

<script setup lang="ts"></script>

<style scoped>
.card-stat .label { color: var(--td-text-color-secondary); }
.card-stat .value { font-size: 24px; font-weight: 600; }
</style>
