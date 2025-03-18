package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global database instance
var db *gorm.DB

// Stock model for GORM
type Stock struct {
	ID         int64   `gorm:"primaryKey" json:"id"`
	Ticker     string  `json:"ticker"`
	Company    string  `json:"company"`
	TargetFrom float64 `json:"target_from"`
	TargetTo   float64 `json:"target_to"`
	Action     string  `json:"action"`
	Brokerage  string  `json:"brokerage"`
	RatingFrom string  `json:"rating_from"`
	RatingTo   string  `json:"rating_to"`
	Time       string  `json:"time"`
}

// Stock API response structure
type StockResponse struct {
	Items    []StockData `json:"items"`
	NextPage string      `json:"next_page"`
}

type StockData struct {
	Ticker     string `json:"ticker"`
	Company    string `json:"company"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

type StockRecommendation struct {
	Stock          Stock   `json:"stock"`
	Score          float64 `json:"score"`
	Recommendation string  `json:"recommendation"`
}

// Initialize database connection
func initDB() {
	var err error
	dsn := "postgresql://root@localhost:26257/stocks_db?sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(" Error connecting to CockroachDB: %v", err)
	}

	// Auto-migrate the database schema
	db.AutoMigrate(&Stock{})
}

// Fetch stock data from API and store it in the database
func fetchAndStoreStockData() {
	stockData, err := fetchStockData()
	if err != nil {
		log.Fatalf(" Error fetching data from API: %v", err)
	}

	// Insert into DB only if the stock does not exist
	for _, stock := range stockData.Items {
		targetFrom := parsePrice(stock.TargetFrom)
		targetTo := parsePrice(stock.TargetTo)

		var existingStock Stock
		result := db.Where("ticker = ? AND time = ?", stock.Ticker, stock.Time).First(&existingStock)

		if result.RowsAffected == 0 {
			db.Create(&Stock{
				Ticker:     stock.Ticker,
				Company:    stock.Company,
				TargetFrom: targetFrom,
				TargetTo:   targetTo,
				Action:     stock.Action,
				Brokerage:  stock.Brokerage,
				RatingFrom: stock.RatingFrom,
				RatingTo:   stock.RatingTo,
				Time:       stock.Time,
			})
			fmt.Printf(" Inserted stock: %s (%s)\n", stock.Ticker, stock.Time)
		} else {
			fmt.Printf("🔄 Stock already exists: %s (%s), skipping insert\n", stock.Ticker, stock.Time)
		}
	}
	fmt.Println("📊 Data storage process completed.")
}

// Fetch stock data from the API
func fetchStockData() (*StockResponse, error) {
	apiURL := os.Getenv("API_URL")
	apiKey := os.Getenv("API_KEY")

	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+apiKey).
		SetHeader("Content-Type", "application/json").
		Get(apiURL)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("API error: %s", resp.Status())
	}

	// Parse API response
	var stockData StockResponse
	err = json.Unmarshal(resp.Body(), &stockData)
	if err != nil {
		return nil, err
	}

	return &stockData, nil
}

// Convert price string to float64
func parsePrice(price string) float64 {
	var value float64
	price = strings.Replace(price, "$", "", -1)
	fmt.Sscanf(price, "%f", &value)
	return value
}

// Get all stocks with optional filters
func getStocks(c *gin.Context) {
	var stocks []Stock

	ticker := c.Query("ticker")
	company := c.Query("company")
	brokerage := c.Query("brokerage")

	query := db

	if ticker != "" {
		query = query.Where("ticker = ?", ticker)
	}
	if company != "" {
		query = query.Where("company ILIKE ?", "%"+company+"%")
	}
	if brokerage != "" {
		query = query.Where("brokerage ILIKE ?", "%"+brokerage+"%")
	}

	result := query.Find(&stocks)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, stocks)
}

// Get best stocks based on calculated score
func getBestStocks(c *gin.Context) {
	var stocks []Stock
	db.Find(&stocks)

	recommendations := make([]StockRecommendation, 0)

	for _, stock := range stocks {
		score := calculateStockScore(stock)
		recommendation := ""

		if score >= 7 {
			recommendation = "Strong Buy"
		} else if score >= 5 {
			recommendation = "Buy"
		} else if score >= 3 {
			recommendation = "Hold"
		} else {
			recommendation = "Sell"
		}

		recommendations = append(recommendations, StockRecommendation{
			Stock: stock, Score: score, Recommendation: recommendation,
		})
	}

	// Sort by highest score
	sort.Slice(recommendations, func(i, j int) bool {
		return recommendations[i].Score > recommendations[j].Score
	})

	c.JSON(http.StatusOK, recommendations)
}

// Calculate stock score based on target price, ratings, and brokerage
func calculateStockScore(stock Stock) float64 {
	score := 0.0

	// Target price increase impact
	priceIncrease := stock.TargetTo - stock.TargetFrom
	if priceIncrease > 0 {
		score += 3
		if priceIncrease > 10 {
			score += 2
		}
	}

	// Rating changes impact
	if stock.RatingFrom == "Sell" && stock.RatingTo == "Buy" {
		score += 3
	} else if stock.RatingFrom == "Neutral" && stock.RatingTo == "Buy" {
		score += 2
	} else if stock.RatingFrom == "Sell" && stock.RatingTo == "Neutral" {
		score += 1
	}

	// Weight based on brokerage reputation
	topBrokers := map[string]float64{
		"The Goldman Sachs Group": 2,
		"JPMorgan Chase":          1.5,
		"Bank of America":         1,
	}

	if value, exists := topBrokers[stock.Brokerage]; exists {
		score += value
	}

	return score
}

func main() {
	_ = godotenv.Load()
	initDB()
	fetchAndStoreStockData()

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/stocks", getStocks)
	r.GET("/stocks/recommendations", getBestStocks)

	log.Println("API running at http://localhost:9090")
	r.Run(":9090")
}
