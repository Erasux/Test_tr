<template>
    <div class="overflow-x-auto bg-white shadow-md rounded-lg p-4">
        <table class="min-w-full border border-gray-300 rounded-lg">
            <thead class="bg-gray-200 text-gray-700 uppercase text-sm">
                <tr>
                    <th class="px-6 py-3 text-left">Ticker</th>
                    <th class="px-6 py-3 text-left">Company</th>
                    <th class="px-6 py-3 text-left">Brokerage</th>
                    <th class="px-6 py-3 text-left">Action</th>
                    <th class="px-6 py-3 text-left">Rating From</th>
                    <th class="px-6 py-3 text-left">Rating To</th>
                    <th class="px-6 py-3 text-left">Target From</th>
                    <th class="px-6 py-3 text-left">Target To</th>
                    <th class="px-6 py-3 text-left">Score</th>
                    <th class="px-6 py-3 text-left">Recommendation</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="stock in recommendations" :key="stock.ticker" class="border-t hover:bg-gray-100">
                    <td class="px-6 py-3">{{ stock.ticker }}</td>
                    <td class="px-6 py-3">{{ stock.company }}</td>
                    <td class="px-6 py-3">{{ stock.brokerage }}</td>
                    <td class="px-6 py-3">{{ stock.action }}</td>
                    <td class="px-6 py-3">{{ stock.rating_from }}</td>
                    <td class="px-6 py-3">{{ stock.rating_to }}</td>
                    <td class="px-6 py-3">${{ stock.target_from.toFixed(2) }}</td>
                    <td class="px-6 py-3">${{ stock.target_to.toFixed(2) }}</td>
                    <td class="px-6 py-3 font-bold text-blue-600">{{ stock.score }}</td>
                    <td class="px-6 py-3 font-bold" :class="getRecommendationClass(stock.recommendation)">
                        {{ stock.recommendation }}
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from "vue";
import { useStockRecommendationStore } from "../../stores/stockRecommendationStore";

const store = useStockRecommendationStore();
const recommendations = computed(() => store.recommendations);

onMounted(() => {
    store.loadRecommendations();
});

const getRecommendationClass = (recommendation: string) => {
    return {
        "text-green-600": recommendation === "Strong Buy",
        "text-yellow-500": recommendation === "Buy",
        "text-gray-500": recommendation === "Hold",
        "text-red-600": recommendation === "Sell",
    };
};
</script>