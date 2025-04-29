package repositories

import (
	"context"
	"errors"
	"sport-app-backend/helper"
	"sport-app-backend/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *models.Product) (*models.Product, helper.Error)
	GetAllProduct(ctx context.Context) ([]models.Product, helper.Error)
	GetProductByID(ctx context.Context, id string) (*models.Product, helper.Error)
	UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, helper.Error)
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

func (pr *productRepository) CreateProduct(ctx context.Context, product *models.Product) (*models.Product, helper.Error) {
	err := pr.db.WithContext(ctx).Save(product).Error
	if err != nil {
		return nil, helper.NewInternalServerError("Something went wrong")
	}
	return product, nil
}

func (pr *productRepository) GetAllProduct(ctx context.Context) ([]models.Product, helper.Error) {
	var products []models.Product
	err := pr.db.WithContext(ctx).Find(&products).Error
	if err != nil {
		return nil, helper.NewInternalServerError("Something went wrong")
	}
	return products, nil
}

func (pr *productRepository) GetProductByID(ctx context.Context, id string) (*models.Product, helper.Error) {
	var product models.Product
	err := pr.db.WithContext(ctx).Where("id = ?", id).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewNotFoundError("product not found")
		}
		return nil, helper.NewInternalServerError("Something went wrong")
	}
	return &product, nil
}

func (pr *productRepository) UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, helper.Error) {
	err := pr.db.WithContext(ctx).Save(product).Error
	if err != nil {
		return nil, helper.NewInternalServerError("Something went wrong")
	}
	return product, nil
}

func (pr *productRepository) DeleteProduct(ctx context.Context, id string) (*models.Product, helper.Error) {
	var product models.Product
	err := pr.db.WithContext(ctx).Where("id = ?", id).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewNotFoundError("product not found")
		}
		return nil, helper.NewInternalServerError("Something went wrong")
	}

	err = pr.db.WithContext(ctx).Delete(&product).Error
	if err != nil {
		return nil, helper.NewInternalServerError("Something went wrong")
	}

	return &product, nil	
}