import { defineStore } from "pinia";
import { ref, computed } from "vue";
import { fetchStocks } from "../services/stockService";

export const useStockStore = defineStore("stock", () => {
  const stocks = ref<any[]>([]);
  const filters = ref({
    ticker: "",
    brokerage: "",
    rating: "",
  });
  const errorMessage = ref<string | null>(null);
  const loading = ref(false);
  const isLoaded = ref(false);

  const filteredStocks = computed(() => {
    return stocks.value.filter((stock) => {
      return (
        (!filters.value.ticker || stock.ticker.toLowerCase().includes(filters.value.ticker.toLowerCase())) &&
        (!filters.value.brokerage || stock.brokerage.toLowerCase().includes(filters.value.brokerage.toLowerCase())) &&
        (!filters.value.rating || stock.ratingTo.toLowerCase().includes(filters.value.rating.toLowerCase()))
      );
    });
  });
  
  async function loadStocks(forceReload = false) {
    if (isLoaded.value && !forceReload) return; // Evita múltiples llamadas innecesarias

    try {
      loading.value = true;
      errorMessage.value = null;

      const response = await fetch("http://localhost:9090/stocks");
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const data = await response.json();
      stocks.value = data.data;
      isLoaded.value = true; // Marcar como cargado
    } catch (error) {
      console.error("Error fetching stocks:", error);
      errorMessage.value = "Error loading stocks. Please try again later.";
    } finally {
      loading.value = false;
    }
  }

  function setFilters(newFilters: any) {
    filters.value = newFilters;
  }

  return { stocks, filteredStocks,errorMessage, loading, loadStocks, setFilters };
});
