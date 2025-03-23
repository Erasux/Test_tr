export const API_CONFIG = {
  BASE_URL: import.meta.env.VITE_API_URL || 'http://localhost:9090',
  TIMEOUT: 5000,
  RETRY_ATTEMPTS: 1,
  RETRY_DELAY: 1000,
}
