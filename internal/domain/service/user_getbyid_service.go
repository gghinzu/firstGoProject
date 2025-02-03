package service

import (
	"firstGoProject/internal/domain/entity"
)

// GetUserByID gets specified user with the given id using an instance of UserService
// (implementation of the interface UserServicePort)
func (s *UserService) GetUserByID(id int) *entity.User {
	user := s.repo.GetUserByID(id)
	return user
}
