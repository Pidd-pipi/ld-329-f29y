import type { NeedResponse, SwapOrder } from '../types/domain';
import { request } from './api';

export interface CreateResponsePayload {
  responder: string;
  offerSkill: string;
  timeSlot: string;
  placeType: string;
  place: string;
  note: string;
}

export async function fetchResponses(needId: number): Promise<NeedResponse[]> {
  return request<NeedResponse[]>(`/needs/${needId}/responses`);
}

export async function createResponse(needId: number, payload: CreateResponsePayload): Promise<NeedResponse> {
  return request<NeedResponse>(`/needs/${needId}/responses`, {
    method: 'POST',
    body: JSON.stringify(payload),
  });
}

export async function acceptResponse(needId: number, responseId: number): Promise<SwapOrder> {
  return request<SwapOrder>(`/needs/${needId}/responses/${responseId}/accept`, { method: 'POST' });
}
