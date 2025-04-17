package repositories

import (
	"context"
	"errors"
	"sport-app-backend/helper"
	"sport-app-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryProductRepository interface {
	CreateCategoryProduct(ctx context.Context, categoryProduct *models.CategoryProduct) (*models.CategoryProduct, helper.Error)
	GetAllCategoryProduct(ctx context.Context) ([]models.CategoryProduct, helper.Error)
	GetCategoryProductByID(ctx context.Context, id string) (*models.CategoryProduct, helper.Error)
	UpdateCategoryProduct(ctx context.Context, id string, categoryProduct *models.CategoryProduct) (*models.CategoryProduct, helper.Error)
	DeleteCategoryProduct(ctx context.Context, id string) (*models.CategoryProduct, helper.Error)
}

type categoryProductRepository struct {
	db *gorm.DB
}

func NewCategoryProductRepository(db *gorm.DB) *categoryProductRepository {
	return &categoryProductRepository{
		db: db,
	}
}

func (cpr *categoryProductRepository) CreateCategoryProduct(ctx context.Context, categoryProduct *models.CategoryProduct) (*models.CategoryProduct, helper.Error) {
	err := cpr.db.WithContext(ctx).Create(categoryProduct).Error
	if err != nil {
		return nil, helper.NewInternalServerError("Something went wrong")
	}
	return categoryProduct, nil
}

func (cpr *categoryProductRepository) GetAllCategoryProduct(ctx context.Context) ([]models.CategoryProduct, helper.Error) {
	var categoryProducts []models.CategoryProduct
	err := cpr.db.WithContext(ctx).Find(&categoryProducts).Error
	if err != nil {
		return nil, helper.NewInternalServerError("Something went wrong")
	}
	return categoryProducts, nil
}

func (cpr *categoryProductRepository) GetCategoryProductByID(ctx context.Context, id string) (*models.CategoryProduct, helper.Error) {
	var categoryProduct models.CategoryProduct

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, helper.NewBadRequestError("invalid UUID format")
	}

	err = cpr.db.WithContext(ctx).Where("id = ?", parsedID).First(&categoryProduct).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewNotFoundError("category product not found")
		}
		return nil, helper.NewInternalServerError("Something went wrong")
	}
	return &categoryProduct, nil
}

func (cpr *categoryProductRepository) UpdateCategoryProduct(ctx context.Context, id string, input *models.CategoryProduct) (*models.CategoryProduct, helper.Error) {
	var existingCategoryProduct *models.CategoryProduct

	if err := cpr.db.WithContext(ctx).Where("id = ?", id).First(&existingCategoryProduct).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewNotFoundError("category product not found")
		}
		return nil, helper.NewInternalServerError("Something went wrong")
	}

	existingCategoryProduct.Name = input.Name

	if err := cpr.db.WithContext(ctx).Save(&existingCategoryProduct).Error; err != nil {
		return nil, helper.NewInternalServerError("Something went wrong")
	}

	return existingCategoryProduct, nil
}

func (cpr *categoryProductRepository) DeleteCategoryProduct(ctx context.Context, id string) (*models.CategoryProduct, helper.Error) {
	var categoryProduct models.CategoryProduct
	if err := cpr.db.WithContext(ctx).Where("id = ?", id).First(&categoryProduct).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewNotFoundError("category product not found")
		}
		return nil, helper.NewInternalServerError("Something went wrong")
	}

	if err := cpr.db.WithContext(ctx).Delete(&categoryProduct).Error; err != nil {
		return nil, helper.NewInternalServerError("Something went wrong")
	}

	return &categoryProduct, nil
}
