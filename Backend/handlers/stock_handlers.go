package handlers

import (
	"errors"
	"net/http"

	"Backend/config"
	"Backend/models"
	"Backend/repositories"
	"Backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// StockHandler define los manejadores para las rutas relacionadas con las acciones.
type StockHandler struct {
	db *gorm.DB
}

// NewStockHandler crea una nueva instancia de StockHandler.
// Retorna error si la base de datos es nil.
func NewStockHandler(db *gorm.DB) (*StockHandler, error) {
	if db == nil {
		return nil, errors.New("la base de datos no puede ser nil")
	}
	return &StockHandler{db: db}, nil
}

// GetStocks obtiene las acciones filtradas por ticker, company y brokerage.
// Si no se proporcionan filtros, muestra todos los datos.
func (h *StockHandler) GetStocks(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "conexión a base de datos no inicializada",
		})
		return
	}

	// Validación y sanitización de parámetros
	ticker := services.SanitizeInput(c.Query("ticker"))
	company := services.SanitizeInput(c.Query("company"))
	brokerage := services.SanitizeInput(c.Query("brokerage"))

	// Limitar la longitud de los parámetros para prevenir ataques
	if len(ticker) > 10 || len(company) > 100 || len(brokerage) > 100 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "parámetros de búsqueda demasiado largos",
		})
		return
	}

	var stocks []models.Stock
	var err error

	// Si no hay filtros, obtener todos los stocks
	if ticker == "" && company == "" && brokerage == "" {
		stocks, err = repositories.GetAllStocks(h.db)
	} else {
		stocks, err = repositories.GetStocks(h.db, ticker, company, brokerage)
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "no se encontraron acciones",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error al obtener acciones",
		})
		return
	}

	// Verificar si se encontraron stocks
	if len(stocks) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no se encontraron acciones",
		})
		return
	}

	// Obtener la fecha más reciente
	var latestTime string
	if len(stocks) > 0 {
		latestTime = stocks[0].Time
		for _, stock := range stocks {
			if stock.Time > latestTime {
				latestTime = stock.Time
			}
		}
	}

	// Preparar la respuesta
	response := gin.H{
		"data": stocks,
		"metadata": gin.H{
			"total_records": len(stocks),
			"last_update":   latestTime,
			"filters_applied": gin.H{
				"ticker":    ticker != "",
				"company":   company != "",
				"brokerage": brokerage != "",
			},
		},
	}

	c.JSON(http.StatusOK, response)
}

// GetBestStocks obtiene las mejores recomendaciones de acciones.
// Implementa validación y manejo de errores mejorado.
func (h *StockHandler) GetBestStocks(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "conexión a base de datos no inicializada",
		})
		return
	}

	stocks, err := repositories.GetAllStocks(h.db)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "no se encontraron acciones",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error al obtener acciones",
		})
		return
	}

	// Verificar si se encontraron stocks
	if len(stocks) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no se encontraron acciones",
		})
		return
	}

	// Configurar el BrokerScorer con valores constantes
	scorer := services.NewDefaultBrokerScorer(map[string]float64{
		"The Goldman Sachs Group": 2.0,
		"JPMorgan Chase":          1.5,
		"Bank of America":         1.0,
	})

	recommendations := services.CalculateStockRecommendations(stocks, scorer)

	c.JSON(http.StatusOK, gin.H{
		"data": recommendations,
	})
}

// UpdateStocks actualiza los datos de stocks desde la API
func (h *StockHandler) UpdateStocks(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "conexión a base de datos no inicializada",
		})
		return
	}

	// Ejecutar la actualización en una goroutine para no bloquear la respuesta
	go func() {
		if err := repositories.FetchAndStoreStockData(); err != nil {
			config.LogError(err, "Error actualizando datos de stocks")
		} else {
			config.LogInfo("✅ Datos de stocks actualizados exitosamente", "UpdateStocks")
		}
	}()

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Actualización de datos iniciada",
	})
}
