<template>
  <el-dialog
    :model-value="visible"
    title="发起响应"
    width="480px"
    @update:model-value="emit('update:visible', $event)"
    @closed="reset"
  >
    <p v-if="need" class="muted">响应「{{ need.requester }}」的需求：{{ need.title }}</p>
    <el-form label-width="90px">
      <el-form-item label="响应者">
        <el-input :model-value="currentUser" disabled />
      </el-form-item>
      <el-form-item label="提供技能" required>
        <el-input v-model="form.offerSkill" placeholder="如：毕业照人像摄影" />
      </el-form-item>
      <el-form-item label="空闲时段" required>
        <el-select v-model="form.timeSlot" placeholder="选择空闲时段" style="width: 100%">
          <el-option v-for="slot in TIME_SLOTS" :key="slot" :label="slot" :value="slot" />
        </el-select>
      </el-form-item>
      <el-form-item label="地点类型" required>
        <el-radio-group v-model="form.placeType">
          <el-radio-button v-for="type in PLACE_TYPES" :key="type" :value="type">{{ type }}</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="地点" required>
        <el-input v-model="form.place" :placeholder="placePlaceholder" />
      </el-form-item>
      <el-form-item label="备注">
        <el-input v-model="form.note" type="textarea" :rows="2" placeholder="补充说明（可选）" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="emit('update:visible', false)">取消</el-button>
      <el-button type="primary" :loading="submitting" @click="submit">提交响应</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { computed, reactive, ref } from 'vue';
import { ElMessage } from 'element-plus';
import { PLACE_TYPES, TIME_SLOTS } from '../constants/response.constants';
import { createResponse } from '../services/response.service';
import type { Need } from '../types/domain';

const props = defineProps<{ visible: boolean; need: Need | null; currentUser: string }>();
const emit = defineEmits<{ 'update:visible': [value: boolean]; submitted: [] }>();

const submitting = ref(false);
const form = reactive({ offerSkill: '', timeSlot: '', placeType: PLACE_TYPES[1] as string, place: '', note: '' });

const placePlaceholder = computed(() =>
  form.placeType === PLACE_TYPES[0] ? '如：腾讯会议号 / 会议链接' : '如：东校区湖边',
);

function reset() {
  form.offerSkill = '';
  form.timeSlot = '';
  form.placeType = PLACE_TYPES[1];
  form.place = '';
  form.note = '';
}

async function submit() {
  if (!props.need) return;
  if (!form.offerSkill.trim() || !form.timeSlot || !form.place.trim()) {
    ElMessage.warning('请填写提供技能、空闲时段和地点');
    return;
  }
  submitting.value = true;
  try {
    await createResponse(props.need.id, {
      responder: props.currentUser,
      offerSkill: form.offerSkill.trim(),
      timeSlot: form.timeSlot,
      placeType: form.placeType,
      place: form.place.trim(),
      note: form.note.trim(),
    });
    ElMessage.success('响应已提交，等待发布者处理');
    emit('submitted');
    emit('update:visible', false);
  } catch (err) {
    ElMessage.error(err instanceof Error ? err.message : '提交失败');
  } finally {
    submitting.value = false;
  }
}
</script>
