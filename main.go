package main

import (
	"fmt"
	"log"
	"os"
	"sport-app-backend/config"
	"sport-app-backend/handlers"
	"sport-app-backend/repositories"
	"sport-app-backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()

	userOwnerRepository := repositories.NewUserOwnerRepository(db)
	userOwnerService := services.NewUserOwnerService(userOwnerRepository)
	userOwnerHandler := handlers.NewUserOwnerHandler(userOwnerService)

	gin.SetMode(gin.ReleaseMode)
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

	api := router.Group("/api/v1")

	userOwner := api.Group("/owner")
	userOwner.POST("/register", userOwnerHandler.CreateUserOwner)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server running on port %s", port)
	router.Run(fmt.Sprintf(":%s", port))
}
