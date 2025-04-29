package services

import (
	"context"
	"sport-app-backend/helper"
	"sport-app-backend/models"
	"sport-app-backend/repositories"
)

type CategoryProductService interface {
	CreateCategoryProduct(ctx context.Context, input *models.CategoryProductRequest) (*models.CategoryProductResponse, helper.Error)
	GetAllCategoryProduct(ctx context.Context) ([]models.CategoryProductResponse, helper.Error)
	GetCategoryProductByID(ctx context.Context, id string) (*models.CategoryProductResponse, helper.Error)
	UpdateCategoryProduct(ctx context.Context, id string, input *models.CategoryProductRequest) (*models.CategoryProductResponse, helper.Error)
	DeleteCategoryProduct(ctx context.Context, id string) (*models.CategoryProductResponse, helper.Error)
	GetCategoryProducts(ctx context.Context, id string) ([]models.Product, helper.Error)
}

type categoryProductService struct {
	categoryProductRepository repositories.CategoryProductRepository
}

func NewCategoryProductService(categoryProductRepository repositories.CategoryProductRepository) *categoryProductService {
	return &categoryProductService{categoryProductRepository: categoryProductRepository}
}

func (cps *categoryProductService) CreateCategoryProduct(ctx context.Context, input *models.CategoryProductRequest) (*models.CategoryProductResponse, helper.Error) {
	categoryProduct := input.NewCategoryProduct()

	savedProduct, err := cps.categoryProductRepository.CreateCategoryProduct(ctx, &categoryProduct)
	if err != nil {
		return nil, err
	}

	return cps.mapCategoryProductResponse(savedProduct), nil
}

func (cps *categoryProductService) GetAllCategoryProduct(ctx context.Context) ([]models.CategoryProductResponse, helper.Error) {
	categoryProducts, err := cps.categoryProductRepository.GetAllCategoryProduct(ctx)
	if err != nil {
		return nil, err
	}

	var categoryProductResponses []models.CategoryProductResponse
	for _, categoryProduct := range categoryProducts {
		categoryProductResponses = append(categoryProductResponses, *cps.mapCategoryProductResponse(&categoryProduct))
	}

	return categoryProductResponses, nil
}

func (cps *categoryProductService) GetCategoryProductByID(ctx context.Context, id string) (*models.CategoryProductResponse, helper.Error) {
	categoryProduct, err := cps.categoryProductRepository.GetCategoryProductByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return cps.mapCategoryProductResponse(categoryProduct), nil
}

func (cps *categoryProductService) UpdateCategoryProduct(ctx context.Context, id string, input *models.CategoryProductRequest) (*models.CategoryProductResponse, helper.Error) {
	categoryProduct, err := cps.categoryProductRepository.GetCategoryProductByID(ctx, id)
	if err != nil {
		return nil, err
	}

	categoryProduct.Name = input.Name

	categoryProduct, err = cps.categoryProductRepository.UpdateCategoryProduct(ctx, id, categoryProduct)
	if err != nil {
		return nil, err
	}

	return cps.mapCategoryProductResponse(categoryProduct), nil
}

func (cps *categoryProductService) DeleteCategoryProduct(ctx context.Context, id string) (*models.CategoryProductResponse, helper.Error) {
	categoryProduct, err := cps.categoryProductRepository.DeleteCategoryProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	return cps.mapCategoryProductResponse(categoryProduct), nil
}

func (cps *categoryProductService) GetCategoryProducts(ctx context.Context, id string) ([]models.Product, helper.Error) {
	products, err := cps.categoryProductRepository.GetCategoryWithProducts(ctx, id)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (cps *categoryProductService) mapCategoryProductResponse(categoryProduct *models.CategoryProduct) *models.CategoryProductResponse {
	return &models.CategoryProductResponse{
		ID:        categoryProduct.ID,
		Name:      categoryProduct.Name,
		CreatedAt: categoryProduct.CreatedAt,
		UpdatedAt: categoryProduct.UpdatedAt,
	}
}
