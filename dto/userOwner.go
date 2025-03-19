package dto

type RegisterUserOwnerRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=30"`
	Username    string `json:"username" validate:"required,min=3,max=30"`
	PhoneNumber string `json:"phone_number" validate:"required,phone"`
	Password    string `json:"password" validate:"required,min=6,max=30"`
}

type UserOwnerResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}