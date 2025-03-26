package main

import (
	"log"

	"Backend/config"
	"Backend/handlers"
	"Backend/middleware"
	"Backend/repositories"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Inicializar logger
	config.InitLogger()

	// Inicializar configuración de seguridad
	securityConfig, err := config.InitSecurityConfig()
	if err != nil {
		log.Fatalf("Error initializing security config: %v", err)
	}

	// Inicializar la base de datos
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Configurar el repositorio de stocks
	repositories.SetDB(db)

	// Actualizar datos de stocks al inicio
	go func() {
		if err := repositories.FetchAndStoreStockData(); err != nil {
			log.Printf("Error fetching initial stock data: %v", err)
		} else {
			log.Println("✅ Datos de stocks actualizados exitosamente")
		}
	}()

	// Configurar el enrutador
	r := gin.Default()

	// Aplicar middleware de seguridad
	r.Use(middleware.SecurityMiddleware())

	// Configurar CORS de forma segura
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = securityConfig.AllowedOrigins
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	// Configurar los manejadores
	stockHandler, err := handlers.NewStockHandler(db)
	if err != nil {
		log.Fatalf("Error creating stock handler: %v", err)
	}

	// Definir las rutas
	r.GET("/stocks", stockHandler.GetStocks)
	r.GET("/stocks/recommendations", stockHandler.GetBestStocks)
	r.POST("/stocks/update", stockHandler.UpdateStocks)

	// Iniciar el servidor
	config.LogInfo("API running at http://localhost:9090", "main")
	if err := r.Run(":9090"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
