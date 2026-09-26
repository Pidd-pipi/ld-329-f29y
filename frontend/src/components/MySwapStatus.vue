<template>
  <div class="my-swap">
    <h3>我的响应</h3>
    <el-empty v-if="store.myResponses.length === 0" description="暂无响应记录" :image-size="60" />
    <p v-for="resp in store.myResponses" :key="resp.id" class="status-line">
      <span>{{ needTitle(resp.needId) }} · {{ resp.offerSkill }}</span>
      <el-tag size="small" :type="RESPONSE_STATUS_TAG[resp.status]" effect="light">{{ resp.status }}</el-tag>
    </p>
    <h3>我的交换单</h3>
    <el-empty v-if="store.myOrders.length === 0" description="暂无交换单" :image-size="60" />
    <p v-for="order in store.myOrders" :key="order.id" class="status-line">
      <span>{{ order.needTitle }} · {{ order.timeSlot }}</span>
      <el-tag size="small" :type="ORDER_STATUS_TAG[order.status]" effect="light">{{ order.status }}</el-tag>
    </p>
  </div>
</template>

<script setup lang="ts">
import { ORDER_STATUS_TAG, RESPONSE_STATUS_TAG } from '../constants/swap.constants';
import { useDashboardStore } from '../stores/dashboard.store';

const store = useDashboardStore();

function needTitle(needId: number): string {
  return store.overview?.needs.find((need) => need.id === needId)?.title ?? '历史需求';
}
</script>
