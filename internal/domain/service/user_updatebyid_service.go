package service

import (
	"errors"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	"fmt"
)

func (s *UserService) UpdateUserByID(id string, updatedUser *dto.UpdateDTO) error {
	userUpdate, err := s.GetUserByID(id)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	convertedUser := updatedUser.UpdateConvertToUser(updatedUser)
	if convertedUser == nil {
		return errors.New("dto to entity conversion failed")
	}
	if userUpdate.Status != enum.Deleted { //&& (user.Role.Name == enum.Admin || user.ID == convertedUser.ID)
		return s.repo.UpdateUserByID(userUpdate.ID, convertedUser)
	} else {
		return errors.New("user is deactivated and cannot be updated")
	}
}
