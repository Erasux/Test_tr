package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"Backend/config"
	"Backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

// SetDB asigna la instancia de la base de datos
func SetDB(database *gorm.DB) {
	db = database
}

func FetchAndStoreStockData() error {
	stockData, err := fetchStockData()
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

func fetchStockData() (*models.StockResponse, error) {
	apiURL := os.Getenv("API_URL")
	apiKey := os.Getenv("API_KEY")

	// Crear una nueva solicitud HTTP
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	// Agregar headers
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Usar el cliente HTTP compartido
	client := config.GetHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Verificar el código de estado
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	// Parsear la respuesta JSON
	var stockData models.StockResponse
	if err := json.Unmarshal(body, &stockData); err != nil {
		return nil, fmt.Errorf("error parsing API response: %v", err)
	}

	return &stockData, nil
}

func parsePrice(price string) float64 {
	var value float64
	price = strings.Replace(price, "$", "", -1)
	fmt.Sscanf(price, "%f", &value)
	return value
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

// GetAllStocks obtiene todas las acciones de la base de datos.
func GetAllStocks(db *gorm.DB) ([]models.Stock, error) {
	var stocks []models.Stock
	result := db.Find(&stocks)

	if result.Error != nil {
		return nil, result.Error
	}

	return stocks, nil
}

// MockStockRepository es una implementación mock de StockRepository.
type MockStockRepository struct {
	GetStocksFunc    func(db *gorm.DB, ticker, company, brokerage string) ([]models.Stock, error)
	GetAllStocksFunc func(db *gorm.DB) ([]models.Stock, error)
}

func (m *MockStockRepository) GetStocks(db *gorm.DB, ticker, company, brokerage string) ([]models.Stock, error) {
	if m.GetStocksFunc != nil {
		return m.GetStocksFunc(db, ticker, company, brokerage)
	}
	return nil, nil
}

func (m *MockStockRepository) GetAllStocks(db *gorm.DB) ([]models.Stock, error) {
	if m.GetAllStocksFunc != nil {
		return m.GetAllStocksFunc(db)
	}
	return nil, nil
}
