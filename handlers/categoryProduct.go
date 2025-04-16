package handlers

import "sport-app-backend/services"

type categoryProductHandler struct {
	categoryProductService services.CategoryProductService
}

func NewCategoryProductHandler(categoryProductService services.CategoryProductService) *categoryProductHandler {
	return &categoryProductHandler{categoryProductService: categoryProductService}
}
