package service

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
)

func (s *UserService) GetUserWithRole(id string) (*entity.User, enum.UserRole, error) {
	user, err := s.repo.GetUserByID(id)

	if err != nil {
		return nil, "", err
	}

	role, err := s.repo.GetRoleByUserInfo(user)

	if err != nil {
		return nil, "", err
	}

	return user, role, nil
}
