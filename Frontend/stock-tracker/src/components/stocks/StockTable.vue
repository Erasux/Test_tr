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
                </tr>
            </thead>
            <tbody>
                <tr v-for="stock in stocks" :key="stock.id" class="border-t hover:bg-gray-100">
                    <td class="px-6 py-3">{{ stock.ticker }}</td>
                    <td class="px-6 py-3">{{ stock.company }}</td>
                    <td class="px-6 py-3">{{ stock.brokerage }}</td>
                    <td class="px-6 py-3">{{ stock.action }}</td>
                    <td class="px-6 py-3">{{ stock.ratingFrom || "N/A" }}</td>
                    <td class="px-6 py-3">{{ stock.ratingTo || "N/A" }}</td>
                    <td class="px-6 py-3">${{ isNaN(stock.targetFrom) ? "N/A" : stock.targetFrom.toFixed(2) }}</td>
                    <td class="px-6 py-3">${{ isNaN(stock.targetTo) ? "N/A" : stock.targetTo.toFixed(2) }}</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script setup lang="ts">
defineProps<{ stocks: any[] }>();
</script>