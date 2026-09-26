import type { ExchangeMode, NeedStatus, OrderStatus, ResponseStatus } from '../types/domain';

export const TIME_SLOT_OPTIONS = ['周一晚', '周二晚', '周三晚', '周四晚', '周五晚', '周六上午', '周六下午', '周日全天'] as const;

export const EXCHANGE_MODES: readonly ExchangeMode[] = ['线上', '线下'];

export const USER_OPTIONS = ['林澈', '孟野', '周芮', '许安'] as const;

export const DEFAULT_USER = '林澈';

type TagType = 'success' | 'warning' | 'info' | 'primary' | 'danger';

export const NEED_STATUS_TAG: Record<NeedStatus, TagType> = {
  开放中: 'success',
  匹配中: 'warning',
  已约成: 'info',
};

export const RESPONSE_STATUS_TAG: Record<ResponseStatus, TagType> = {
  等待中: 'warning',
  已选中: 'success',
  未选中: 'info',
};

export const ORDER_STATUS_TAG: Record<OrderStatus, TagType> = {
  待确认: 'warning',
  已确认: 'success',
};

export const ONLINE_PLACE_PLACEHOLDER = '如：腾讯会议号 / 会议链接';
export const OFFLINE_PLACE_PLACEHOLDER = '如：东校区湖边';
