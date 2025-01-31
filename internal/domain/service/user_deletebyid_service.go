package service

import (
	"errors"
)

// DeleteUserByID gets all users using an instance of UserService
// (implementation of the interface UserServicePort)
func (service *UserService) DeleteUserByID(ID int) error {
	err := service.repo.DeleteUserByID(ID)
	if err != nil {
		return errors.New("user not found")
	}
	return nil
}
