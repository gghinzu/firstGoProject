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

	convertedUser := updatedUser.UpdateProfileConvertToUser(updatedUser)
	err = s.repo.UpdateUserByID(uid, convertedUser)
	if err != nil {
		return nil, err
	}
	return convertedUser, nil
}
