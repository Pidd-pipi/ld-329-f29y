export interface Skill {
  id: number;
  owner: string;
  title: string;
  category: string;
  level: number;
  campus: string;
  description: string;
  timeSlots: string[];
  rewards: string[];
  portfolio: string;
}

export type NeedStatus = '开放中' | '匹配中' | '已约成';

export interface Need {
  id: number;
  requester: string;
  title: string;
  category: string;
  campus: string;
  expectTime: string;
  budgetType: string;
  description: string;
  responses: number;
  status: NeedStatus;
}

export interface Match {
  id: number;
  provider: string;
  learner: string;
  offerSkill: string;
  wantedSkill: string;
  score: number;
  commonSlots: string[];
  recommendation: string;
}

export type ExchangeMode = '线上' | '线下';

export interface Appointment {
  id: number;
  pair: string;
  time: string;
  mode: ExchangeMode;
  place: string;
  status: string;
  agenda: string;
}

export interface Review {
  id: number;
  from: string;
  to: string;
  rating: number;
  content: string;
}

export interface Conversation {
  id: number;
  withUser: string;
  unread: number;
  messages: string[];
}

export type ResponseStatus = '等待中' | '已选中' | '未选中';

export interface NeedResponse {
  id: number;
  needId: number;
  responder: string;
  offerSkill: string;
  timeSlots: string[];
  mode: ExchangeMode;
  place: string;
  note: string;
  status: ResponseStatus;
  createdAt: string;
}

export type OrderStatus = '待确认' | '已确认';

export interface ExchangeOrder {
  id: number;
  needId: number;
  needTitle: string;
  responseId: number;
  requester: string;
  responder: string;
  offerSkill: string;
  timeSlot: string;
  mode: ExchangeMode;
  place: string;
  agenda: string;
  status: OrderStatus;
  confirmations: string[];
  createdAt: string;
}

export interface Profile {
  name: string;
  major: string;
  creditScore: number;
  creditLevel: string;
  skillWall: Skill[];
  radar: Record<string, number>;
  history: string[];
  reviews: Review[];
}

export interface Overview {
  service: string;
  categories: string[];
  metrics: Record<string, number>;
  skills: Skill[];
  needs: Need[];
  matches: Match[];
  appointments: Appointment[];
  reviews: Review[];
  messages: Conversation[];
  responses: NeedResponse[];
  orders: ExchangeOrder[];
  profile: Profile;
}
