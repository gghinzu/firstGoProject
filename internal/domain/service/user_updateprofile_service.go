package service

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/dto"
)

func (s *UserService) UpdateProfile(id string, updatedUser *dto.UpdateProfileDTO) (*entity.User, error) {
	user := updatedUser.UpdateProfileConvertToUser(updatedUser)
	err := s.repo.UpdateUserByID(id, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
