package config

import (
	"Backend/models"
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}
	envMap, err := godotenv.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading environment variables: %v", err)
	}
	dsn := envMap["DB_URL"]

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
