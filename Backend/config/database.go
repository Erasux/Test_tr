package config

import (
	"Backend/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "postgresql://root@localhost:26257/stocks_db?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to CockroachDB: %v", err)
	}

	// Auto-migrate the database schema
	if err := db.AutoMigrate(&models.Stock{}); err != nil {
		return nil, fmt.Errorf("error migrating database: %v", err)
	}

	return db, nil
}

func GetDB() *gorm.DB {
	return db
}
