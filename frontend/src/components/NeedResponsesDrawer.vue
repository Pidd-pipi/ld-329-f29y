<template>
  <el-drawer
    :model-value="modelValue"
    :title="`响应列表：${need?.title ?? ''}`"
    size="420px"
    @update:model-value="emit('update:modelValue', $event)"
  >
    <el-empty v-if="need && responses.length === 0" description="还没有响应，去邀请同学吧" />
    <article v-for="resp in responses" :key="resp.id" class="feature-card">
      <div class="feature-card__title">
        <strong>{{ resp.responder }} · {{ resp.offerSkill }}</strong>
        <el-tag :type="RESPONSE_STATUS_TAG[resp.status]" effect="light">{{ resp.status }}</el-tag>
      </div>
      <div class="tag-row">
        <el-tag v-for="slot in resp.timeSlots" :key="slot" effect="plain">{{ slot }}</el-tag>
      </div>
      <p class="muted">{{ resp.mode }} · {{ resp.place }} · {{ resp.createdAt }}</p>
      <p v-if="resp.note">{{ resp.note }}</p>
      <div v-if="canSelect(resp)" class="card-actions">
        <el-select
          v-model="pickedSlots[resp.id]"
          placeholder="选择交换时段"
          size="small"
          style="width: 150px"
        >
          <el-option v-for="slot in resp.timeSlots" :key="slot" :label="slot" :value="slot" />
        </el-select>
        <el-button
          size="small"
          type="primary"
          :disabled="!pickedSlots[resp.id]"
          :loading="selectingId === resp.id"
          @click="select(resp)"
        >
          选为交换对象
        </el-button>
      </div>
    </article>
  </el-drawer>
</template>

<script setup lang="ts">
import { computed, reactive, ref } from 'vue';
import { ElMessage, ElMessageBox } from 'element-plus';
import { RESPONSE_STATUS_TAG } from '../constants/swap.constants';
import { useDashboardStore } from '../stores/dashboard.store';
import type { Need, NeedResponse } from '../types/domain';

const props = defineProps<{ modelValue: boolean; need: Need | null }>();
const emit = defineEmits<{ 'update:modelValue': [value: boolean] }>();

const store = useDashboardStore();
const pickedSlots = reactive<Record<number, string>>({});
const selectingId = ref(0);

const responses = computed(() => (props.need ? store.responsesFor(props.need.id) : []));

function canSelect(resp: NeedResponse): boolean {
  return (
    !!props.need &&
    props.need.requester === store.currentUser &&
    props.need.status === '开放中' &&
    resp.status === '等待中'
  );
}

async function select(resp: NeedResponse) {
  if (!props.need) return;
  const slot = pickedSlots[resp.id];
  try {
    await ElMessageBox.confirm(
      `确定选择 ${resp.responder}，在「${slot}」以「${resp.offerSkill}」进行交换吗？其余响应将标记为未选中。`,
      '生成交换单',
      { confirmButtonText: '生成', cancelButtonText: '再想想', type: 'warning' },
    );
  } catch {
    return;
  }
  selectingId.value = resp.id;
  try {
    await store.selectResponder(props.need.id, resp.id, slot);
    ElMessage.success('交换单已生成，等待双方确认');
  } catch (err) {
    ElMessage.error(err instanceof Error ? err.message : '操作失败，请稍后重试');
  } finally {
    selectingId.value = 0;
  }
}
</script>
