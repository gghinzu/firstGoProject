package service

import (
	"firstGoProject/internal/domain/entity"
)

func (s *UserService) GetUserByID(id string) (*entity.User, error) {
	return s.repo.GetUserByID(id)
}
