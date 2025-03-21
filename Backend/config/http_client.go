package config

import (
	"net/http"
	"sync"
	"time"
)

var (
	httpClient *http.Client
	httpOnce   sync.Once
)

// InitHTTPClient inicializa un cliente HTTP compartido con configuraciones seguras
func InitHTTPClient() *http.Client {
	httpOnce.Do(func() {
		httpClient = &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:          100,
				MaxIdleConnsPerHost:   100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				DisableKeepAlives:     false,
				ForceAttemptHTTP2:     true,
			},
		}
	})
	return httpClient
}

// GetHTTPClient retorna la instancia del cliente HTTP
func GetHTTPClient() *http.Client {
	if httpClient == nil {
		return InitHTTPClient()
	}
	return httpClient
}
