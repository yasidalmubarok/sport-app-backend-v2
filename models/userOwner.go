package models

import (
	"time"

	"github.com/google/uuid"
)

type UserOwner struct {
	ID          string    `gorm:"primary_key"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Username    string    `gorm:"type:varchar(255);not null"`
	PhoneNumber string    `gorm:"type:varchar(255);not null"`
	Password    string    `gorm:"type:varchar(15);not null"`
	Role        string    `gorm:"type:varchar(10);default:'owner'"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:now()"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:now()"`

	_ struct{} `gorm:"index:idx_username_phone,unique"`
}

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

func (uor *RegisterUserOwnerRequest) NewUserOwner() UserOwner {
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


