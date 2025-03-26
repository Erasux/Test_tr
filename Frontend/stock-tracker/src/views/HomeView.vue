<template>
  <div class="container mx-auto p-6 bg-gray-50 min-h-screen">
    <h1 class="text-3xl font-bold mb-6 text-center text-gray-800">📊 Stock List</h1>
    <StockFilters @filter-changed="applyFilters" />
    <StockTable :stocks="filteredStocks" />
  </div>
</template>


<script setup lang="ts">
import { computed, onMounted } from "vue";
import { useStockStore } from "@/stores/stockStore";
import StockFilters from "@/components/stocks/StockFilters.vue";
import StockTable from "@/components/stocks/StockTable.vue";
import type { StockFilters as Filters } from "@/types/stock";

const stockStore = useStockStore();
const filteredStocks = computed(() => stockStore.filteredStocks);

const applyFilters = (filters: Filters) => {
  stockStore.setFilters(filters);
};

onMounted(() => {
  stockStore.loadStocks(); // Carga los datos al entrar
});
</script>
