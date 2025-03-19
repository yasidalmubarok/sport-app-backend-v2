package models

import (
	"sport-app-backend/dto"
	"time"

	"github.com/google/uuid"
)


var UserOwnerRole = "owner"

type UserOwner struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}




func (uor *dto.RegisterUserOwnerRequest) NewUserOwner() UserOwner {
	id := uuid.New().String()
	rawCreatedAt := time.Now().Format(time.RFC3339)
	created, _ := time.Parse(time.RFC3339, rawCreatedAt)
	rawUpdatedAt := time.Now().Format(time.RFC3339)
	updated, _ := time.Parse(time.RFC3339, rawUpdatedAt)

	return UserOwner{
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