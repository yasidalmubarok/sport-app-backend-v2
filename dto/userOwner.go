package dto

import (
	"sport-app-backend/models"
	"time"

	"github.com/google/uuid"
)

var UserOwnerRole = "owner"

type RegisterUserOwnerRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=30"`
	Username    string `json:"username" validate:"required,min=3,max=30"`
	PhoneNumber string `json:"phone_number" validate:"required,phone"`
	Password    string `json:"password" validate:"required,min=6,max=30"`
}

type LoginUserOwnerRequest struct {
	Username string `json:"username" validate:"required,min=3,max=30"`
	Password string `json:"password" validate:"required,min=6,max=30"`
}

type RegisterOwnerResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (uor *RegisterUserOwnerRequest) NewUserOwner() models.UserOwner {
	id := uuid.New().String()
	rawCreatedAt := time.Now().Format(time.RFC3339)
	created, _ := time.Parse(time.RFC3339, rawCreatedAt)
	rawUpdatedAt := time.Now().Format(time.RFC3339)
	updated, _ := time.Parse(time.RFC3339, rawUpdatedAt)

	return models.UserOwner{
		ID:          id,
		Name:        uor.Name,
		Username:    uor.Username,
		PhoneNumber: uor.PhoneNumber,
		Password:    uor.Password,
		Role:        UserOwnerRole,
		CreatedAt:   created,
		UpdatedAt:   updated,
	}
}
