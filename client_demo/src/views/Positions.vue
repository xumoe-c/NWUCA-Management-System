<template>
  <t-card title="职务管理" hover-shadow>
    <template #actions>
      <t-button theme="primary" @click="visible=true">新增职务</t-button>
    </template>
    <t-table :data="tableData" :columns="columns" row-key="id" bordered stripe />
  </t-card>

  <t-dialog v-model:visible="visible" header="新增职务" :confirm-on-enter="true" @confirm="onConfirm">
    <t-form :data="form" :rules="rules">
      <t-form-item label="名称" name="name"><t-input v-model="form.name" /></t-form-item>
      <t-form-item label="访问级别" name="access_level"><t-input-number v-model="form.access_level" /></t-form-item>
    </t-form>
  </t-dialog>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
import { MessagePlugin, PrimaryTableCol } from 'tdesign-vue-next';

const visible = ref(false);
const form = reactive({ name: '', access_level: 1 });
const rules = {
  name: [{ required: true, message: '请输入名称' }],
  access_level: [{ required: true, message: '请输入访问级别' }],
};

const tableData = ref([
  { id: 1, name: '部长', access_level: 100 },
  { id: 2, name: '干事', access_level: 10 },
]);

const columns: PrimaryTableCol[] = [
  { colKey: 'name', title: '名称' },
  { colKey: 'access_level', title: '访问级别' },
  { colKey: 'op', title: '操作', cell: () => '—' },
];

function onConfirm() {
  // TODO: 接入后端
  tableData.value.push({ id: Date.now(), name: form.name, access_level: form.access_level });
  visible.value = false;
  MessagePlugin.success('新增成功');
}
</script>
