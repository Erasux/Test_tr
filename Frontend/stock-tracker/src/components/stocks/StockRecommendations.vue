<template>
    <div class="overflow-x-auto bg-white shadow-md rounded-lg p-4">
        <Loader v-if="loading" />
        <p v-if="errorMessage" class="text-red-600 text-center font-bold">{{ errorMessage }}</p>

        <table v-if="!loading && !errorMessage" class="min-w-full border border-gray-300 rounded-lg">
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
                    <td class="px-6 py-3">{{ stock.rating_from || "-" }}</td>
                    <td class="px-6 py-3">{{ stock.rating_to || "-" }}</td>
                    <td class="px-6 py-3">${{ stock.target_from ? stock.target_from.toFixed(2) : "-" }}</td>
                    <td class="px-6 py-3">${{ stock.target_to ? stock.target_to.toFixed(2) : "-" }}</td>

                    <td class="px-6 py-3 text-blue-600 font-bold">
                        {{ isNaN(Number(stock.score)) ? "-" : Number(stock.score).toFixed(2) }}
                    </td>

                    <td class="px-6 py-3 font-bold" :class="{
                        'text-green-600': stock.recommendation === 'Strong Buy' || stock.recommendation === 'Buy',
                        'text-gray-500': stock.recommendation === 'Hold',
                        'text-red-600': stock.recommendation === 'Sell'
                    }">
                        {{ stock.recommendation }}
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from "vue";
import { useStockRecommendationStore } from "@/stores/stockRecommendationStore";
import Loader from "@/components/Loader.vue";

const store = useStockRecommendationStore();
const recommendations = computed(() => store.recommendations);
const loading = computed(() => store.loading);
const errorMessage = computed(() => store.errorMessage);

onMounted(() => {
    store.loadRecommendations();
});


</script>