<template>
  <el-table :data="needs" size="small">
    <el-table-column label="需求" min-width="180">
      <template #default="{ row }">
        <strong>{{ row.title }}</strong>
        <div class="muted">{{ row.requester }} · 期望 {{ row.expectTime }}</div>
      </template>
    </el-table-column>
    <el-table-column prop="category" label="类别" width="76" />
    <el-table-column prop="campus" label="校区" width="90" />
    <el-table-column prop="responses" label="响应" width="66" sortable />
    <el-table-column label="状态" width="92">
      <template #default="{ row }">
        <el-tag :type="statusTag(row.status)" effect="light">{{ row.status }}</el-tag>
      </template>
    </el-table-column>
    <el-table-column label="操作" width="196">
      <template #default="{ row }">
        <el-button size="small" type="primary" :disabled="!canRespond(row)" @click="emit('respond', row)">
          发起响应
        </el-button>
        <el-button size="small" :disabled="row.responses === 0" @click="emit('manage', row)">
          响应处理
        </el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script setup lang="ts">
import { NEED_STATUS_OPEN, NEED_STATUS_TAG } from '../constants/response.constants';
import type { Need } from '../types/domain';

const props = defineProps<{ needs: Need[]; currentUser: string }>();
const emit = defineEmits<{ respond: [need: Need]; manage: [need: Need] }>();

const canRespond = (need: Need) => need.status === NEED_STATUS_OPEN && need.requester !== props.currentUser;
const statusTag = (status: string) => NEED_STATUS_TAG[status] ?? 'info';
</script>
