<template>
    <div class="container mx-auto p-6 bg-gray-50 min-h-screen">
        <h1 class="text-3xl font-bold mb-6 text-center text-gray-800">📈 Stock Recommendations</h1>
        <StockRecommendations />
    </div>
</template>

<script setup lang="ts">
import StockRecommendations from "../components/stocks/StockRecommendations.vue";
</script>