import { defineStore } from "pinia";
import { fetchStocks } from "../services/stockService";

export const useStockStore = defineStore("stocks", {
  state: () => ({
    stocks: [] as any[],
    loading: false,
    error: null as string | null,
  }),
  actions: {
    async fetchStocks(filters: Record<string, string> = {}) {
      this.loading = true;
      this.error = null;

      try {
        this.stocks = await fetchStocks(filters);
      } catch (error) {
        this.error = error instanceof Error ? error.message : "Failed to fetch stocks";
      } finally {
        this.loading = false;
      }
    },
  },
});