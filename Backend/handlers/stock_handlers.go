package handlers

import (
	"net/http"

	"Backend/models"

	"Backend/repositories"
	"Backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func GetStocks(c *gin.Context) {
	var stocks []models.Stock

	ticker := c.Query("ticker")
	company := c.Query("company")
	brokerage := c.Query("brokerage")

	stocks, err := repositories.GetStocks(db, ticker, company, brokerage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stocks)
}
func GetBestStocks(c *gin.Context) {
	stocks, err := repositories.GetAllStocks(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	recommendations := services.CalculateStockRecommendations(stocks)
	c.JSON(http.StatusOK, recommendations)
}
