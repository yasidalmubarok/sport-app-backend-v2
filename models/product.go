package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        string    `gorm:"type:uuid;primary_key"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Category  string    `gorm:"type:varchar(255);not null"`
	PriceSell float64   `gorm:"type:float;not null"`
	PriceBuy  float64   `gorm:"type:float;not null"`
	Stock     int       `gorm:"type:int;not null"`
	Status    string    `gorm:"type:varchar(255);not null"`
	Image     string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now()"`
}

type CreateProductRequest struct {
	Name      string  `json:"name" binding:"required"`
	Category  string  `json:"category" binding:"required"`
	PriceSell float64 `json:"price_sell" binding:"required"`
	PriceBuy  float64 `json:"price_buy" binding:"required"`
	Stock     int     `json:"stock" binding:"required"`
	Status    string  `json:"status" binding:"required,oneof=aktif 'non aktif'"`
	Image     string  `json:"image" binding:"required"`
}

type CreateProductResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	PriceSell float64   `json:"price_sell"`
	PriceBuy  float64   `json:"price_buy"`
	Stock     int       `json:"stock"`
	Status    string    `json:"status"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Product) NewProduct() Product {
	id := uuid.New().String()
	rawCreatedAt := time.Now().Format(time.RFC3339)
	created, _ := time.Parse(time.RFC3339, rawCreatedAt)
	rawUpdatedAt := time.Now().Format(time.RFC3339)
	updated, _ := time.Parse(time.RFC3339, rawUpdatedAt)

	return Product{
		ID:        id,
		Name:      p.Name,
		Category:  p.Category,
		PriceSell: p.PriceSell,
		PriceBuy:  p.PriceBuy,
		Stock:     p.Stock,
		Status:    p.Status,
		Image:     p.Image,
		CreatedAt: created,
		UpdatedAt: updated,
	}
}
