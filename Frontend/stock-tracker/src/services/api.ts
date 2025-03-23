import { $fetch } from 'ofetch'

const API_BASE_URL = 'http://localhost:9090'

export const apiFetch = $fetch.create({
  baseURL: API_BASE_URL,
  retry: 1,
  onRequestError: ({ error }: { error: Error }) => {
    console.error('Error en la petición:', error)
  },
  onResponseError: ({ response }: { response: { _data?: unknown; statusText?: string } }) => {
    console.error('Error en la respuesta:', response?._data || response?.statusText)
  },
})

export interface ApiResponse<T> {
  data: T
  message?: string
  status?: number
}
