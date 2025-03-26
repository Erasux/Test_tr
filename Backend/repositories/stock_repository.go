package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"Backend/config"
	"Backend/models"
	"Backend/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

// SetDB asigna la instancia de la base de datos
func SetDB(database *gorm.DB) {
	db = database
}

func FetchAndStoreStockData() error {
	var totalStocks int64
	var nextPage string
	var pageCount int
	const maxPages = 20

	for {
		pageCount++
		fmt.Printf(" Obteniendo página %d de %d...\n", pageCount, maxPages)

		// Verificar si hemos alcanzado el límite de páginas
		if pageCount >= maxPages {
			fmt.Printf("Límite de %d páginas alcanzado\n", maxPages)
			break
		}

		stockData, err := fetchStockData(nextPage)
		if err != nil {
			return fmt.Errorf("error fetching data from API: %v", err)
		}

		// Convertir los datos de la API a una lista de modelos Stock
		var stocks []models.Stock
		for _, item := range stockData.Items {
			// Convertir el map[string]interface{} a los tipos correctos
			stock := models.Stock{
				Ticker:     item["ticker"].(string),
				Company:    item["company"].(string),
				TargetFrom: utils.ParsePrice(item["target_from"].(string)),
				TargetTo:   utils.ParsePrice(item["target_to"].(string)),
				Action:     item["action"].(string),
				Brokerage:  item["brokerage"].(string),
				RatingFrom: item["rating_from"].(string),
				RatingTo:   item["rating_to"].(string),
				Time:       item["time"].(string),
			}
			stocks = append(stocks, stock)
		}

		// Realizar un UPSERT usando GORM
		result := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "ticker"}, {Name: "time"}},
			DoUpdates: clause.AssignmentColumns([]string{"target_from", "target_to", "action", "brokerage", "rating_from", "rating_to"}),
		}).Create(&stocks)

		if result.Error != nil {
			return fmt.Errorf("error inserting/updating stocks: %v", result.Error)
		}

		totalStocks += result.RowsAffected

		// Verificar si hay más páginas
		nextPage = stockData.NextPage
		if nextPage == "" {
			fmt.Printf("Ultima Pagina\n")
			break
		}

		// Pequeña pausa para no sobrecargar la API
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("✅ Proceso completado:\n")
	fmt.Printf("   - Total de stocks procesados: %d\n", totalStocks)
	return nil
}

func fetchStockData(nextPage string) (*models.StockResponse, error) {
	apiURL := os.Getenv("API_URL")
	apiKey := os.Getenv("API_KEY")

	// Construir la URL con el parámetro next_page si existe
	if nextPage != "" {
		// Verificar si la URL ya tiene un parámetro
		if strings.Contains(apiURL, "?") {
			apiURL = fmt.Sprintf("%s&next_page=%s", apiURL, nextPage)
		} else {
			apiURL = fmt.Sprintf("%s?next_page=%s", apiURL, nextPage)
		}
	}

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

// GetAllStocks obtiene todas las acciones de la base de datos, mostrando solo los registros más recientes por ticker.
func GetAllStocks(db *gorm.DB) ([]models.Stock, error) {
	var stocks []models.Stock

	// Subconsulta para obtener el ID más reciente por ticker
	subQuery := db.Model(&models.Stock{}).
		Select("MAX(id) as id").
		Group("ticker")

	// Consulta principal que obtiene los registros más recientes
	result := db.Where("id IN (?)", subQuery).
		Order("time DESC").
		Find(&stocks)

	if result.Error != nil {
		return nil, result.Error
	}

	return stocks, nil
}

// GetStocks obtiene las acciones filtradas por ticker, company y brokerage, mostrando solo los registros más recientes.
func GetStocks(db *gorm.DB, ticker, company, brokerage string) ([]models.Stock, error) {
	var stocks []models.Stock

	// Subconsulta para obtener el ID más reciente por ticker
	subQuery := db.Model(&models.Stock{}).
		Select("MAX(id) as id").
		Group("ticker")

	// Consulta base con los registros más recientes
	query := db.Where("id IN (?)", subQuery)

	if ticker != "" {
		query = query.Where("ticker = ?", ticker)
	}
	if company != "" {
		query = query.Where("company ILIKE ?", "%"+company+"%")
	}
	if brokerage != "" {
		query = query.Where("brokerage ILIKE ?", "%"+brokerage+"%")
	}

	// Ordenar por fecha descendente
	query = query.Order("time DESC")

	result := query.Find(&stocks)

	if result.Error != nil {
		return nil, result.Error
	}

	return stocks, nil
}
