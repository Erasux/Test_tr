import axios from "axios";

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || "http://localhost:9090";

/**
 * Fetch stocks from the API with optional filters.
 * @param filters - Optional filters for ticker, company, and brokerage.
 * @returns An array of stocks.
 * @throws Error if the request fails or the response format is invalid.
 */
export const fetchStocks = async (filters: Record<string, string> = {}): Promise<any[]> => {
  try {
    const response = await axios.get(`${API_BASE_URL}/stocks`, { params: filters });

    // Verificar la estructura de la respuesta
    if (!response.data || !Array.isArray(response.data.data)) {
      throw new Error("Invalid response format: expected an array of stocks");
    }

    return response.data.data; // Extraer el array de la clave "data"
  } catch (error) {
    console.error("Error fetching stocks:", error);
    throw new Error("Failed to fetch stocks. Please try again later.");
  }
};
/**
 * Fetch stock recommendations from the API.
 * @returns An array of stock recommendations.
 * @throws Error if the request fails or the response format is invalid.
 */
export const fetchStockRecommendations = async (): Promise<any[]> => {
  try {
    const response = await axios.get(`${API_BASE_URL}/stocks/recommendations`);

    // Verificar la estructura de la respuesta
    if (!response.data || !Array.isArray(response.data.data)) {
      throw new Error("Invalid response format: expected an array of recommendations");
    }

    return response.data.data;
  } catch (error) {
    console.error("Error fetching recommendations:", error);
    throw new Error("Failed to fetch recommendations. Please try again later.");
  }
};