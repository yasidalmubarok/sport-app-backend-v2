package services

import (
	"context"
	"sport-app-backend/helper"
	"sport-app-backend/models"
	"sport-app-backend/repositories"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product *models.CreateProductRequest) (*models.CreateProductResponse, helper.Error)
	GetAllProduct(ctx context.Context) ([]models.CreateProductResponse, helper.Error)
	GetProductByID(ctx context.Context, id string) (*models.CreateProductResponse, helper.Error)
	UpdateProduct(ctx context.Context,id string, product *models.CreateProductRequest) (*models.CreateProductResponse, helper.Error)
	DeleteProduct(ctx context.Context, id string) (*models.CreateProductResponse, helper.Error)
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) *productService {
	return &productService{
		productRepository: productRepository,
	}
}

func (ps *productService) CreateProduct(ctx context.Context, req *models.CreateProductRequest) (*models.CreateProductResponse, helper.Error) {
	product := req.ToProduct().NewProduct()
	result, err := ps.productRepository.CreateProduct(ctx, &product)
	if err != nil {
		return nil, helper.NewInternalServerError("failed to create product")
	}
	return result.ToResponse(), nil
}

func (ps *productService) GetAllProduct(ctx context.Context) ([]models.CreateProductResponse, helper.Error) {
	result, err := ps.productRepository.GetAllProduct(ctx)
	if err != nil {
		return nil, helper.NewInternalServerError("failed to get all product")
	}
	
	var response []models.CreateProductResponse
	for _, product := range result {
		response = append(response, *ps.mapProductsResponse(&product))
	}
	return response, nil
}

func (ps *productService) GetProductByID(ctx context.Context, id string) (*models.CreateProductResponse, helper.Error) {
	result, err := ps.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return nil, helper.NewInternalServerError("failed to get product by id")
	}
	return result.ToResponse(), nil
}

func (ps *productService) UpdateProduct(ctx context.Context,id string, input *models.CreateProductRequest) (*models.CreateProductResponse, helper.Error) {
	product, err := ps.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return nil, helper.NewInternalServerError("failed to get product by id")
	}

	product.Name = input.Name
	product.Category = input.Category
	product.PriceSell = input.PriceSell
	product.PriceBuy = input.PriceBuy
	product.Stock = input.Stock
	product.Status = input.Status

	result, err := ps.productRepository.UpdateProduct(ctx, product)
	if err != nil {
		return nil, helper.NewInternalServerError("failed to update product")
	}
	return result.ToResponse(), nil
}

func (ps *productService) DeleteProduct(ctx context.Context, id string) (*models.CreateProductResponse, helper.Error) {
	result, err := ps.productRepository.DeleteProduct(ctx, id)
	if err != nil {
		return nil, helper.NewInternalServerError("failed to delete product")
	}
	return result.ToResponse(), nil
}

func (ps *productService) mapProductsResponse(product *models.Product) *models.CreateProductResponse {
	return &models.CreateProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Category:  product.Category,
		PriceSell: product.PriceSell,
		PriceBuy:  product.PriceBuy,
		Stock:     product.Stock,
		Status:    product.Status,
		Image:     product.Image,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}