<template>
  <div class="panel">
    <h2>交换单确认</h2>
    <el-empty v-if="sortedOrders.length === 0" description="还没有交换单，去响应需求或选择响应者吧" />
    <article v-for="order in sortedOrders" :key="order.id" class="feature-card">
      <div class="feature-card__title">
        <strong>{{ order.needTitle }}</strong>
        <el-tag :type="ORDER_STATUS_TAG[order.status]" effect="light">{{ order.status }}</el-tag>
      </div>
      <p>{{ order.requester }} ↔ {{ order.responder }} · {{ order.offerSkill }}</p>
      <p class="muted">{{ order.timeSlot }} · {{ order.mode }} · {{ order.place }}</p>
      <div class="tag-row">
        <el-tag
          v-for="participant in participants(order)"
          :key="participant.name"
          :type="participant.confirmed ? 'success' : 'info'"
          effect="plain"
        >
          {{ participant.name }}{{ participant.confirmed ? ' 已确认' : ' 待确认' }}
        </el-tag>
      </div>
      <div v-if="order.status === '待确认'" class="card-actions">
        <el-button
          v-if="canConfirmAs(order, store.currentUser)"
          size="small"
          type="primary"
          :loading="confirmingId === order.id"
          @click="confirm(order, store.currentUser)"
        >
          确认接受
        </el-button>
        <el-button
          v-for="name in otherUnconfirmed(order)"
          :key="name"
          size="small"
          text
          type="primary"
          @click="confirm(order, name)"
        >
          以 {{ name }} 身份确认
        </el-button>
      </div>
    </article>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import { ElMessage } from 'element-plus';
import { ORDER_STATUS_TAG } from '../constants/swap.constants';
import { useDashboardStore } from '../stores/dashboard.store';
import type { ExchangeOrder } from '../types/domain';

const store = useDashboardStore();
const confirmingId = ref(0);

const sortedOrders = computed(() =>
  [...(store.overview?.orders ?? [])].sort((a, b) => {
    if (a.status !== b.status) {
      return a.status === '待确认' ? -1 : 1;
    }
    return b.id - a.id;
  }),
);

function participants(order: ExchangeOrder) {
  return [order.requester, order.responder].map((name) => ({
    name,
    confirmed: order.confirmations.includes(name),
  }));
}

function canConfirmAs(order: ExchangeOrder, user: string): boolean {
  return order.status === '待确认' && !order.confirmations.includes(user) &&
    (order.requester === user || order.responder === user);
}

function otherUnconfirmed(order: ExchangeOrder): string[] {
  return [order.requester, order.responder].filter(
    (name) => name !== store.currentUser && !order.confirmations.includes(name),
  );
}

async function confirm(order: ExchangeOrder, user: string) {
  confirmingId.value = order.id;
  try {
    const updated = await store.confirm(order.id, user);
    ElMessage.success(
      updated.status === '已确认' ? '双方已确认，交换预约生效' : `${user} 已确认，等待对方确认`,
    );
  } catch (err) {
    ElMessage.error(err instanceof Error ? err.message : '确认失败，请稍后重试');
  } finally {
    confirmingId.value = 0;
  }
}
</script>
