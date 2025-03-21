<template>
    <div class="overflow-x-auto bg-white shadow-md rounded-lg p-4">
        <Loader v-if="loading" />
        <p v-if="errorMessage" class="text-red-600 text-center font-bold">{{ errorMessage }}</p>

        <table v-if="!loading && !errorMessage && stocks.length > 0"
            class="min-w-full border border-gray-300 rounded-lg">
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
                </tr>
            </thead>
            <tbody>
                <tr v-for="stock in stocks" :key="stock.id" class="border-t hover:bg-gray-100">
                    <td class="px-6 py-3">{{ stock.ticker }}</td>
                    <td class="px-6 py-3">{{ stock.company }}</td>
                    <td class="px-6 py-3">{{ stock.brokerage }}</td>
                    <td class="px-6 py-3">{{ stock.action }}</td>
                    <td class="px-6 py-3">{{ stock.rating_from || "N/A" }}</td>
                    <td class="px-6 py-3">{{ stock.rating_to || "N/A" }}</td>
                    <td class="px-6 py-3">${{ stock.target_from ? stock.target_from.toFixed(2) : "N/A" }}</td>
                    <td class="px-6 py-3">${{ stock.target_to ? stock.target_to.toFixed(2) : "N/A" }}</td>
                </tr>
            </tbody>
        </table>

        <p v-if="!loading && !errorMessage && stocks.length === 0" class="text-center text-gray-500">
            No stocks available.
        </p>
    </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from "vue";
import { useStockStore } from "@/stores/stockStore";
import Loader from "@/components/Loader.vue";

const stockStore = useStockStore();
const stocks = computed(() => stockStore.filteredStocks);
const loading = computed(() => stockStore.loading);
const errorMessage = computed(() => stockStore.errorMessage);

onMounted(() => {
    stockStore.loadStocks(); // No duplica la petición si los datos ya están cargados
});
</script>
