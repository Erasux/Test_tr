import { defineStore } from 'pinia'
import { ref } from 'vue'

interface ErrorState {
  message: string
  code?: string
  details?: unknown
}

export const useErrorStore = defineStore('error', () => {
  const error = ref<ErrorState | null>(null)

  function setError(errorState: ErrorState) {
    error.value = errorState
  }

  function clearError() {
    error.value = null
  }

  return {
    error,
    setError,
    clearError,
  }
})
