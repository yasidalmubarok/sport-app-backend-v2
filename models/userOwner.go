package models

import (
	"time"
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