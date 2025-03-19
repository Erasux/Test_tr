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
	// Cargar variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Inicializar la base de datos
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Pasar la instancia de db a los repositorios
	repositories.SetDB(db)

	// Fetch y almacenar datos de la API (obtener 5 páginas)
	if err := repositories.FetchAndStoreStockData(1); err != nil {
		log.Fatalf("Error fetching and storing stock data: %v", err)
	}

	// Configurar el enrutador
	r := gin.Default()
	r.Use(cors.Default())

	// Configurar los manejadores
	handlers.SetDB(db)

	r.GET("/stocks", handlers.GetStocks)
	r.GET("/stocks/recommendations", handlers.GetBestStocks)

	log.Println("API running at http://localhost:9090")
	if err := r.Run(":9090"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
