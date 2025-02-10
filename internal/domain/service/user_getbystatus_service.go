package service

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
)

func (s *UserService) GetUsersByStatus(status enum.UserStatus) (*[]entity.User, error) {
	return s.repo.GetUsersByStatus(status)
}
