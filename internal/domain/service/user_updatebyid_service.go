package service

import (
	"errors"
	"firstGoProject/internal/domain/entity"
)

func (service *UserService) UpdateUserByID(id int, updatedUser entity.User) error {
	existingUser, err := service.repo.GetUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	if updatedUser.Name != "" {
		existingUser.Name = updatedUser.Name
	}
	if updatedUser.Surname != "" {
		existingUser.Surname = updatedUser.Surname
	}
	if updatedUser.Age > 0 {
		existingUser.Age = updatedUser.Age
	}
	if updatedUser.Gender != "" {
		existingUser.Gender = updatedUser.Gender
	}
	if updatedUser.Education != "" {
		existingUser.Education = updatedUser.Education
	}

	return service.repo.UpdateUserByID(id, *existingUser)
}
