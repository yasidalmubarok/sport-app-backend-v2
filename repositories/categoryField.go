package repositories

import (
	"context"
	"errors"
	"sport-app-backend/helper"
	"sport-app-backend/models"

	"gorm.io/gorm"
)

type CategoryFieldRepository interface {
	CreateCategoryField(ctx context.Context, categoryField *models.CategoryField) (*models.CategoryField, helper.Error)
	GetAllCategoryField(ctx context.Context) ([]models.CategoryField, helper.Error)
	GetCategoryFieldByID(ctx context.Context, id string) (*models.CategoryField, helper.Error)
	UpdateCategoryField(ctx context.Context, id string, categoryField models.CategoryField) (*models.CategoryField, helper.Error)
	DeleteCategoryField(ctx context.Context, id string) (*models.CategoryField, helper.Error)
}

type categoryFieldRepository struct {
	db *gorm.DB
}

func NewCategoryFieldRepository(db *gorm.DB) *categoryFieldRepository {
	return &categoryFieldRepository{
		db: db,
	}
}

func (cfr *categoryFieldRepository) CreateCategoryField(ctx context.Context, categoryField *models.CategoryField) (*models.CategoryField, helper.Error) {
	err := cfr.db.WithContext(ctx).Create(&categoryField).Error
	if err != nil {
		return nil, helper.NewInternalServerError("Something went wrong")
	}
	return categoryField, nil
}

func (cfr *categoryFieldRepository) GetAllCategoryField(ctx context.Context) ([]models.CategoryField, helper.Error) {
	var categoryFields []models.CategoryField
	err := cfr.db.WithContext(ctx).Find(&categoryFields).Error
	if err != nil {
		return nil, helper.NewInternalServerError("Something went wrong")
	}
	return categoryFields, nil
}

func (cfr *categoryFieldRepository) GetCategoryFieldByID(ctx context.Context, id string) (*models.CategoryField, helper.Error) {
	var categoryField models.CategoryField
	err := cfr.db.WithContext(ctx).Where("id = ?", id).First(&categoryField).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewNotFoundError("category field not found")
		}
		return nil, helper.NewInternalServerError("Something went wrong")
	}
	return &categoryField, nil
}

func (cfr *categoryFieldRepository) UpdateCategoryField(ctx context.Context, id string, categoryField models.CategoryField) (*models.CategoryField, helper.Error) {
	var existingCategoryField models.CategoryField
	err := cfr.db.WithContext(ctx).Where("id = ?", id).First(&existingCategoryField).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewNotFoundError("category field not found")
		}
		return nil, helper.NewInternalServerError("Something went wrong")
	}
	existingCategoryField.CategoryName = categoryField.CategoryName
	err = cfr.db.WithContext(ctx).Save(&existingCategoryField).Error
	if err != nil {
		return nil, helper.NewInternalServerError("Something went wrong")
	}
	return &existingCategoryField, nil
}

func (cfr *categoryFieldRepository) DeleteCategoryField(ctx context.Context, id string) (*models.CategoryField, helper.Error) {
	var categoryField models.CategoryField
	err := cfr.db.WithContext(ctx).Where("id = ?", id).First(&categoryField).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewNotFoundError("category field not found")
		}
		return nil, helper.NewInternalServerError("Something went wrong")
	}
	err = cfr.db.WithContext(ctx).Delete(&categoryField).Error
	if err != nil {
		return nil, helper.NewInternalServerError("Something went wrong")
	}
	return &categoryField, nil
}