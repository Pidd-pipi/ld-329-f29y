import type { SwapOrder } from '../types/domain';
import { request } from './api';

export async function fetchSwapOrders(): Promise<SwapOrder[]> {
  return request<SwapOrder[]>('/swap-orders');
}

export async function confirmSwapOrder(id: number, user: string): Promise<SwapOrder> {
  return request<SwapOrder>(`/swap-orders/${id}/confirm`, {
    method: 'POST',
    body: JSON.stringify({ user }),
  });
}
