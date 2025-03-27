<template>
  <div class="overflow-x-auto bg-white shadow-md rounded-lg p-4">
    <RecommendationExplanation />
    <StockFilters @filter-changed="applyFilters" />
    <Loader v-if="loading" />
    <ErrorMessage v-if="errorMessage" :message="errorMessage" />

    <div v-if="!loading && !errorMessage">
      <table class="min-w-full border border-gray-300 rounded-lg">
        <thead class="bg-gray-200 text-gray-700 uppercase text-sm">
          <tr>
            <th v-for="header in tableHeaders" :key="header.key"
              class="px-6 py-3 text-left cursor-pointer hover:bg-gray-300 transition-colors" @click="sortBy(header.key)"
              :aria-sort="sortColumn === header.key ? sortDirection : 'none'" role="columnheader">
              <div class="flex items-center gap-2">
                {{ header.label }}
                <span v-if="sortColumn === header.key" class="text-xs">
                  {{ sortDirection === 'ascending' ? '↑' : '↓' }}
                </span>
              </div>
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="stock in paginatedRecommendations" :key="stock.ticker" class="border-t hover:bg-gray-100">
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

      <TablePagination v-model:current-page="currentPage" :total-items="filteredRecommendations.length"
        :items-per-page="itemsPerPage" />
    </div>

    <EmptyState v-if="!loading && !errorMessage && recommendations.length === 0" />
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
import { useStockRecommendationStore } from "@/stores/stockRecommendationStore";
import Loader from "@/components/Loader.vue";
import ErrorMessage from "@/components/common/ErrorMessage.vue";
import EmptyState from "@/components/common/EmptyState.vue";
import TablePagination from "@/components/common/TablePagination.vue";
import StockFilters from "@/components/stocks/StockFilters.vue";
import RecommendationExplanation from "@/components/stocks/RecommendationExplanation.vue";
import type { StockFilters as Filters } from "@/types/stock";

interface TableHeader {
  key: string;
  label: string;
}

const tableHeaders: TableHeader[] = [
  { key: 'ticker', label: 'Ticker' },
  { key: 'company', label: 'Company' },
  { key: 'brokerage', label: 'Brokerage' },
  { key: 'action', label: 'Action' },
  { key: 'rating_from', label: 'Rating From' },
  { key: 'rating_to', label: 'Rating To' },
  { key: 'target_from', label: 'Target From' },
  { key: 'target_to', label: 'Target To' },
  { key: 'score', label: 'Score' },
  { key: 'recommendation', label: 'Recommendation' }
];

const store = useStockRecommendationStore();
const recommendations = computed(() => store.recommendations);
const loading = computed(() => store.loading);
const errorMessage = computed(() => store.errorMessage);

// Estado para ordenamiento
const sortColumn = ref<string | null>(null);
const sortDirection = ref<'ascending' | 'descending'>('ascending');

// Estado para paginación
const currentPage = ref(1);
const itemsPerPage = ref(15);

// Estado para filtros
const filters = ref<Filters>({
  ticker: "",
  brokerage: "",
  company: "",
});

// Aplicar filtros
const filteredRecommendations = computed(() => {
  let filtered = [...recommendations.value];

  // Aplicar filtros de búsqueda
  if (filters.value.ticker) {
    filtered = filtered.filter(stock =>
      stock.ticker.toLowerCase().includes(filters.value.ticker.toLowerCase())
    );
  }
  if (filters.value.brokerage) {
    filtered = filtered.filter(stock =>
      stock.brokerage.toLowerCase().includes(filters.value.brokerage.toLowerCase())
    );
  }
  if (filters.value.company) {
    filtered = filtered.filter(stock =>
      stock.company.toLowerCase().includes(filters.value.company.toLowerCase())
    );
  }

  // Aplicar ordenamiento
  if (sortColumn.value) {
    filtered.sort((a, b) => {
      const aValue = a[sortColumn.value as keyof typeof a];
      const bValue = b[sortColumn.value as keyof typeof b];

      if (aValue === null) return 1;
      if (bValue === null) return -1;

      const comparison = aValue < bValue ? -1 : aValue > bValue ? 1 : 0;
      return sortDirection.value === 'ascending' ? comparison : -comparison;
    });
  }

  return filtered;
});

// Recomendaciones paginadas
const paginatedRecommendations = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value;
  return filteredRecommendations.value.slice(start, start + itemsPerPage.value);
});

// Funciones de interacción
function sortBy(column: string) {
  if (sortColumn.value === column) {
    sortDirection.value = sortDirection.value === 'ascending' ? 'descending' : 'ascending';
  } else {
    sortColumn.value = column;
    sortDirection.value = 'ascending';
  }
  // Resetear a la primera página cuando se ordena
  currentPage.value = 1;
}

function applyFilters(newFilters: Filters) {
  filters.value = newFilters;
  currentPage.value = 1; // Resetear a la primera página cuando se aplican filtros
}

// Cargar recomendaciones al montar el componente
store.loadRecommendations();
</script>
