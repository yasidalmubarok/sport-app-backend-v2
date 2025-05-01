package services

import (
	"context"
	"sport-app-backend/helper"
	"sport-app-backend/models"
	"sport-app-backend/repositories"
)

type CategoryFieldService interface {
	CreateCategoryField(ctx context.Context, input *models.CategoryFieldRequest) (*models.CategoryFieldResponse, helper.Error)
	GetAllCategoryField(ctx context.Context) ([]models.CategoryFieldResponse, helper.Error)
	GetCategoryFieldByID(ctx context.Context, id string) (*models.CategoryFieldResponse, helper.Error)
	UpdateCategoryField(ctx context.Context, id string, categoryField *models.CategoryFieldRequest) (*models.CategoryFieldResponse, helper.Error)
	DeleteCategoryField(ctx context.Context, id string) (*models.CategoryFieldResponse, helper.Error)
}

type categoryFieldService struct {
	categoryFieldRepository repositories.CategoryFieldRepository
}

func NewCategoryFieldService(categoryFieldRepository repositories.CategoryFieldRepository) *categoryFieldService {
	return &categoryFieldService{
		categoryFieldRepository: categoryFieldRepository,
	}
}

func (cs *categoryFieldService) CreateCategoryField(ctx context.Context, input *models.CategoryFieldRequest) (*models.CategoryFieldResponse, helper.Error) {
	categoryField := input.NewCategoryField()
	categoryFieldResponse, err := cs.categoryFieldRepository.CreateCategoryField(ctx, &categoryField)
	if err != nil {
		return nil, err
	}
	return cs.mapCategoryFieldsResponse(categoryFieldResponse), nil
}

func (cs *categoryFieldService) GetAllCategoryField(ctx context.Context) ([]models.CategoryFieldResponse, helper.Error) {
	categoryFields, err := cs.categoryFieldRepository.GetAllCategoryField(ctx)
	if err != nil {
		return nil, err
	}
	var categoryFieldsResponse []models.CategoryFieldResponse
	for _, categoryField := range categoryFields {
		categoryFieldsResponse = append(categoryFieldsResponse, *cs.mapCategoryFieldsResponse(&categoryField))
	}
	return categoryFieldsResponse, nil
}

func (cs *categoryFieldService) GetCategoryFieldByID(ctx context.Context, id string) (*models.CategoryFieldResponse, helper.Error) {
	categoryField, err := cs.categoryFieldRepository.GetCategoryFieldByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return cs.mapCategoryFieldsResponse(categoryField), nil
}

func (cs *categoryFieldService) UpdateCategoryField(ctx context.Context, id string, input *models.CategoryFieldRequest) (*models.CategoryFieldResponse, helper.Error) {
	categoryFieldResponse, err := cs.categoryFieldRepository.UpdateCategoryField(ctx, id, input.NewCategoryField())
	if err != nil {
		return nil, err
	}
	return cs.mapCategoryFieldsResponse(categoryFieldResponse), nil
}

func (cs *categoryFieldService) DeleteCategoryField(ctx context.Context, id string) (*models.CategoryFieldResponse, helper.Error) {
	categoryFieldResponse, err := cs.categoryFieldRepository.DeleteCategoryField(ctx, id)
	if err != nil {
		return nil, err
	}
	return cs.mapCategoryFieldsResponse(categoryFieldResponse), nil
}

func (cs *categoryFieldService) mapCategoryFieldsResponse(categoryFields *models.CategoryField) *models.CategoryFieldResponse {
	return &models.CategoryFieldResponse{
		ID:           categoryFields.ID,
		CategoryName: categoryFields.CategoryName,
		CreatedAt:    categoryFields.CreatedAt,
		UpdatedAt:    categoryFields.UpdatedAt,
	}
}