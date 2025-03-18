import axios from "axios";

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || "http://localhost:9090";

export const fetchStocks = async (filters: Record<string, string> = {}) => {
  try {
    const response = await axios.get(`${API_BASE_URL}/stocks`, { params: filters });
    return response.data;
  } catch (error) {
    console.error("Error fetching stocks:", error);
    throw new Error("Failed to fetch stocks");
  }
};

export const fetchStockRecommendations = async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}/stocks/recommendations`);
    return response.data;
  } catch (error) {
    console.error("Error fetching recommendations:", error);
    throw new Error("Failed to fetch recommendations");
  }
};
