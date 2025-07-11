package models

import (
	"time"

	"github.com/google/uuid"
)

type CategoryProduct struct {
	ID        string `gorm:"type:uuid;primary_key"`
	Name      string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now()"`
}

type CategoryProductRequest struct {
	Name string `json:"name" binding:"required"`
}

type CategoryProductResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryProducts struct {
	CategoryID string `json:"category_id"`
	Category   string `json:"category"`
	Products   []CreateProductResponse
}

func CategoryWithProducts(category *CategoryProductResponse, products []Product) CategoryProducts {
	var productList []CreateProductResponse
	for _, product := range products {
		productList = append(productList, *product.ToResponse())
	}

	data := CategoryProducts{
		CategoryID: category.ID,
		Category:   category.Name,
		Products:   productList,
	}
	return data
}

func (cpr *CategoryProductRequest) NewCategoryProduct() CategoryProduct {
	id := uuid.New().String()
	rawCreatedAt := time.Now().Format(time.RFC3339)
	created, _ := time.Parse(time.RFC3339, rawCreatedAt)
	rawUpdatedAt := time.Now().Format(time.RFC3339)
	updated, _ := time.Parse(time.RFC3339, rawUpdatedAt)

	return CategoryProduct{
		ID:        id,
		Name:      cpr.Name,
		CreatedAt: created,
		UpdatedAt: updated,
	}
}
