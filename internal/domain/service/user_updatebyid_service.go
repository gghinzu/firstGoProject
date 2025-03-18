package service

import (
	"errors"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	error2 "firstGoProject/internal/error"
	"fmt"
)

func (s *UserService) UpdateUserByID(id string, updatedUser *dto.UpdateDTO) error {
	userUpdate, err := s.GetUserByID(id)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	convertedUser := updatedUser.UpdateConvertToUser(updatedUser)

	if convertedUser == nil {
		return error2.ConversionError
	}

	if userUpdate.Status != enum.Deleted {
		return s.repo.UpdateUserByID(userUpdate.ID, convertedUser)
	} else {
		return errors.New("user is deactivated and cannot be updated")
	}
}
