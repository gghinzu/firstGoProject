package service

import (
	"firstGoProject/internal/domain/entity"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetAllUsers gets all users using an instance of UserService
// (implementation of the interface UserServicePort)
func (s *UserService) GetAllUsers() *[]entity.User {
	return s.repo.GetAllUsers()
}
