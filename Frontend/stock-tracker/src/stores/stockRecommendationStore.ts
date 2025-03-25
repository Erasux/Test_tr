import { defineStore } from "pinia";
import { ref } from "vue";
import type { StockRecommendation, StockRecommendationResponse } from "@/types/StockRecommendation";

export const useStockRecommendationStore = defineStore("stockRecommendation", () => {
  const recommendations = ref<StockRecommendation[]>([]);
  const errorMessage = ref<string | null>(null);
  const loading = ref(true);

  async function loadRecommendations() {
    try {
      loading.value = true; // Iniciamos el estado de carga
      errorMessage.value = null; // Reseteamos errores previos

      const response = await fetch("http://localhost:9090/stocks/recommendations");

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const data = await response.json() as StockRecommendationResponse;

      recommendations.value = data.data.map((item) => ({
        ...item.stock,
        score: item.score.toFixed(2),
        recommendation: item.recommendation
      }));
    } catch (error) {
      console.error("Error fetching stock recommendations:", error);
      errorMessage.value = "Error loading stock recommendations. Please try again later.";
    } finally {
      loading.value = false; // Aseguramos que el loading se detenga siempre
    }
  }

  return { recommendations, errorMessage, loading, loadRecommendations };
});
