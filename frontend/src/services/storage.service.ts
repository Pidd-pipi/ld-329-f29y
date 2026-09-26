import type { Overview } from '../types/domain';
import { request } from './api';

export async function fetchOverview(): Promise<Overview> {
  return request<Overview>('/dashboard/overview');
}
