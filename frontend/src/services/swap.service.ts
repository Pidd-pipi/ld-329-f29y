import { AppException } from '../errors/AppException';
import type { ExchangeMode, ExchangeOrder, NeedResponse } from '../types/domain';

const API_BASE = '/api';

async function parseError(response: Response): Promise<never> {
  let message = '请求失败，请稍后重试';
  try {
    const body = (await response.json()) as { message?: string };
    if (body.message) {
      message = body.message;
    }
  } catch {
    // 保留默认提示
  }
  throw new AppException(message);
}

async function request<T>(url: string, init?: RequestInit): Promise<T> {
  const response = await fetch(url, init);
  if (!response.ok) {
    return parseError(response);
  }
  return response.json() as Promise<T>;
}

export interface CreateResponsePayload {
  responder: string;
  offerSkill: string;
  timeSlots: string[];
  mode: ExchangeMode;
  place: string;
  note: string;
}

export function createResponse(needId: number, payload: CreateResponsePayload): Promise<NeedResponse> {
  return request<NeedResponse>(`${API_BASE}/needs/${needId}/responses`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  });
}

export function selectResponse(needId: number, responseId: number, timeSlot: string): Promise<ExchangeOrder> {
  return request<ExchangeOrder>(`${API_BASE}/needs/${needId}/select`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ responseId, timeSlot }),
  });
}

export function confirmOrder(orderId: number, user: string): Promise<ExchangeOrder> {
  return request<ExchangeOrder>(`${API_BASE}/orders/${orderId}/confirm`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ user }),
  });
}
