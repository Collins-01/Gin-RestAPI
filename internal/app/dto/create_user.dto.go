package dto

type CreateUserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
}
