import { defineStore } from "pinia";
import { ref } from "vue";

export const useStockRecommendationStore = defineStore("stockRecommendation", () => {
  const recommendations = ref<any[]>([]);

  async function loadRecommendations() {
    try {
      const response = await fetch("http://localhost:9090/stocks/recommendations");
      const data = await response.json();
      recommendations.value = data.data.map((item: any) => ({
        ...item.stock,
        score: item.score.toFixed(2),
        recommendation: item.recommendation
      }));
    } catch (error) {
      console.error("Error fetching stock recommendations:", error);
    }
  }

  return { recommendations, loadRecommendations };
});
