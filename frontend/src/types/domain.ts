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
  status: string;
}

export interface NeedResponse {
  id: number;
  needId: number;
  responder: string;
  offerSkill: string;
  timeSlot: string;
  placeType: string;
  place: string;
  note: string;
  status: string;
}

export interface SwapOrder {
  id: number;
  needId: number;
  needTitle: string;
  responseId: number;
  requester: string;
  responder: string;
  offerSkill: string;
  timeSlot: string;
  place: string;
  status: string;
  requesterConfirmed: boolean;
  responderConfirmed: boolean;
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

export interface Appointment {
  id: number;
  pair: string;
  time: string;
  timeSlot: string;
  place: string;
  status: string;
  agenda: string;
  participants: string[];
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

export interface Profile {
  name: string;
  major: string;
  creditScore: number;
  creditLevel: string;
  skillWall: Skill[];
  radar: Record<string, number>;
  history: string[];
  reviews: Review[];
  orders: SwapOrder[];
  myResponses: NeedResponse[];
}

export interface Overview {
  service: string;
  categories: string[];
  metrics: Record<string, number>;
  skills: Skill[];
  needs: Need[];
  matches: Match[];
  appointments: Appointment[];
  swapOrders: SwapOrder[];
  reviews: Review[];
  messages: Conversation[];
  profile: Profile;
}
