<template>
  <div class="flex flex-wrap gap-4 mb-6 bg-white p-4 shadow-md rounded-lg">
    <input v-model="filters.ticker" placeholder="Buscar por Ticker"
      class="p-2 border border-gray-300 rounded shadow-sm focus:ring focus:ring-blue-300" @input="debouncedFilter" />
    <input v-model="filters.brokerage" placeholder="Buscar por Brokerage"
      class="p-2 border border-gray-300 rounded shadow-sm focus:ring focus:ring-blue-300" @input="debouncedFilter" />
    <input v-model="filters.company" placeholder="Buscar por Company"
      class="p-2 border border-gray-300 rounded shadow-sm focus:ring focus:ring-blue-300" @input="debouncedFilter" />
    <button @click="applyFilters" :disabled="isLoading"
      class="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600 transition disabled:opacity-50 disabled:cursor-not-allowed">
      {{ isLoading ? 'Filtrando...' : 'Aplicar Filtros' }}
    </button>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import { useStockStore } from "@/stores/stockStore";

interface Filters {
  ticker: string;
  brokerage: string;
  company: string;
}

const stockStore = useStockStore();
const isLoading = computed(() => stockStore.loading);

const filters = ref<Filters>({
  ticker: "",
  brokerage: "",
  company: "",
});

const emit = defineEmits(["filter-changed"]);

// Función para crear un debounce
function debounce<T extends (...args: unknown[]) => void>(
  fn: T,
  delay: number
): (...args: Parameters<T>) => void {
  let timeoutId: ReturnType<typeof setTimeout>;
  return (...args: Parameters<T>) => {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => fn(...args), delay);
  };
}

// Crear versión debounced de la función de filtrado
const debouncedFilter = debounce(() => {
  applyFilters();
}, 300);

const applyFilters = () => {
  if (!isLoading.value) {
    emit("filter-changed", filters.value);
  }
};
</script>