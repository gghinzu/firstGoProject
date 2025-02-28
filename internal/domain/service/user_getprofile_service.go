package service

import (
	"firstGoProject/internal/domain/entity"
	"github.com/google/uuid"
)

func (s *UserService) GetProfile(id string) (*entity.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.repo.GetUserByID(uid)
}
