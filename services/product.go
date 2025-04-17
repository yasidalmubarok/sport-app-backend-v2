package services

import (
	"context"
	"sport-app-backend/helper"
	"sport-app-backend/models"
	"sport-app-backend/repositories"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product *models.Product) (*models.Product, helper.Error)
	GetAllProduct(ctx context.Context) ([]models.Product, helper.Error)
	GetProductByID(ctx context.Context, id string) (*models.Product, helper.Error)
	UpdateProduct(ctx context.Context, id string, product *models.Product) (*models.Product, helper.Error)
	DeleteProduct(ctx context.Context, id string) (*models.Product, helper.Error)
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) *productService {
	return &productService{
		productRepository: productRepository,
	}
}