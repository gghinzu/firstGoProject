package service

import (
	"firstGoProject/internal/domain/entity"
)

// GetAllUsers returns all the users in the database
// uses the instance of UserService (it connects the service with the repo)
func (s *UserService) GetAllUsers() *[]entity.User {
	return s.repo.GetAllUsers()
}
