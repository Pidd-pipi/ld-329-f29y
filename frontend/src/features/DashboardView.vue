<template>
  <main class="page-shell" v-loading="loading">
    <AppHeader :unread="overview?.metrics.unread ?? 0" />

    <section v-if="overview" class="metrics-grid">
      <MetricCard label="已发布技能" :value="overview.metrics.skills" />
      <MetricCard label="活跃需求" :value="overview.metrics.needs" />
      <MetricCard label="智能匹配" :value="overview.metrics.matches" />
      <MetricCard label="交换单" :value="overview.metrics.orders ?? 0" />
      <MetricCard label="评价记录" :value="overview.metrics.reviews" />
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

      <div class="panel">
        <h2>需求浏览</h2>
        <NeedBoard
          :needs="overview.needs"
          :current-user="currentUser"
          @respond="openRespond"
          @manage="openManage"
        />
      </div>

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
        <h2>交换预约</h2>
        <el-timeline>
          <el-timeline-item v-for="item in overview.appointments" :key="item.id" :timestamp="item.time">
            <strong>{{ item.pair }}</strong>
            <p>{{ item.place }} · {{ item.status }}</p>
            <p class="muted">{{ item.agenda }}</p>
          </el-timeline-item>
        </el-timeline>
        <el-divider content-position="left">交换单</el-divider>
        <SwapOrderPanel :orders="overview.swapOrders" :current-user="currentUser" @confirm="confirmOrder" />
      </div>

      <div class="panel">
        <ProfilePanel :profile="overview.profile" />
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

    <ResponseFormDialog
      v-model:visible="respondVisible"
      :need="respondNeed"
      :current-user="currentUser"
      @submitted="reload"
    />
    <ResponseManageDialog
      v-model:visible="manageVisible"
      :need="manageNeed"
      :current-user="currentUser"
      @changed="reload"
    />
  </main>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { ElMessage } from 'element-plus';
import AppHeader from '../components/AppHeader.vue';
import FeatureCard from '../components/FeatureCard.vue';
import MetricCard from '../components/MetricCard.vue';
import NeedBoard from '../components/NeedBoard.vue';
import ProfilePanel from '../components/ProfilePanel.vue';
import ResponseFormDialog from '../components/ResponseFormDialog.vue';
import ResponseManageDialog from '../components/ResponseManageDialog.vue';
import SwapOrderPanel from '../components/SwapOrderPanel.vue';
import { fetchOverview } from '../services/storage.service';
import { confirmSwapOrder } from '../services/swaporder.service';
import type { Need, Overview, SwapOrder } from '../types/domain';

const overview = ref<Overview | null>(null);
const loading = ref(true);
const error = ref('');

const respondVisible = ref(false);
const respondNeed = ref<Need | null>(null);
const manageVisible = ref(false);
const manageNeedId = ref<number | null>(null);

const currentUser = computed(() => overview.value?.profile.name ?? '');
const manageNeed = computed(
  () => overview.value?.needs.find((need) => need.id === manageNeedId.value) ?? null,
);

async function reload() {
  overview.value = await fetchOverview();
}

onMounted(async () => {
  try {
    await reload();
  } catch (err) {
    error.value = err instanceof Error ? err.message : '加载失败';
  } finally {
    loading.value = false;
  }
});

function openRespond(need: Need) {
  respondNeed.value = need;
  respondVisible.value = true;
}

function openManage(need: Need) {
  manageNeedId.value = need.id;
  manageVisible.value = true;
}

async function confirmOrder(order: SwapOrder) {
  try {
    await confirmSwapOrder(order.id, currentUser.value);
    ElMessage.success('已确认交换单');
    await reload();
  } catch (err) {
    ElMessage.error(err instanceof Error ? err.message : '确认失败');
  }
}
</script>
