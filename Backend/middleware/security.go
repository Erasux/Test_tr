package middleware

import (
	"net/http"

	"Backend/config"

	"github.com/gin-gonic/gin"
)

// SecurityMiddleware implementa medidas de seguridad básicas
func SecurityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Prevenir clickjacking
		c.Header("X-Frame-Options", "DENY")
		// Habilitar XSS protection
		c.Header("X-XSS-Protection", "1; mode=block")
		// Prevenir MIME type sniffing
		c.Header("X-Content-Type-Options", "nosniff")
		// Configurar CSP
		c.Header("Content-Security-Policy", "default-src 'self'")
		// Configurar HSTS
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		// Validar origen de la petición
		origin := c.GetHeader("Origin")
		if origin != "" {
			allowedOrigins := config.GetSecurityConfig().AllowedOrigins
			isAllowed := false
			for _, allowed := range allowedOrigins {
				if origin == allowed {
					isAllowed = true
					break
				}
			}
			if !isAllowed {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"error": "origen no permitido",
				})
				return
			}
		}

		// Rate limiting básico (implementar con Redis en producción)
		clientIP := c.ClientIP()
		if isRateLimited(clientIP) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "demasiadas peticiones",
			})
			return
		}

		c.Next()
	}
}

// isRateLimited implementa un rate limiting básico para pruebas locales
func isRateLimited(ip string) bool {

	return false // Deshabilitado para pruebas
}
