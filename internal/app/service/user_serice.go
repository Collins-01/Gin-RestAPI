package service

import (
	"rest_api/internal/app/dto"
	"rest_api/internal/app/model"
	repository "rest_api/internal/app/repository/user"
)

// UserService represents a service for user-related operations.
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new instance of UserService with the provided UserRepository.
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(user *dto.CreateUserDTO) error {
	return s.userRepo.CreateUser(user)
}

// GetUserByID retrieves a user by their ID.
func (s *UserService) GetUserByID(id int) (*model.User, error) {
	return s.userRepo.GetUserByID(id)
}

// UpdateUser updates a user's information.
func (s *UserService) UpdateUser(user *model.User) error {
	return s.userRepo.UpdateUser(user)
}

// DeleteUser deletes a user by their ID.
func (s *UserService) DeleteUser(id int) error {
	return s.userRepo.DeleteUser(id)
}
