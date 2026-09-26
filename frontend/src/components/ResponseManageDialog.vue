<template>
  <el-dialog
    :model-value="visible"
    :title="need ? `响应处理 · ${need.title}` : '响应处理'"
    width="720px"
    @update:model-value="emit('update:visible', $event)"
  >
    <el-table v-loading="loading" :data="responses" size="small" empty-text="暂无响应">
      <el-table-column prop="responder" label="响应者" width="90" />
      <el-table-column prop="offerSkill" label="提供技能" min-width="130" />
      <el-table-column prop="timeSlot" label="时段" width="90" />
      <el-table-column label="地点" min-width="150">
        <template #default="{ row }">
          <el-tag size="small" effect="plain">{{ row.placeType }}</el-tag>
          {{ row.place }}
        </template>
      </el-table-column>
      <el-table-column label="状态" width="86">
        <template #default="{ row }">
          <el-tag :type="statusTag(row.status)" effect="light">{{ row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="96">
        <template #default="{ row }">
          <el-button
            v-if="canAccept && row.status === RESPONSE_STATUS_PENDING"
            size="small"
            type="primary"
            :loading="acceptingId === row.id"
            @click="accept(row)"
          >
            接受
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <p v-if="!canAccept && need" class="muted dialog-hint">
      {{ need.requester === currentUser ? '需求已约成，交换单等待双方确认。' : '仅发布者可以处理响应。' }}
    </p>
  </el-dialog>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { ElMessage } from 'element-plus';
import {
  NEED_STATUS_OPEN,
  RESPONSE_STATUS_PENDING,
  RESPONSE_STATUS_TAG,
} from '../constants/response.constants';
import { acceptResponse, fetchResponses } from '../services/response.service';
import type { Need, NeedResponse } from '../types/domain';

const props = defineProps<{ visible: boolean; need: Need | null; currentUser: string }>();
const emit = defineEmits<{ 'update:visible': [value: boolean]; changed: [] }>();

const responses = ref<NeedResponse[]>([]);
const loading = ref(false);
const acceptingId = ref<number | null>(null);

const canAccept = computed(
  () => !!props.need && props.need.requester === props.currentUser && props.need.status === NEED_STATUS_OPEN,
);

const statusTag = (status: string) => RESPONSE_STATUS_TAG[status] ?? 'info';

watch(
  () => [props.visible, props.need?.id, props.need?.status],
  () => {
    if (props.visible && props.need) {
      void load();
    }
  },
);

async function load() {
  if (!props.need) return;
  loading.value = true;
  try {
    responses.value = await fetchResponses(props.need.id);
  } catch (err) {
    ElMessage.error(err instanceof Error ? err.message : '加载响应失败');
  } finally {
    loading.value = false;
  }
}

async function accept(row: NeedResponse) {
  if (!props.need) return;
  acceptingId.value = row.id;
  try {
    await acceptResponse(props.need.id, row.id);
    ElMessage.success('已生成交换单，等待双方确认');
    emit('changed');
    await load();
  } catch (err) {
    ElMessage.error(err instanceof Error ? err.message : '操作失败');
  } finally {
    acceptingId.value = null;
  }
}
</script>
