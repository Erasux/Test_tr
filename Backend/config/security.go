package config

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"time"
)

// SecurityConfig contiene configuraciones de seguridad
type SecurityConfig struct {
	JWTSecret      string
	TokenDuration  time.Duration
	AllowedOrigins []string
}

var securityConfig *SecurityConfig

// InitSecurityConfig inicializa las configuraciones de seguridad
func InitSecurityConfig() (*SecurityConfig, error) {
	if securityConfig != nil {
		return securityConfig, nil
	}

	// Generar un secreto JWT seguro si no está definido
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		secret, err := generateSecureToken(32)
		if err != nil {
			return nil, err
		}
		jwtSecret = secret
	}

	securityConfig = &SecurityConfig{
		JWTSecret:     jwtSecret,
		TokenDuration: 24 * time.Hour,
		AllowedOrigins: []string{
			"http://localhost:3000", // Frontend local
			"https://tudominio.com", // Tu dominio de producción
		},
	}

	return securityConfig, nil
}

// generateSecureToken genera un token seguro
func generateSecureToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// GetSecurityConfig retorna la configuración de seguridad
func GetSecurityConfig() *SecurityConfig {
	if securityConfig == nil {
		InitSecurityConfig()
	}
	return securityConfig
}
