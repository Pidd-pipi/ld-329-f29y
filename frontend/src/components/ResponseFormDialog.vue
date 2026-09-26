<template>
  <el-dialog
    :model-value="modelValue"
    :title="`响应需求：${need?.title ?? ''}`"
    width="480px"
    @update:model-value="emit('update:modelValue', $event)"
    @open="resetForm"
  >
    <el-form label-position="top">
      <el-form-item label="响应者">
        <el-input :model-value="store.currentUser" disabled />
      </el-form-item>
      <el-form-item label="我能提供的技能" required>
        <el-input v-model="form.offerSkill" placeholder="如：毕业照人像摄影" maxlength="40" />
      </el-form-item>
      <el-form-item label="我的空闲时段（可多选）" required>
        <el-checkbox-group v-model="form.timeSlots">
          <el-checkbox v-for="slot in TIME_SLOT_OPTIONS" :key="slot" :value="slot">{{ slot }}</el-checkbox>
        </el-checkbox-group>
      </el-form-item>
      <el-form-item label="交换方式" required>
        <el-radio-group v-model="form.mode">
          <el-radio v-for="mode in EXCHANGE_MODES" :key="mode" :value="mode">{{ mode }}</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="地点" required>
        <el-input v-model="form.place" :placeholder="placePlaceholder" maxlength="60" />
      </el-form-item>
      <el-form-item label="备注">
        <el-input v-model="form.note" type="textarea" :rows="2" placeholder="补充说明，如期望的交换回报" maxlength="120" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="emit('update:modelValue', false)">取消</el-button>
      <el-button type="primary" :loading="submitting" @click="submit">提交响应</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { computed, reactive, ref } from 'vue';
import { ElMessage } from 'element-plus';
import {
  EXCHANGE_MODES,
  OFFLINE_PLACE_PLACEHOLDER,
  ONLINE_PLACE_PLACEHOLDER,
  TIME_SLOT_OPTIONS,
} from '../constants/swap.constants';
import { useDashboardStore } from '../stores/dashboard.store';
import type { ExchangeMode, Need } from '../types/domain';

const props = defineProps<{ modelValue: boolean; need: Need | null }>();
const emit = defineEmits<{ 'update:modelValue': [value: boolean] }>();

const store = useDashboardStore();
const submitting = ref(false);
const form = reactive({
  offerSkill: '',
  timeSlots: [] as string[],
  mode: '线上' as ExchangeMode,
  place: '',
  note: '',
});

const placePlaceholder = computed(() =>
  form.mode === '线上' ? ONLINE_PLACE_PLACEHOLDER : OFFLINE_PLACE_PLACEHOLDER,
);

function resetForm() {
  form.offerSkill = '';
  form.timeSlots = [];
  form.mode = '线上';
  form.place = '';
  form.note = '';
}

async function submit() {
  if (!props.need) return;
  if (!form.offerSkill.trim() || form.timeSlots.length === 0 || !form.place.trim()) {
    ElMessage.warning('请完整填写技能、空闲时段和地点');
    return;
  }
  submitting.value = true;
  try {
    await store.submitResponse(props.need.id, { ...form });
    ElMessage.success('响应已提交，等待发布者选择');
    emit('update:modelValue', false);
  } catch (err) {
    ElMessage.error(err instanceof Error ? err.message : '提交失败，请稍后重试');
  } finally {
    submitting.value = false;
  }
}
</script>
