<template>
  <div class="panel">
    <h2>需求浏览与响应</h2>
    <el-table :data="needs" size="small">
      <el-table-column label="需求" min-width="180">
        <template #default="{ row }">
          <strong>{{ row.title }}</strong>
          <div class="muted">{{ row.requester }} · 期望 {{ row.expectTime }}</div>
        </template>
      </el-table-column>
      <el-table-column prop="category" label="类别" width="80" />
      <el-table-column prop="campus" label="校区" width="90" />
      <el-table-column label="状态" width="90">
        <template #default="{ row }">
          <el-tag :type="NEED_STATUS_TAG[row.status as NeedStatus]" effect="light">{{ row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="210">
        <template #default="{ row }">
          <el-button v-if="canRespond(row)" size="small" type="primary" @click="openRespond(row)">
            发起响应
          </el-button>
          <el-button v-else-if="hasResponded(row)" size="small" disabled>已响应</el-button>
          <el-button size="small" @click="openResponses(row)">
            {{ row.requester === store.currentUser ? `管理响应(${pendingCount(row)})` : '查看响应' }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <ResponseFormDialog v-model="respondVisible" :need="activeNeed" />
    <NeedResponsesDrawer v-model="drawerVisible" :need="activeNeed" />
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import { NEED_STATUS_TAG } from '../constants/swap.constants';
import { useDashboardStore } from '../stores/dashboard.store';
import type { Need, NeedStatus } from '../types/domain';
import NeedResponsesDrawer from './NeedResponsesDrawer.vue';
import ResponseFormDialog from './ResponseFormDialog.vue';

const store = useDashboardStore();
const respondVisible = ref(false);
const drawerVisible = ref(false);
const activeNeed = ref<Need | null>(null);

const needs = computed(() => store.overview?.needs ?? []);

function hasResponded(need: Need): boolean {
  return store.responsesFor(need.id).some((resp) => resp.responder === store.currentUser);
}

function canRespond(need: Need): boolean {
  return need.status === '开放中' && need.requester !== store.currentUser && !hasResponded(need);
}

function pendingCount(need: Need): number {
  return store.responsesFor(need.id).filter((resp) => resp.status === '等待中').length;
}

function openRespond(need: Need) {
  activeNeed.value = need;
  respondVisible.value = true;
}

function openResponses(need: Need) {
  activeNeed.value = need;
  drawerVisible.value = true;
}
</script>
