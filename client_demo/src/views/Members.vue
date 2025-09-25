<template>
  <t-card title="成员管理" hover-shadow>
    <template #actions>
      <t-button theme="primary">新增成员</t-button>
    </template>
    <t-form layout="inline" label-width="0" style="margin-bottom: 12px">
      <t-form-item><t-input placeholder="关键词：姓名/学号/电话/邮箱" clearable /></t-form-item>
      <t-form-item><t-select placeholder="成员状态" :options="statusOptions" /></t-form-item>
      <t-form-item>
        <t-space>
          <t-button theme="primary" variant="outline">查询</t-button>
          <t-button variant="outline">重置</t-button>
        </t-space>
      </t-form-item>
    </t-form>
    <t-table :data="tableData" :columns="columns" row-key="id" bordered stripe />
  </t-card>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const statusOptions = [
  { label: '在读', value: 'studying' },
  { label: '已毕业', value: 'graduated' },
  { label: '休学', value: 'suspend' },
  { label: '退会', value: 'quit' },
];

const tableData = ref([
  { id: 1, name: '张三', student_id: '20210001', grade: '2021', college: '计算机学院', major: '软件工程', status: '在读' },
]);

const columns = [
  { colKey: 'name', title: '姓名' },
  { colKey: 'student_id', title: '学号' },
  { colKey: 'grade', title: '年级' },
  { colKey: 'college', title: '学院' },
  { colKey: 'major', title: '专业' },
  { colKey: 'status', title: '状态' },
  { colKey: 'op', title: '操作', cell: () => '—' },
];
</script>
