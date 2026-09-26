type TagType = 'primary' | 'success' | 'info' | 'warning' | 'danger';

export const NEED_STATUS_OPEN = '等待响应';
export const NEED_STATUS_BOOKED = '已约成';

export const RESPONSE_STATUS_PENDING = '等待中';
export const RESPONSE_STATUS_CHOSEN = '已选中';
export const RESPONSE_STATUS_REJECTED = '未选中';

export const ORDER_STATUS_PENDING = '待双方确认';
export const ORDER_STATUS_CONFIRMED = '已确认';

export const TIME_SLOTS = ['周二晚', '周三晚', '周六上午', '周六下午', '周日全天'] as const;

export const PLACE_TYPES = ['线上', '线下'] as const;

export const NEED_STATUS_TAG: Record<string, TagType> = {
  [NEED_STATUS_OPEN]: 'warning',
  [NEED_STATUS_BOOKED]: 'success',
};

export const RESPONSE_STATUS_TAG: Record<string, TagType> = {
  [RESPONSE_STATUS_PENDING]: 'warning',
  [RESPONSE_STATUS_CHOSEN]: 'success',
  [RESPONSE_STATUS_REJECTED]: 'info',
};

export const ORDER_STATUS_TAG: Record<string, TagType> = {
  [ORDER_STATUS_PENDING]: 'warning',
  [ORDER_STATUS_CONFIRMED]: 'success',
};
