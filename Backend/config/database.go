package config

import (
	"Backend/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	dsn := "postgresql://root@localhost:26257/stocks_db?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(" Error connecting to CockroachDB: %v", err)
	}

	// Auto-migrate the database schema
	db.AutoMigrate(&models.Stock{})

	return db
}
func GetDB() *gorm.DB {
	return db
}
