package main

import (
	"fmt"
	"log"
	"os"
	"sport-app-backend/config"
	"sport-app-backend/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Bisa disesuaikan
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	userOwnerRepository := repositories.NewUserOwnerRepository(db)

	// Tambahkan router di sini...

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server running on port %s", port)
	router.Run(fmt.Sprintf(":%s", port))
}
