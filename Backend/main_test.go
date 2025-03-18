package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 🔹 Initialize test database (CockroachDB)
func setupTestDB() {
	var err error
	dsn := "postgresql://root@localhost:26257/stocks_db?sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Error connecting to CockroachDB for tests: %v", err)
	}

	// Migrate schema in the test database
	db.AutoMigrate(&Stock{})

	// Ensure stock does not already exist before inserting
	var existing Stock
	result := db.Where("ticker = ?", "AAPL").First(&existing)

	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		// Insert stock data only if it does not exist
		db.Create(&Stock{
			Ticker:     "AAPL",
			Company:    "Apple Inc.",
			TargetFrom: 150.00,
			TargetTo:   180.00,
			Action:     "upgraded by",
			Brokerage:  "The Goldman Sachs Group",
			RatingFrom: "Neutral",
			RatingTo:   "Buy",
		})
	}
}

// 🔹 Test stock score calculation
func TestCalculateStockScore(t *testing.T) {
	setupTestDB() // Initialize test DB

	stock := Stock{
		Ticker:     "AAPL",
		Company:    "Apple Inc.",
		TargetFrom: 150.00,
		TargetTo:   180.00,
		Action:     "upgraded by",
		Brokerage:  "The Goldman Sachs Group",
		RatingFrom: "Neutral",
		RatingTo:   "Buy",
	}

	score := calculateStockScore(stock)

	assert.GreaterOrEqual(t, score, 7.0, "The stock score should be at least 7.0 for a strong recommendation.")
}

// 🔹 Test fetching best stock recommendations
func TestGetBestStocks(t *testing.T) {
	setupTestDB() // Initialize test DB

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/stocks/recommendations", getBestStocks)

	req, _ := http.NewRequest("GET", "/stocks/recommendations", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "The endpoint should return HTTP 200 OK")
	assert.Contains(t, w.Body.String(), "score", "The response should contain stock scores")
}

// 🔹 Test fetching all stocks
func TestGetStocks(t *testing.T) {
	setupTestDB() // Initialize test DB

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/stocks", getStocks)

	req, _ := http.NewRequest("GET", "/stocks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "The endpoint should return HTTP 200 OK")
	assert.Contains(t, w.Body.String(), "ticker", "The response should contain stock ticker information")
}
