<template>
  <t-card title="任期分配" hover-shadow>
    <template #actions>
      <t-button theme="primary">新增任期</t-button>
    </template>
    <t-form layout="inline" label-width="0" style="margin-bottom: 12px">
      <t-form-item><t-input placeholder="成员" /></t-form-item>
      <t-form-item><t-input placeholder="部门" /></t-form-item>
      <t-form-item><t-input placeholder="职务" /></t-form-item>
      <t-form-item><t-select :options="[{label:'在任',value:'active'},{label:'历史',value:'inactive'}]" placeholder="状态" /></t-form-item>
    </t-form>
    <t-table :data="tableData" :columns="columns" row-key="id" bordered stripe />
  </t-card>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const tableData = ref([
  { id: 1, member: '张三', department: '宣传部', position: '部长', start: '2024-09-01', end: '' },
]);

const columns = [
  { colKey: 'member', title: '成员' },
  { colKey: 'department', title: '部门/队伍' },
  { colKey: 'position', title: '职务' },
  { colKey: 'start', title: '开始日期' },
  { colKey: 'end', title: '结束日期', cell: ({ row }: any) => row.end || '在任' },
  { colKey: 'op', title: '操作', cell: () => '—' },
];
</script>
