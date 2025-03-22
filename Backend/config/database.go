package config

import (
	"Backend/models"
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

// InitDB inicializa la conexión a la base de datos con configuraciones seguras
func InitDB() (*gorm.DB, error) {
	var initErr error

	dbOnce.Do(func() {
		// Cargar variables de entorno de forma segura
		if err := godotenv.Load(); err != nil {
			initErr = fmt.Errorf("error al cargar archivo .env: %v", err)
			return
		}

		envMap, err := godotenv.Read()
		if err != nil {
			initErr = fmt.Errorf("error al leer variables de entorno: %v", err)
			return
		}

		dsn, exists := envMap["DB_URL"]
		if !exists || dsn == "" {
			initErr = fmt.Errorf("DB_URL no está definida en las variables de entorno")
			return
		}

		// Configuración segura de la base de datos
		config := &gorm.Config{
			Logger: logger.New(
				log.Default(),
				logger.Config{
					SlowThreshold:             time.Second,
					LogLevel:                  logger.Error,
					IgnoreRecordNotFoundError: true,
				},
			),
		}

		// Establecer conexión con timeout
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		db, err = gorm.Open(postgres.Open(dsn), config)
		if err != nil {
			initErr = fmt.Errorf("error al conectar con la base de datos: %v", err)
			return
		}

		// Configurar pool de conexiones
		sqlDB, err := db.DB()
		if err != nil {
			initErr = fmt.Errorf("error al configurar pool de conexiones: %v", err)
			return
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)

		// Auto-migrar el esquema de la base de datos
		if err := db.AutoMigrate(&models.Stock{}); err != nil {
			initErr = fmt.Errorf("error al migrar la base de datos: %v", err)
			return
		}
	})

	return db, initErr
}

// GetDB retorna la instancia de la base de datos
// Retorna nil si la base de datos no ha sido inicializada
func GetDB() *gorm.DB {
	if db == nil {
		log.Println("Advertencia: Intentando acceder a la base de datos no inicializada")
	}
	return db
}
