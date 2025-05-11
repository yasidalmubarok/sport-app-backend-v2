package models

import (
	"time"

	"github.com/google/uuid"
)

type CategoryField struct {
	ID           string    `gorm:"type:uuid;primary_key"`
	CategoryName string    `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:now()"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:now()"`
}

type CategoryFieldRequest struct {
	Name string `json:"name" binding:"required"`
}
type CategoryFieldResponse struct {
	ID           string    `json:"id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (cfr *CategoryFieldRequest) NewCategoryField() CategoryField {
	id := uuid.New().String()
	rawCreatedAt := time.Now().Format(time.RFC3339)
	created, _ := time.Parse(time.RFC3339, rawCreatedAt)
	rawUpdatedAt := time.Now().Format(time.RFC3339)
	updated, _ := time.Parse(time.RFC3339, rawUpdatedAt)

	return CategoryField{
		ID:           id,
		CategoryName: cfr.Name,
		CreatedAt:    created,
		UpdatedAt:    updated,
	}
}

