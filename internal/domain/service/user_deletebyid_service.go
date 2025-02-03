package service

import (
	"errors"
)

// DeleteUserByID gets all users using an instance of UserService
// (implementation of the interface UserServicePort)
func (s *UserService) DeleteUserByID(ID int) error {
	err := s.repo.DeleteUserByID(ID)
	if err != nil {
		return errors.New("user not found")
	}
	return nil
}
