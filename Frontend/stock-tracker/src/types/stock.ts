export interface Stock {
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

export interface StockFilters {
  ticker: string
  brokerage: string
  company: string
}

export interface ApiResponse<T> {
  data: T
  message?: string
  status?: number
}
