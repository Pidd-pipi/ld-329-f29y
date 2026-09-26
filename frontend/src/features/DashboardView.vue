<template>
  <main class="page-shell" v-loading="loading">
    <AppHeader :unread="overview?.metrics.unread ?? 0" />

    <section v-if="overview" class="metrics-grid">
      <MetricCard label="已发布技能" :value="overview.metrics.skills" />
      <MetricCard label="活跃需求" :value="overview.metrics.needs" />
      <MetricCard label="需求响应" :value="overview.metrics.responses" />
      <MetricCard label="交换单" :value="overview.metrics.orders" />
    </section>

    <el-alert v-if="error" :title="error" type="error" show-icon />

    <section v-if="overview" class="workspace-grid">
      <div class="panel">
        <h2>技能发布</h2>
        <FeatureCard v-for="skill in overview.skills" :key="skill.id" :title="skill.title" :description="skill.description">
          <template #tag><el-tag>{{ skill.category }} {{ skill.level }}%</el-tag></template>
          <div class="tag-row">
            <el-tag v-for="slot in skill.timeSlots" :key="slot" effect="plain">{{ slot }}</el-tag>
            <el-tag v-for="reward in skill.rewards" :key="reward" type="success" effect="plain">{{ reward }}</el-tag>
          </div>
          <small>{{ skill.owner }} · {{ skill.campus }} · {{ skill.portfolio }}</small>
        </FeatureCard>
      </div>

      <NeedBoard />

      <ExchangeOrderPanel />

      <div class="panel">
        <h2>智能匹配</h2>
        <FeatureCard v-for="match in overview.matches" :key="match.id" :title="`${match.provider} × ${match.learner}`" :description="match.recommendation">
          <template #tag><el-tag type="warning">{{ match.score }}%</el-tag></template>
          <p class="muted">{{ match.offerSkill }} ↔ {{ match.wantedSkill }}</p>
          <div class="tag-row">
            <el-tag v-for="slot in match.commonSlots" :key="slot">{{ slot }}</el-tag>
          </div>
        </FeatureCard>
      </div>

      <div class="panel">
        <h2>已确认预约</h2>
        <el-empty v-if="overview.appointments.length === 0" description="还没有生效的预约" />
        <el-timeline v-else>
          <el-timeline-item v-for="item in overview.appointments" :key="item.id" :timestamp="item.time">
            <strong>{{ item.pair }}</strong>
            <p>{{ item.mode }} · {{ item.place }} · {{ item.status }}</p>
            <p class="muted">{{ item.agenda }}</p>
          </el-timeline-item>
        </el-timeline>
      </div>

      <div class="panel profile-panel">
        <div>
          <h2>个人主页与技能墙</h2>
          <h3>{{ overview.profile.name }}</h3>
          <p>{{ overview.profile.major }} · {{ overview.profile.creditLevel }}</p>
          <el-progress :percentage="overview.profile.creditScore" />
          <ul>
            <li v-for="item in overview.profile.history" :key="item">{{ item }}</li>
          </ul>
          <MySwapStatus />
        </div>
        <RadarChart :radar="overview.profile.radar" />
      </div>

      <div class="panel">
        <h2>评价信用</h2>
        <FeatureCard v-for="review in overview.reviews" :key="review.id" :title="`${review.from} → ${review.to}`" :description="review.content">
          <template #tag><el-rate :model-value="review.rating" disabled size="small" /></template>
        </FeatureCard>
      </div>

      <div class="panel">
        <h2>消息通知</h2>
        <FeatureCard v-for="conversation in overview.messages" :key="conversation.id" :title="conversation.withUser" :description="conversation.messages.join(' / ')">
          <template #tag><el-badge :value="conversation.unread" /></template>
        </FeatureCard>
      </div>
    </section>
  </main>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { storeToRefs } from 'pinia';
import AppHeader from '../components/AppHeader.vue';
import ExchangeOrderPanel from '../components/ExchangeOrderPanel.vue';
import FeatureCard from '../components/FeatureCard.vue';
import MetricCard from '../components/MetricCard.vue';
import MySwapStatus from '../components/MySwapStatus.vue';
import NeedBoard from '../components/NeedBoard.vue';
import RadarChart from '../components/RadarChart.vue';
import { useDashboardStore } from '../stores/dashboard.store';

const store = useDashboardStore();
const { overview, loading, error } = storeToRefs(store);

onMounted(store.load);
</script>
