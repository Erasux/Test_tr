import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { apiFetch } from '@/services/api'
import type { ApiResponse } from '@/services/api'

interface Stock {
  id: number
  ticker: string
  company: string
  brokerage: string
  action: string
  rating_from: string | null
  rating_to: string | null
  target_from: number | null
  target_to: number | null
}

interface StockFilters {
  ticker: string
  brokerage: string
  company: string
}

export const useStockStore = defineStore('stock', () => {
  const stocks = ref<Stock[]>([])
  const filters = ref<StockFilters>({
    ticker: '',
    brokerage: '',
    company: '',
  })
  const errorMessage = ref<string | null>(null)
  const loading = ref(false)
  const isLoaded = ref(false)

  const filteredStocks = computed(() => {
    return stocks.value.filter((stock) => {
      return (
        (!filters.value.ticker ||
          stock.ticker.toLowerCase().includes(filters.value.ticker.toLowerCase())) &&
        (!filters.value.brokerage ||
          stock.brokerage.toLowerCase().includes(filters.value.brokerage.toLowerCase())) &&
        (!filters.value.company ||
          stock.company.toLowerCase().includes(filters.value.company.toLowerCase()))
      )
    })
  })

  async function loadStocks(forceReload = false) {
    if (isLoaded.value && !forceReload) return

    try {
      loading.value = true
      errorMessage.value = null

      const response = await apiFetch<ApiResponse<Stock[]>>('/stocks')
      stocks.value = response.data
      isLoaded.value = true
    } catch (error) {
      console.error('Error al cargar los stocks:', error)
      errorMessage.value = 'Error al cargar los stocks. Por favor, intente nuevamente más tarde.'
    } finally {
      loading.value = false
    }
  }

  function setFilters(newFilters: StockFilters) {
    filters.value = newFilters
  }

  return { stocks, filteredStocks, errorMessage, loading, loadStocks, setFilters }
})
