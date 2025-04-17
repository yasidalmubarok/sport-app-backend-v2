package main

import (
	"fmt"
	"log"
	"os"
	"sport-app-backend/config"
	"sport-app-backend/handlers"
	"sport-app-backend/middlewares"
	"sport-app-backend/repositories"
	"sport-app-backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	redis := config.ConnectRedis()
	db := config.ConnectDB()

	// Dependency Injection
	userOwnerRepository := repositories.NewUserOwnerRepository(db, redis)
	userOwnerService := services.NewUserOwnerService(userOwnerRepository)
	userOwnerHandler := handlers.NewUserOwnerHandler(userOwnerService)

	categoryProductRepository := repositories.NewCategoryProductRepository(db)
	categoryProductService := services.NewCategoryProductService(categoryProductRepository)
	categoryProductHandler := handlers.NewCategoryProductHandler(categoryProductService)

	authService := middlewares.NewAuthService(db, userOwnerRepository)

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
	userOwner.POST("/session", userOwnerHandler.LoginUserOwner)
	userOwner.POST("/request-otp", userOwnerHandler.RequestResetPasswordHandler)
	userOwner.POST("/verify-otp", userOwnerHandler.ResetPasswordHandler)

	categoryProduct := api.Group("/category-product")
	categoryProduct.POST("/category", authService.AuthMiddleware(), categoryProductHandler.CreateCategoryProduct)
	categoryProduct.GET("/category", categoryProductHandler.GetAllCategoryProduct)
	categoryProduct.GET("/category/:id", categoryProductHandler.GetCategoryProductByID)
	categoryProduct.PUT("/category/:id", authService.AuthMiddleware(), categoryProductHandler.UpdateCategoryProduct)
	categoryProduct.DELETE("/category/:id", authService.AuthMiddleware(), categoryProductHandler.DeleteCategoryProduct)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server running on port %s", port)
	router.Run(fmt.Sprintf(":%s", port))
}
