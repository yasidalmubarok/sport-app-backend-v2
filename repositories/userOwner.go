package repositories

import (
	"context"
	"errors"
	"sport-app-backend/helper"
	"sport-app-backend/models"

	"gorm.io/gorm"
)

type UserOwnerRepository interface {
	CreateUserOwner(ctx context.Context, userOwner *models.UserOwner) (*models.UserOwner, helper.Error)
	GetUserOwnerByUsername(ctx context.Context, username string) (*models.UserOwner, helper.Error)
	GetUserOwnerByPhoneNumber(ctx context.Context, phoneNumber string) (*models.UserOwner, helper.Error)
	GetUserOwnerByEmail(ctx context.Context, email string) (*models.UserOwner, helper.Error)
	IsUsernameExists(ctx context.Context, username string) (bool, helper.Error)
	IsPhoneNumberExists(ctx context.Context, phoneNumber string) (bool, helper.Error)
	IsEmailExists(ctx context.Context, email string) (bool, helper.Error)
	GetUserByUsernameOrPhone(ctx context.Context, identifier string) (*models.UserOwner, helper.Error)
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

func (uor *userOwnerRepository) GetUserOwnerByEmail(ctx context.Context, email string) (*models.UserOwner, helper.Error) {
	var userOwner models.UserOwner
	err := uor.db.WithContext(ctx).Where("email = ?", email).First(&userOwner).Error
	if err != nil {
		return nil, helper.NewNotFoundError("email not found")
	}
	return &userOwner, nil
}

func (uor *userOwnerRepository) IsUsernameExists(ctx context.Context, username string) (bool, helper.Error) {
	var count int64
	err := uor.db.WithContext(ctx).Model(&models.UserOwner{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, helper.NewConflictError("username already exists")
	}
	return count > 0, nil
}

func (uor *userOwnerRepository) IsPhoneNumberExists(ctx context.Context, phoneNumber string) (bool, helper.Error) {
	var count int64
	err := uor.db.WithContext(ctx).Model(&models.UserOwner{}).Where("phone_number = ?", phoneNumber).Count(&count).Error
	if err != nil {
		return false, helper.NewConflictError("phone number already exists")
	}
	return count > 0, nil
}

func (uor *userOwnerRepository) IsEmailExists(ctx context.Context, email string) (bool, helper.Error) {
	var count int64
	err := uor.db.WithContext(ctx).Model(&models.UserOwner{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, helper.NewConflictError("email already exists")
	}
	return count > 0, nil
}

func (uor *userOwnerRepository) GetUserByUsernameOrPhone(ctx context.Context, identifier string) (*models.UserOwner, helper.Error) {
	var user models.UserOwner
	
	// Coba cari berdasarkan username
	err := uor.db.WithContext(ctx).
		Where("username = ?", identifier).
		First(&user).Error

	if err == nil {
		return &user, nil
	}

	// Jika tidak ditemukan, cari berdasarkan phone_number
	err = uor.db.WithContext(ctx).
		Where("phone_number = ?", identifier).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewNotFoundError("user not found")
		}
		return nil, helper.NewInternalServerError(err.Error())
	}

	return &user, nil
}