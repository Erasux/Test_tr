<template>
    <div class="flex flex-wrap gap-4 mb-6 bg-white p-4 shadow-md rounded-lg">
        <input v-model="filters.ticker" placeholder="Search by Ticker"
            class="p-2 border border-gray-300 rounded shadow-sm focus:ring focus:ring-blue-300" />
        <input v-model="filters.brokerage" placeholder="Search by Brokerage"
            class="p-2 border border-gray-300 rounded shadow-sm focus:ring focus:ring-blue-300" />
        <input v-model="filters.rating" placeholder="Search by Rating"
            class="p-2 border border-gray-300 rounded shadow-sm focus:ring focus:ring-blue-300" />
        <button @click="applyFilters" class="px-4 py-2 bg-gray-500 text-white rounded hover:bg-gray-600 transition">
            Apply Filters
        </button>
    </div>
</template>


<script setup lang="ts">
import { ref } from "vue";

const filters = ref({
    ticker: "",
    brokerage: "",
    rating: "",
});

const emit = defineEmits(["filter-changed"]);

const applyFilters = () => {
    emit("filter-changed", filters.value);
};
</script>