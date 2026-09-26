<template>
  <div class="swap-orders">
    <el-empty v-if="orders.length === 0" description="暂无交换单" :image-size="60" />
    <el-timeline v-else>
      <el-timeline-item v-for="order in orders" :key="order.id" :timestamp="order.timeSlot" placement="top">
        <div class="swap-orders__head">
          <strong>{{ order.requester }} ↔ {{ order.responder }}</strong>
          <el-tag :type="statusTag(order.status)" effect="light">{{ order.status }}</el-tag>
        </div>
        <p>{{ order.needTitle }} · {{ order.offerSkill }}</p>
        <p class="muted">{{ order.place }}</p>
        <div class="tag-row">
          <el-tag :type="order.requesterConfirmed ? 'success' : 'info'" effect="plain" size="small">
            发布者{{ order.requesterConfirmed ? '已确认' : '待确认' }}
          </el-tag>
          <el-tag :type="order.responderConfirmed ? 'success' : 'info'" effect="plain" size="small">
            响应者{{ order.responderConfirmed ? '已确认' : '待确认' }}
          </el-tag>
          <el-button v-if="canConfirm(order)" size="small" type="primary" @click="emit('confirm', order)">
            确认交换
          </el-button>
        </div>
      </el-timeline-item>
    </el-timeline>
  </div>
</template>

<script setup lang="ts">
import { ORDER_STATUS_PENDING, ORDER_STATUS_TAG } from '../constants/response.constants';
import type { SwapOrder } from '../types/domain';

const props = defineProps<{ orders: SwapOrder[]; currentUser: string }>();
const emit = defineEmits<{ confirm: [order: SwapOrder] }>();

const statusTag = (status: string) => ORDER_STATUS_TAG[status] ?? 'info';

function canConfirm(order: SwapOrder): boolean {
  if (order.status !== ORDER_STATUS_PENDING) return false;
  if (props.currentUser === order.requester) return !order.requesterConfirmed;
  if (props.currentUser === order.responder) return !order.responderConfirmed;
  return false;
}
</script>
