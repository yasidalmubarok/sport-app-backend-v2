package repositories

import (
	"context"
	"errors"
	"fmt"
	"sport-app-backend/helper"
	"sport-app-backend/models"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserOwnerRepository interface {
	CreateUserOwner(ctx context.Context, userOwner *models.UserOwner) (*models.UserOwner, helper.Error)
	GetUserOwnerByEmail(ctx context.Context, email string) (*models.UserOwner, helper.Error)
	IsUsernameExists(ctx context.Context, username string) (bool, helper.Error)
	IsPhoneNumberExists(ctx context.Context, phoneNumber string) (bool, helper.Error)
	IsEmailExists(ctx context.Context, email string) (bool, helper.Error)
	GetUserByUsernameOrPhone(ctx context.Context, identifier string) (*models.UserOwner, helper.Error)
	UpdateUserOwnerPassword(ctx context.Context, email string, newPassword string) (*models.UserOwner, helper.Error)
	SaveResetOTP(ctx context.Context, email string, otp string) helper.Error
	GetResetOTP(ctx context.Context, email string) (string, helper.Error)
	DeleteResetOTP(ctx context.Context, email string) helper.Error
}

type userOwnerRepository struct {
	db    *gorm.DB
	redisClient *redis.Client
}

func NewUserOwnerRepository(db *gorm.DB, redisClient *redis.Client) *userOwnerRepository {
	return &userOwnerRepository{
		db:    db,
		redisClient: redisClient,
	}
}

func (uor *userOwnerRepository) CreateUserOwner(ctx context.Context, userOwner *models.UserOwner) (*models.UserOwner, helper.Error) {
	err := uor.db.WithContext(ctx).Create(userOwner).Error
	if err != nil {
		return nil, helper.NewInternalServerError("Something went wrong")
	}
	return userOwner, nil
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

func (uor *userOwnerRepository) UpdateUserOwnerPassword(ctx context.Context, email string, newPassword string) (*models.UserOwner, helper.Error) {
	var userOwner models.UserOwner

	// Cari user berdasarkan ID
	err := uor.db.WithContext(ctx).First(&userOwner, "id = ?", userOwner.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewNotFoundError("user not found")
		}
		return nil, helper.NewInternalServerError(err.Error())
	}

	// Update password
	userOwner.Password = newPassword
	err = uor.db.WithContext(ctx).Save(&userOwner).Error
	if err != nil {
		return nil, helper.NewInternalServerError("failed to update password")
	}

	return &userOwner, nil
}

func (uor *userOwnerRepository) SaveResetOTP(ctx context.Context, email string, otp string) helper.Error {
	key := fmt.Sprintf("reset_otp:%s", email)
	err := uor.redisClient.Set(ctx, key, otp, 5*time.Minute).Err()
	if err != nil {
		return helper.NewInternalServerError("failed to store OTP in Redis")
	}
	return nil
}

func (uor *userOwnerRepository) GetResetOTP(ctx context.Context, email string) (string, helper.Error) {
	key := fmt.Sprintf("reset_otp:%s", email)
	otp, err := uor.redisClient.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", helper.NewNotFoundError("OTP not found or expired")
	} else if err != nil {
		return "", helper.NewInternalServerError("failed to retrieve OTP from Redis")
	}
	return otp, nil
}

func (uor *userOwnerRepository) DeleteResetOTP(ctx context.Context, email string) helper.Error {
	key := fmt.Sprintf("reset_otp:%s", email)
	_, err := uor.redisClient.Del(ctx, key).Result()
	if err != nil {
		return helper.NewInternalServerError("failed to delete OTP from Redis")
	}
	return nil
}
