<template>
  <t-row :gutter="16">
    <t-col :span="3">
      <t-input v-model="treeFilter" placeholder="搜索部门" clearable />
      <t-tree :data="treeData" :filter="treeFilter" activable hover transition style="margin-top: 8px" @active="onActive" />
    </t-col>
    <t-col :span="9">
      <t-card title="部门详情" hover-shadow>
        <template #actions>
          <t-space v-if="isAdmin">
            <t-button theme="primary" @click="openForm('create')">新增部门</t-button>
            <t-button variant="outline" :disabled="!currentDept" @click="openForm('edit')">编辑</t-button>
            <t-popconfirm content="确认删除该部门？" @confirm="onDelete">
              <t-button theme="danger" variant="outline" :disabled="!currentDept">删除</t-button>
            </t-popconfirm>
          </t-space>
        </template>
        <div v-if="currentDept">{{ currentDept.name }}</div>
        <t-empty v-else description="请选择左侧部门" />
      </t-card>
    </t-col>
  </t-row>

  <t-drawer v-model:visible="formVisible" :header="formMode==='create' ? '新增部门' : '编辑部门'" :footer="false" size="40%">
    <t-form :data="form" :rules="rules" @submit.prevent="onSubmit">
      <t-form-item label="名称" name="name"><t-input v-model="form.name" /></t-form-item>
      <t-form-item>
        <t-space>
          <t-button type="submit" theme="primary" :loading="saving">保存</t-button>
          <t-button variant="outline" @click="formVisible=false">取消</t-button>
        </t-space>
      </t-form-item>
    </t-form>
  </t-drawer>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue';
import { MessagePlugin } from 'tdesign-vue-next';
import { getDepartments, createDepartment, deleteDepartment, type Department } from '../api/department';
import { useAuthStore } from '../store/auth';

const auth = useAuthStore();
const isAdmin = computed(() => auth.isAdmin);

const treeFilter = ref('');
const treeData = ref<any[]>([]);
const currentDept = ref<Department | null>(null);

async function load() {
  try {
    const list = await getDepartments();
    treeData.value = list.map((d) => ({ value: String(d.id), label: d.name }));
  } catch (e: any) {
    MessagePlugin.error(e?.response?.data?.error || '加载失败');
  }
}

function onActive({ node }: any) {
  const id = Number(node.value);
  const name = String(node.label);
  currentDept.value = { id, name } as Department;
}

const formVisible = ref(false);
const formMode = ref<'create' | 'edit'>('create');
const form = reactive({ name: '' });
const rules = { name: [{ required: true, message: '请输入名称' }] };
const saving = ref(false);

function openForm(mode: 'create' | 'edit') {
  formMode.value = mode;
  formVisible.value = true;
  form.name = mode === 'edit' && currentDept.value ? currentDept.value.name : '';
}

async function onSubmit() {
  saving.value = true;
  try {
    if (formMode.value === 'create') {
      await createDepartment({ name: form.name });
      MessagePlugin.success('创建成功');
    } else {
      // TODO: 后端提供 PUT /api/v1/departments/:id 时对接
      MessagePlugin.info('编辑暂未对接，请稍后');
    }
    formVisible.value = false;
    await load();
  } catch (e: any) {
    MessagePlugin.error(e?.response?.data?.error || '保存失败');
  } finally {
    saving.value = false;
  }
}

async function onDelete() {
  if (!currentDept.value) return;
  try {
    await deleteDepartment(currentDept.value.id);
    MessagePlugin.success('删除成功');
    currentDept.value = null;
    await load();
  } catch (e: any) {
    MessagePlugin.error(e?.response?.data?.error || '删除失败');
  }
}

onMounted(load);
</script>
