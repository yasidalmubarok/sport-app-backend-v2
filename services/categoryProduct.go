package services

import (
	"context"
	"sport-app-backend/helper"
	"sport-app-backend/models"
	"sport-app-backend/repositories"
)

type CategoryProductService interface {
	CreateCategoryProduct(ctx context.Context, categoryProduct *models.CategoryProductRequest) (*models.CategoryProductResponse, helper.Error)
	GetAllCategoryProduct(ctx context.Context) ([]models.CategoryProductResponse, helper.Error)
	GetCategoryProductByID(ctx context.Context, id int) (*models.CategoryProductResponse, helper.Error)
	UpdateCategoryProduct(ctx context.Context, categoryProduct *models.CategoryProductRequest) (*models.CategoryProductResponse, helper.Error)
	DeleteCategoryProduct(ctx context.Context, id int) helper.Error
}

type categoryProductService struct {
	categoryProductRepository repositories.CategoryProductRepository
}

func NewCategoryProductService(categoryProductRepository repositories.CategoryProductRepository) *categoryProductService {
	return &categoryProductService{categoryProductRepository: categoryProductRepository}
}