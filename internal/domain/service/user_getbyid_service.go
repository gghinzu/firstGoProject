package service

import (
	"firstGoProject/internal/domain/entity"
	"github.com/google/uuid"
)

// GetUserByID gets specified user with the given id using an instance of UserService
// (implementation of the interface UserServicePort)
func (s *UserService) GetUserByID(id string) (*entity.User, error) {
	uid := uuid.MustParse(id)
	return s.repo.GetUserByID(uid)
}
