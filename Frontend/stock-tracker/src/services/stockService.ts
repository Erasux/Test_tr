export async function fetchStocks() {
    try {
      const response = await fetch("http://localhost:9090/stocks");
      if (!response.ok) {
        throw new Error("Failed to fetch stocks");
      }
      const data = await response.json();
      return data.data; // Suponiendo que la API devuelve { data: [...] }
    } catch (error) {
      console.error("Error fetching stocks:", error);
      return [];
    }
  }
  