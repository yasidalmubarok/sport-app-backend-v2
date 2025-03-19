package repositories

import (
	"context"
	"sport-app-backend/helper"
	"sport-app-backend/models"

	"gorm.io/gorm"
)

type UserOwnerRepository interface {
	CreateUserOwner(ctx context.Context, userOwner *models.UserOwner) (*models.UserOwner, helper.Error)
	GetUserOwnerByUsername(ctx context.Context, username string) (*models.UserOwner, helper.Error)
	GetUserOwnerByPhoneNumber(ctx context.Context, phoneNumber string) (*models.UserOwner, helper.Error)
}

type userOwnerRepository struct {
	db *gorm.DB
}

func NewUserOwnerRepository(db *gorm.DB) *userOwnerRepository {
	return &userOwnerRepository{db: db}
}

func (uor *userOwnerRepository) CreateUserOwner(ctx context.Context, userOwner *models.UserOwner) (*models.UserOwner, helper.Error) {
	err := uor.db.WithContext(ctx).Create(userOwner).Error
	if err != nil {
		return nil, helper.NewInternalServerError("Something went wrong")
	}
	return userOwner, nil
}

func (uor *userOwnerRepository) GetUserOwnerByUsername(ctx context.Context, username string) (*models.UserOwner, helper.Error) {
	var userOwner models.UserOwner
	err := uor.db.WithContext(ctx).Where("username = ?", username).First(&userOwner).Error
	if err != nil {
		return nil, helper.NewNotFoundError("username not found")
	}
	return &userOwner, nil
}

func (uor *userOwnerRepository) GetUserOwnerByPhoneNumber(ctx context.Context, phoneNumber string) (*models.UserOwner, helper.Error) {
	var userOwner models.UserOwner
	err := uor.db.WithContext(ctx).Where("phone_number = ?", phoneNumber).First(&userOwner).Error
	if err != nil {
		return nil, helper.NewNotFoundError("phone number not found")
	}
	return &userOwner, nil
}
