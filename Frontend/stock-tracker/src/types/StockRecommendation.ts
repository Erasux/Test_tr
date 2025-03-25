import type { Stock } from './stock'

export interface StockRecommendation extends Stock {
  score: string;
  recommendation: string;
}

export interface StockRecommendationResponse {
  data: {
    stock: Stock;
    score: number;
    recommendation: string;
  }[];
}
