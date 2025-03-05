package service

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"github.com/google/uuid"
)

func (s *UserService) GetUserWithRole(userID string) (*entity.User, enum.UserRole, error) {
	uid := uuid.MustParse(userID)
	user, err := s.repo.GetUserByID(uid)
	if err != nil {
		return nil, "", err
	}
	role, err := s.repo.GetUserWithRole(user)
	if err != nil {
		return nil, "", err
	}
	return user, role, nil
}
