package handlers

import (
	"errors"
	"net/http"

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
// Implementa validación de parámetros y manejo de errores mejorado.
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

	stocks, err := repositories.GetStocks(h.db, ticker, company, brokerage)
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

	c.JSON(http.StatusOK, gin.H{
		"data": stocks,
	})
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
