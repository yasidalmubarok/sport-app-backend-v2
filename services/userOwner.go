package services

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"sport-app-backend/config"
	"sport-app-backend/helper"
	"sport-app-backend/models"
	"sport-app-backend/repositories"
	"time"
)

type UserOwnerService interface {
	CreateUserOwner(ctx context.Context, userOwner *models.RegisterUserOwnerRequest) (*models.RegisterOwnerResponse, helper.Error)
	LoginUserOwner(ctx context.Context, userOwner *models.LoginUserOwnerRequest) (*models.LoginOwnerResponse, helper.Error)
	RequestResetPassword(ctx context.Context, email string) helper.Error
	ResetPassword(ctx context.Context, email, otp, newPassword string) helper.Error
}

type userOwnerService struct {
	userOwnerRepository repositories.UserOwnerRepository
}

func NewUserOwnerService(userOwnerRepository repositories.UserOwnerRepository) *userOwnerService {
	return &userOwnerService{userOwnerRepository: userOwnerRepository}
}

func (uos *userOwnerService) CreateUserOwner(ctx context.Context, userOwnerPayLoad *models.RegisterUserOwnerRequest) (*models.RegisterOwnerResponse, helper.Error) {
	err := helper.ValidateStruct(userOwnerPayLoad)
	if err != nil {
		return nil, err
	}

	if err := helper.ValidateUsername(userOwnerPayLoad.Username); err != nil {
		return nil, err
	}

	if err := helper.ValidateEmail(userOwnerPayLoad.Email); err != nil {
		return nil, err
	}

	if err := helper.ValidatePhoneNumber(userOwnerPayLoad.PhoneNumber); err != nil {
		return nil, err
	}

	hashedPassword, err := helper.HashPassword(userOwnerPayLoad.Password)
	if err != nil {
		return nil, err
	}

	userOwner := userOwnerPayLoad.NewUserOwner()
	userOwner.Password = hashedPassword

	// Periksa apakah username sudah ada
	existingUser, _ := uos.userOwnerRepository.IsUsernameExists(ctx, userOwner.Username)
	if existingUser {
		return nil, helper.NewConflictError("username already exists")
	}

	existingEmail, _ := uos.userOwnerRepository.IsEmailExists(ctx, userOwner.Email)
	if existingEmail {
		return nil, helper.NewConflictError("email already exists")
	}

	// Periksa apakah nomor telepon sudah ada
	existingPhone, _ := uos.userOwnerRepository.IsPhoneNumberExists(ctx, userOwner.PhoneNumber)
	if existingPhone {
		return nil, helper.NewConflictError("phone number already exists")
	}

	newUserOwner, err := uos.userOwnerRepository.CreateUserOwner(ctx, &userOwner)
	if err != nil {
		return nil, err
	}

	return uos.mapRegisterOwnerResponse(newUserOwner), nil
}

func (uos *userOwnerService) LoginUserOwner(ctx context.Context, userOwnerPayLoad *models.LoginUserOwnerRequest) (*models.LoginOwnerResponse, helper.Error) {
	err := helper.ValidateStruct(userOwnerPayLoad)
	if err != nil {
		return nil, helper.NewBadRequestError(err.Error())
	}

	owner, err := uos.userOwnerRepository.GetUserByUsernameOrPhone(ctx, userOwnerPayLoad.Identifier)
	if err != nil {
		return nil, err
	}

	if owner == nil {
		return nil, helper.NewNotFoundError("invalid username or phone number")
	}

	if !helper.ComparePassword(owner.Password, userOwnerPayLoad.Password) {
		return nil, helper.NewUnauthenticatedError("incorrect password")
	}

	token, err := helper.GenerateJWT(owner.ID, owner.Name, owner.Username, owner.PhoneNumber, owner.Role)
	if err != nil {
		return nil, helper.NewInternalServerError("failed to generate token")
	}

	return uos.mapLoginOwnerWithTokenResponse(owner, token), nil
}

func (uos *userOwnerService) RequestResetPassword(ctx context.Context, email string) helper.Error {
	// Cek apakah email terdaftar
	_, err := uos.userOwnerRepository.GetUserOwnerByEmail(ctx, email)
	if err != nil {
		return helper.NewNotFoundError("email not registered")
	}

	// Generate OTP 4 digit
	otp, genErr := rand.Int(rand.Reader, big.NewInt(10000)) // Rentang 0000 - 9999
	if genErr != nil {
		return helper.NewInternalServerError("failed to generate OTP")
	}

	// Simpan OTP ke Redis
	err = uos.userOwnerRepository.SaveResetOTP(ctx, email, fmt.Sprintf("%04d", otp))
	if err != nil {
		return err
	}

	emailBody := fmt.Sprintf("Your OTP is: %04d", otp.Int64())
	if err := config.SendEmail(email, "Reset Password OTP", emailBody); err != nil {
		return err
	}
	return nil
}

func (uos *userOwnerService) ResetPassword(ctx context.Context, email, otp, newPassword string) helper.Error {
	// Ambil OTP dari Redis
	storedOTP, err := uos.userOwnerRepository.GetResetOTP(ctx, email)
	if err != nil {
		return helper.NewUnathorizedError("invalid or expired OTP")
	}

	// Verifikasi OTP
	if storedOTP != otp {
		return helper.NewUnathorizedError("incorrect OTP")
	}

	// Hash password baru
	hashedPassword, err := helper.HashPassword(newPassword)
	if err != nil {
		return helper.NewInternalServerError("failed to hash password")
	}

	// Update password di database
	_, err = uos.userOwnerRepository.UpdateUserOwnerPassword(ctx, email, hashedPassword)
	if err != nil {
		return err
	}

	// Hapus OTP dari Redis setelah berhasil reset password
	uos.userOwnerRepository.DeleteResetOTP(ctx, email)

	return nil
}
func (uos *userOwnerService) mapLoginOwnerWithTokenResponse(userOwner *models.UserOwner, token string) *models.LoginOwnerResponse {
	return &models.LoginOwnerResponse{
		Username:    userOwner.Username,
		PhoneNumber: userOwner.PhoneNumber,
		Token:       token,
	}
}

func (uos *userOwnerService) mapRegisterOwnerResponse(userOwner *models.UserOwner) *models.RegisterOwnerResponse {
	return &models.RegisterOwnerResponse{
		ID:          userOwner.ID,
		Name:        userOwner.Name,
		Username:    userOwner.Username,
		Email:       userOwner.Email,
		PhoneNumber: userOwner.PhoneNumber,
		Role:        userOwner.Role,
		CreatedAt:   userOwner.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   userOwner.UpdatedAt.Format(time.RFC3339),
	}
}
