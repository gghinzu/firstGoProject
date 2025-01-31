package service

import (
	"errors"
	"firstGoProject/internal/domain/entity"
)

// GetUserByID gets specified user with the given id using an instance of UserService
// (implementation of the interface UserServicePort)
func (service *UserService) GetUserByID(id int) (*entity.User, error) {
	user, err := service.repo.GetUserByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
