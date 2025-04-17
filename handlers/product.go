package handlers

import "sport-app-backend/services"

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *productHandler {
	return &productHandler{productService: productService}
}