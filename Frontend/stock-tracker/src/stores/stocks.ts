import { defineStore } from "pinia";
import type { Stock } from "../types/Stock";
import { fetchStocks } from "../services/stockService";

export const useStockStore = defineStore("stocks", {
  state: () => ({
    stocks: [] as Stock[],
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
        this.error = "Error fetching stocks";
      } finally {
        this.loading = false;
      }
    },
  },
});
