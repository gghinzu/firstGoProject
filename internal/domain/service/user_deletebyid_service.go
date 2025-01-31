package service

import (
	"errors"
)

func (service *UserService) DeleteUserByID(ID int) error {
	err := service.repo.DeleteUserByID(ID)
	if err != nil {
		return errors.New("user not found")
	}
	return nil
}
