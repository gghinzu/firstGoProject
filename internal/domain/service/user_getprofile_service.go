package service

import (
	"firstGoProject/internal/domain/entity"
)

func (s *UserService) GetProfile(id string) (*entity.User, error) {
	return s.repo.GetUserByID(id)
}
