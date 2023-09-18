package user

import (
	"rest_api/internal/app/dto"
	"rest_api/internal/app/model"
)

// UserRepository defines the methods for user-related operations.
type UserRepository interface {
	CreateUser(user *dto.CreateUserDTO) error
	GetUserByID(id int) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
}
