<template>
  <div class="overflow-x-auto bg-white shadow-md rounded-lg p-4">
    <TableLoader v-if="loading" />
    <ErrorMessage v-if="errorMessage" :message="errorMessage" />

    <div v-if="!loading && !errorMessage && stocks.length > 0">
      <table class="min-w-full border border-gray-300 rounded-lg" role="table" aria-label="Lista de stocks">
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
          <tr v-for="stock in paginatedStocks" :key="stock.id" class="border-t hover:bg-gray-100 transition-colors"
            :class="{ 'bg-blue-50': selectedStock?.id === stock.id }" @click="selectStock(stock)" role="row"
            tabindex="0">
            <td v-for="header in tableHeaders" :key="header.key" class="px-6 py-3" role="cell">
              <template v-if="header.key === 'target_from' || header.key === 'target_to'">
                {{ formatCurrency(stock[header.key]) }}
              </template>
              <template v-else>
                {{ stock[header.key] || "N/A" }}
              </template>
            </td>
          </tr>
        </tbody>
      </table>

      <TablePagination v-model:current-page="currentPage" :total-items="sortedStocks.length"
        :items-per-page="itemsPerPage" />
    </div>

    <EmptyState v-if="!loading && !errorMessage && stocks.length === 0" />
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, onUnmounted } from "vue"
import type { Stock } from "@/types/stock"
import { useStockStore } from "@/stores/stockStore"
import TableLoader from "@/components/common/TableLoader.vue"
import ErrorMessage from "@/components/common/ErrorMessage.vue"
import EmptyState from "@/components/common/EmptyState.vue"
import TablePagination from "@/components/common/TablePagination.vue"
import { formatCurrency } from "@/utils/formatters"

interface TableHeader {
  key: keyof Stock
  label: string
}

const tableHeaders: TableHeader[] = [
  { key: 'ticker', label: 'Ticker' },
  { key: 'company', label: 'Compañía' },
  { key: 'brokerage', label: 'Brokerage' },
  { key: 'action', label: 'Acción' },
  { key: 'rating_from', label: 'Rating Desde' },
  { key: 'rating_to', label: 'Rating Hasta' },
  { key: 'target_from', label: 'Objetivo Desde' },
  { key: 'target_to', label: 'Objetivo Hasta' }
]

const stockStore = useStockStore()
const stocks = computed<Stock[]>(() => stockStore.filteredStocks)
const loading = computed(() => stockStore.loading)
const errorMessage = computed(() => stockStore.errorMessage)

// Estado para ordenamiento
const sortColumn = ref<keyof Stock | null>(null)
const sortDirection = ref<'ascending' | 'descending'>('ascending')
const selectedStock = ref<Stock | null>(null)

// Estado para paginación
const currentPage = ref(1)
const itemsPerPage = ref(10)

// Ordenamiento de stocks
const sortedStocks = computed(() => {
  if (!sortColumn.value) return stocks.value

  return [...stocks.value].sort((a, b) => {
    const aValue = a[sortColumn.value!]
    const bValue = b[sortColumn.value!]

    if (aValue === null) return 1
    if (bValue === null) return -1

    const comparison = aValue < bValue ? -1 : aValue > bValue ? 1 : 0
    return sortDirection.value === 'ascending' ? comparison : -comparison
  })
})

// Stocks paginados
const paginatedStocks = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  return sortedStocks.value.slice(start, start + itemsPerPage.value)
})

// Funciones de interacción
function sortBy(column: keyof Stock) {
  if (sortColumn.value === column) {
    sortDirection.value = sortDirection.value === 'ascending' ? 'descending' : 'ascending'
  } else {
    sortColumn.value = column
    sortDirection.value = 'ascending'
  }
  // Resetear a la primera página cuando se ordena
  currentPage.value = 1
}

function selectStock(stock: Stock) {
  selectedStock.value = selectedStock.value?.id === stock.id ? null : stock
}

// Keyboard navigation
function handleKeyDown(event: KeyboardEvent) {
  if (event.key === 'Enter' || event.key === ' ') {
    event.preventDefault()
    const target = event.target as HTMLElement
    if (target.tagName === 'TR') {
      target.click()
    }
  }
}

onMounted(() => {
  stockStore.loadStocks()
  document.addEventListener('keydown', handleKeyDown)
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeyDown)
})
</script>

<style scoped>
tr:focus {
  outline: 2px solid var(--color-blue-500);
  outline-offset: -2px;
}
</style>
