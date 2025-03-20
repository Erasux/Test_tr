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

  const filteredStocks = computed(() => {
    return stocks.value.filter((stock) => {
      return (
        (!filters.value.ticker || stock.ticker.toLowerCase().includes(filters.value.ticker.toLowerCase())) &&
        (!filters.value.brokerage || stock.brokerage.toLowerCase().includes(filters.value.brokerage.toLowerCase())) &&
        (!filters.value.rating || stock.ratingTo.toLowerCase().includes(filters.value.rating.toLowerCase()))
      );
    });
  });
  
  async function loadStocks() {
    try {
      const response = await fetch("http://localhost:9090/stocks");
      const data = await response.json();
      stocks.value = data.data.map((stock: any) => ({
        ...stock,
        targetFrom: typeof stock.target_from === "number" ? stock.target_from : parseFloat(stock.target_from),
        targetTo: typeof stock.target_to === "number" ? stock.target_to : parseFloat(stock.target_to),
        ratingFrom: stock.rating_from || "N/A",  // Si el valor es null, mostrar "N/A"
        ratingTo: stock.rating_to || "N/A",
      }));
    } catch (error) {
      console.error("Error fetching stocks:", error);
    }
  }
  
  

  function setFilters(newFilters: any) {
    filters.value = newFilters;
  }

  return { stocks, filteredStocks, loadStocks, setFilters };
});
