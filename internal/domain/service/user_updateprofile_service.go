package service

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/dto"
	"github.com/google/uuid"
)

func (s *UserService) UpdateProfile(id string, updatedUser *dto.UpdateProfileDTO) (*entity.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	converted := updatedUser.UpdateProfileConvertToUser(updatedUser)
	errRepo := s.repo.UpdateUserByID(uid, converted)
	if errRepo != nil {
		return nil, errRepo
	}
	return converted, nil
}
