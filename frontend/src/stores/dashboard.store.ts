import { defineStore } from 'pinia';
import { DEFAULT_USER } from '../constants/swap.constants';
import { logger } from '../logger/logger';
import { fetchOverview } from '../services/storage.service';
import { confirmOrder, createResponse, selectResponse, type CreateResponsePayload } from '../services/swap.service';
import type { ExchangeOrder, NeedResponse, Overview } from '../types/domain';

export const useDashboardStore = defineStore('dashboard', {
  state: () => ({
    overview: null as Overview | null,
    currentUser: DEFAULT_USER as string,
    loading: false,
    error: '',
  }),
  getters: {
    myResponses(state): NeedResponse[] {
      return state.overview?.responses.filter((item) => item.responder === state.currentUser) ?? [];
    },
    myOrders(state): ExchangeOrder[] {
      return (
        state.overview?.orders.filter(
          (order) => order.requester === state.currentUser || order.responder === state.currentUser,
        ) ?? []
      );
    },
    responsesFor(state) {
      return (needId: number): NeedResponse[] =>
        state.overview?.responses.filter((item) => item.needId === needId) ?? [];
    },
  },
  actions: {
    async load() {
      this.loading = true;
      this.error = '';
      try {
        this.overview = await fetchOverview();
      } catch (err) {
        this.error = err instanceof Error ? err.message : '加载失败';
        logger.error('overview load failed', err);
      } finally {
        this.loading = false;
      }
    },
    async submitResponse(needId: number, payload: Omit<CreateResponsePayload, 'responder'>) {
      await createResponse(needId, { ...payload, responder: this.currentUser });
      logger.info('response submitted', needId);
      await this.load();
    },
    async selectResponder(needId: number, responseId: number, timeSlot: string) {
      await selectResponse(needId, responseId, timeSlot);
      logger.info('responder selected', needId, responseId);
      await this.load();
    },
    async confirm(orderId: number, asUser?: string): Promise<ExchangeOrder> {
      const order = await confirmOrder(orderId, asUser ?? this.currentUser);
      logger.info('order confirmed', orderId, order.status);
      await this.load();
      return order;
    },
  },
});
