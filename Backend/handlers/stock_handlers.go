package handlers

import (
	"net/http"

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
func NewStockHandler(db *gorm.DB) *StockHandler {
	return &StockHandler{db: db}
}

// GetStocks obtiene las acciones filtradas por ticker, company y brokerage.
func (h *StockHandler) GetStocks(c *gin.Context) {
	ticker := c.Query("ticker")
	company := c.Query("company")
	brokerage := c.Query("brokerage")

	stocks, err := repositories.GetStocks(h.db, ticker, company, brokerage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch stocks",
			"details": err.Error(),
		})
		return
	}

	if len(stocks) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No stocks found",
			"data":    []models.Stock{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": stocks,
	})
}

// GetBestStocks obtiene las mejores recomendaciones de acciones.
func (h *StockHandler) GetBestStocks(c *gin.Context) {
	stocks, err := repositories.GetAllStocks(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch stocks",
			"details": err.Error(),
		})
		return
	}

	// Configurar el BrokerScorer
	scorer := &services.DefaultBrokerScorer{
		TopBrokers: map[string]float64{
			"The Goldman Sachs Group": 2,
			"JPMorgan Chase":          1.5,
			"Bank of America":         1,
		},
	}

	recommendations := services.CalculateStockRecommendations(stocks, scorer)
	c.JSON(http.StatusOK, gin.H{
		"data": recommendations,
	})
}
