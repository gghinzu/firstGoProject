package service

import (
	"errors"
	"firstGoProject/internal/domain/entity"
)

func (s *UserService) UpdateUserByID(id int, updatedUser *entity.UserDTO) error {
	existingUser := s.repo.GetUserByID(id)

	if existingUser == nil {
		return errors.New("user not found")
	} else {
		if &updatedUser.Name != nil {
			existingUser.Name = updatedUser.Name
		}
		if &updatedUser.Surname != nil {
			existingUser.Surname = updatedUser.Surname
		}
		if updatedUser.Age > 0 && &updatedUser.Age != nil {
			existingUser.Age = updatedUser.Age
		}
		if &updatedUser.Gender != nil {
			existingUser.Gender = updatedUser.Gender
		}
		if &updatedUser.Education != nil {
			existingUser.Education = updatedUser.Education
		}

		return s.repo.UpdateUserByID(id, existingUser)
	}
}
