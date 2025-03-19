package main

import (
	"log"

	"Backend/config"
	"Backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Inicializar la base de datos
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Configurar el enrutador
	r := gin.Default()

	// Configurar CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// Configurar los manejadores
	stockHandler := handlers.NewStockHandler(db)

	// Definir las rutas
	r.GET("/stocks", stockHandler.GetStocks)
	r.GET("/stocks/recommendations", stockHandler.GetBestStocks)

	// Iniciar el servidor
	log.Println("API running at http://localhost:9090")
	if err := r.Run(":9090"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
