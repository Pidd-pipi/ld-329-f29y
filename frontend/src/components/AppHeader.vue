<template>
  <header class="app-header">
    <div>
      <p class="eyebrow">Skill Swap Campus</p>
      <h1>{{ APP_NAME }}</h1>
      <p>{{ APP_TAGLINE }}</p>
    </div>
    <div class="header-actions">
      <span class="muted">当前身份</span>
      <el-select v-model="store.currentUser" style="width: 130px">
        <el-option v-for="user in USER_OPTIONS" :key="user" :label="user" :value="user" />
      </el-select>
      <el-badge :value="unread" class="message-badge">
        <el-button type="primary">站内消息</el-button>
      </el-badge>
    </div>
  </header>
</template>

<script setup lang="ts">
import { APP_NAME, APP_TAGLINE } from '../constants/app.constants';
import { USER_OPTIONS } from '../constants/swap.constants';
import { useDashboardStore } from '../stores/dashboard.store';

defineProps<{ unread: number }>();

const store = useDashboardStore();
</script>
