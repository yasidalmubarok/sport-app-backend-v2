package repositories

import (
	"context"
	"sport-app-backend/helper"
	"sport-app-backend/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *models.Product) (*models.Product, helper.Error)
	GetAllProduct(ctx context.Context) ([]models.Product, helper.Error)
	GetProductByID(ctx context.Context, id string) (*models.Product, helper.Error)
	UpdateProduct(ctx context.Context, id string, product *models.Product) (*models.Product, helper.Error)
	DeleteProduct(ctx context.Context, id string) (*models.Product, helper.Error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}
