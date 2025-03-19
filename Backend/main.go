package main

import (
	"Backend/config"
	"Backend/handlers"
	"Backend/repositories"
	"log"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// Inicializar la base de datos
	db := config.InitDB()

	// Pasar la instancia de db a los repositorios
	repositories.SetDB(db)

	// Fetch y almacenar datos de la API
	repositories.FetchAndStoreStockData()

	// Configurar el enrutador
	r := gin.Default()
	r.Use(cors.Default())

	// Configurar los manejadores
	handlers.SetDB(db)

	r.GET("/stocks", handlers.GetStocks)
	r.GET("/stocks/recommendations", handlers.GetBestStocks)

	log.Println("API running at http://localhost:9090")
	r.Run(":9090")
}
