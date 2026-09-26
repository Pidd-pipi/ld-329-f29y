<template>
  <div class="profile-panel">
    <div>
      <h2>个人主页与技能墙</h2>
      <h3>{{ profile.name }}</h3>
      <p>{{ profile.major }} · {{ profile.creditLevel }}</p>
      <el-progress :percentage="profile.creditScore" />
      <ul>
        <li v-for="item in profile.history" :key="item">{{ item }}</li>
      </ul>
      <template v-if="profile.orders.length">
        <h4>我的交换单</h4>
        <div v-for="order in profile.orders" :key="order.id" class="profile-line">
          <span>{{ order.requester }} ↔ {{ order.responder }} · {{ order.timeSlot }}</span>
          <el-tag :type="orderTag(order.status)" size="small" effect="light">{{ order.status }}</el-tag>
        </div>
      </template>
      <template v-if="profile.myResponses.length">
        <h4>我的响应</h4>
        <div v-for="resp in profile.myResponses" :key="resp.id" class="profile-line">
          <span>{{ resp.offerSkill }} · {{ resp.timeSlot }}</span>
          <el-tag :type="responseTag(resp.status)" size="small" effect="light">{{ resp.status }}</el-tag>
        </div>
      </template>
    </div>
    <RadarChart :radar="profile.radar" />
  </div>
</template>

<script setup lang="ts">
import { ORDER_STATUS_TAG, RESPONSE_STATUS_TAG } from '../constants/response.constants';
import type { Profile } from '../types/domain';
import RadarChart from './RadarChart.vue';

defineProps<{ profile: Profile }>();

const orderTag = (status: string) => ORDER_STATUS_TAG[status] ?? 'info';
const responseTag = (status: string) => RESPONSE_STATUS_TAG[status] ?? 'info';
</script>
