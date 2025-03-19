package repositories

import (
	"Backend/config"
	"Backend/models"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

func init() {
	db = config.GetDB()
}
func SetDB(database *gorm.DB) {
	db = database
}

func FetchAndStoreStockData(maxPages int) error {
	stockData, err := fetchStockData(maxPages)
	if err != nil {
		return fmt.Errorf("error fetching data from API: %v", err)
	}

	// Convertir los datos de la API a una lista de modelos Stock
	var stocks []models.Stock
	for _, stock := range stockData.Items {
		stocks = append(stocks, models.Stock{
			Ticker:     stock.Ticker,
			Company:    stock.Company,
			TargetFrom: parsePrice(stock.TargetFrom),
			TargetTo:   parsePrice(stock.TargetTo),
			Action:     stock.Action,
			Brokerage:  stock.Brokerage,
			RatingFrom: stock.RatingFrom,
			RatingTo:   stock.RatingTo,
			Time:       stock.Time,
		})
	}

	// Realizar un UPSERT usando GORM
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "ticker"}, {Name: "time"}},                                                                 // Conflict keys
		DoUpdates: clause.AssignmentColumns([]string{"target_from", "target_to", "action", "brokerage", "rating_from", "rating_to"}), // Campos a actualizar
	}).Create(&stocks)

	if result.Error != nil {
		return fmt.Errorf("error inserting/updating stocks: %v", result.Error)
	}

	fmt.Printf("📊 Inserted/Updated %d stocks\n", result.RowsAffected)
	return nil
}

func fetchStockData(maxPages int) (*models.StockResponse, error) {
	apiBaseURL := os.Getenv("API_URL") // URL base de la API
	apiKey := os.Getenv("API_KEY")

	client := resty.New()
	var allItems []models.StockData
	nextPage := ""

	for page := 1; page <= maxPages; page++ {
		url := apiBaseURL
		if nextPage != "" {
			// Construir la URL completa usando el identificador de la siguiente página
			url = fmt.Sprintf("%s/%s", apiBaseURL, nextPage)
		}

		resp, err := client.R().
			SetHeader("Authorization", "Bearer "+apiKey).
			SetHeader("Content-Type", "application/json").
			Get(url)

		if err != nil {
			return nil, fmt.Errorf("error fetching data from API (page %d): %v", page, err)
		}

		if resp.IsError() {
			return nil, fmt.Errorf("API error (page %d): %s", page, resp.Status())
		}

		// Parse API response
		var stockData models.StockResponse
		if err := json.Unmarshal(resp.Body(), &stockData); err != nil {
			return nil, fmt.Errorf("error parsing API response (page %d): %v", page, err)
		}

		// Agregar los ítems de esta página a la lista completa
		allItems = append(allItems, stockData.Items...)

		// Verificar si hay más páginas
		nextPage = stockData.NextPage
		if nextPage == "" {
			break // No hay más páginas
		}
	}

	// Devolver todos los ítems concatenados
	return &models.StockResponse{
		Items:    allItems,
		NextPage: "", // No es necesario devolver nextPage aquí
	}, nil
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
