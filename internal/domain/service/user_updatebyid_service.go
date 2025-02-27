package service

import (
	"errors"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	"fmt"
)

func (s *UserService) UpdateUserByID(id string, updatedUser *dto.UpdateUserDTO) error {
	userUpdate, err := s.GetUserByID(id)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	converted := updatedUser.UpdateConvertToUser(updatedUser)
	if converted == nil {
		return errors.New("dto to entity conversion failed")
	}
	if userUpdate.Status != enum.Deleted { //&& (user.Role.Name == enum.Admin || user.ID == converted.ID)
		return s.repo.UpdateUserByID(userUpdate.ID, converted)
	} else {
		return errors.New("user cannot be updated")
	}
}
