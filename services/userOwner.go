package services

import (
	"context"
	"sport-app-backend/helper"
	"sport-app-backend/models"
	"sport-app-backend/repositories"
	"time"
)

type UserOwnerService interface {
	CreateUserOwner(ctx context.Context, userOwner *models.RegisterUserOwnerRequest) (*models.UserOwner, helper.Error)
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

	hashedPassword, err := helper.HashPassword(userOwnerPayLoad.Password)
	if err != nil {
		return nil, helper.NewInternalServerError(err.Error())
	}

	userOwner := userOwnerPayLoad.NewUserOwner()
	userOwner.Password = hashedPassword

	u, _ := uos.userOwnerRepository.GetUserOwnerByUsername(ctx, userOwner.Username)
	if &u.Username == &userOwnerPayLoad.Username {
		return nil, helper.NewConflictError("username already exists")
	}

	p, _ := uos.userOwnerRepository.GetUserOwnerByPhoneNumber(ctx, userOwner.PhoneNumber)
	if &p.PhoneNumber == &userOwnerPayLoad.PhoneNumber {
		return nil, helper.NewConflictError("phone number already exists")
	}

	newUserOwner, err := uos.userOwnerRepository.CreateUserOwner(ctx, &userOwner)
	if err != nil {
		return nil, helper.NewInternalServerError(err.Error())
	}

	return uos.mapRegisterOwnerResponse(newUserOwner), nil
}

func (uos *userOwnerService) mapRegisterOwnerResponse(userOwner *models.UserOwner) *models.RegisterOwnerResponse {
	return &models.RegisterOwnerResponse{
		ID:          userOwner.ID,
		Name:        userOwner.Name,
		Username:    userOwner.Username,
		PhoneNumber: userOwner.PhoneNumber,
		Role:        userOwner.Role,
		CreatedAt:   userOwner.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   userOwner.UpdatedAt.Format(time.RFC3339),
	}
}
