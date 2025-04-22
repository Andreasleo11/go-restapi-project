package main

import (
	"go_RESTapi/database"
	"go_RESTapi/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := gin.Default()

	// 💡 Custom CORS config di sini:
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(config))

	// Routes
	routes.SetupBookRoutes(r)

	// Jalankan server di semua IP biar bisa diakses React
	if err := r.Run("0.0.0.0:8000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
