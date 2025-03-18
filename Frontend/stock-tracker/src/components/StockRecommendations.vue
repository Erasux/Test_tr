<script setup lang="ts">
import { ref, onMounted } from "vue";
import { fetchStockRecommendations } from "../services/stockService";
import type { Stock } from "../types/Stock";

interface Recommendation {
  stock: Stock;
  score: number;
  recommendation: string;
}

const recommendations = ref<Recommendation[]>([]);
const loading = ref(false);
const error = ref<string | null>(null);

const fetchRecommendations = async () => {
  loading.value = true;
  error.value = null;

  try {
    recommendations.value = await fetchStockRecommendations();
  } catch (err) {
    error.value = "Error fetching recommendations";
  } finally {
    loading.value = false;
  }
};

onMounted(fetchRecommendations);
</script>

<template>
  <div class="container mt-4">
    <h2 class="text-center">📈 Investment Recommendations</h2>

    <button @click="fetchRecommendations" class="btn btn-primary my-3">
      Refresh Recommendations
    </button>

    <div v-if="loading" class="text-center">Loading recommendations...</div>
    <div v-else-if="error" class="text-danger">{{ error }}</div>

    <table v-else class="table table-striped">
      <thead>
        <tr>
          <th>Company</th>
          <th>Ticker</th>
          <th>Target Price</th>
          <th>Brokerage</th>
          <th>Rating</th>
          <th>Score</th>
          <th>Recommendation</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="rec in recommendations" :key="rec.stock.ticker">
          <td>{{ rec.stock.company }}</td>
          <td>{{ rec.stock.ticker }}</td>
          <td>${{ rec.stock.target_from }} → ${{ rec.stock.target_to }}</td>
          <td>{{ rec.stock.brokerage }}</td>
          <td>{{ rec.stock.rating_from }} → {{ rec.stock.rating_to }}</td>
          <td>{{ rec.score.toFixed(1) }}</td>
          <td>{{ rec.recommendation }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
