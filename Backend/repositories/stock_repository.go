package repositories

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"Backend/config"

	"Backend/models"

	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = config.GetDB()
}
func SetDB(database *gorm.DB) {
	db = database
}

func FetchAndStoreStockData() {
	stockData, err := fetchStockData()
	if err != nil {
		log.Fatalf(" Error fetching data from API: %v", err)
	}

	// Insert into DB only if the stock does not exist
	for _, stock := range stockData.Items {
		targetFrom := parsePrice(stock.TargetFrom)
		targetTo := parsePrice(stock.TargetTo)

		var existingStock models.Stock
		result := db.Where("ticker = ? AND time = ?", stock.Ticker, stock.Time).First(&existingStock)

		if result.RowsAffected == 0 {
			db.Create(&models.Stock{
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

func fetchStockData() (*models.StockResponse, error) {
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
	var stockData models.StockResponse
	err = json.Unmarshal(resp.Body(), &stockData)
	if err != nil {
		return nil, err
	}

	return &stockData, nil
}

func GetStocks(db *gorm.DB, ticker, company, brokerage string) ([]models.Stock, error) {
	var stocks []models.Stock

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
		return nil, result.Error
	}

	return stocks, nil
}

func GetAllStocks(db *gorm.DB) ([]models.Stock, error) {
	var stocks []models.Stock
	result := db.Find(&stocks)

	if result.Error != nil {
		return nil, result.Error
	}

	return stocks, nil
}

func parsePrice(price string) float64 {
	var value float64
	price = strings.Replace(price, "$", "", -1)
	fmt.Sscanf(price, "%f", &value)
	return value
}
