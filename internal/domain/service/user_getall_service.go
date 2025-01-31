package service

import (
	"firstGoProject/internal/domain/entity"
)

// GetAllUsers gets all users using an instance of UserService
// (implementation of the interface UserServicePort)
func (service *UserService) GetAllUsers() []entity.User {
	return service.repo.GetAllUsers()
}
