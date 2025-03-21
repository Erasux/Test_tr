package config

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

// InitLogger inicializa los loggers con formato estructurado
func InitLogger() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogError registra un error de forma segura
func LogError(err error, context string) {
	if err != nil {
		ErrorLogger.Printf("[%s] %v", context, err)
	}
}

// LogInfo registra información de forma segura
func LogInfo(message string, context string) {
	InfoLogger.Printf("[%s] %s", context, message)
}
