package service

import (
	"errors"
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"fmt"
)

// UpdateUserByID gets an id and UserDTO as parameters, converts the DTO into the entity and sends it to database
// uses the instance of UserService (it connects the service with the repo)
func (s *UserService) UpdateUserByID(id int, updatedUser *entity.UpdateUserDTO) error {
	user, err := s.GetUserByID(id)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	converted := UpdateConvertToUser(updatedUser)
	if converted == nil {
		return errors.New("dto to entity conversion failed")
	}
	if user.Status != enum.Deleted {
		return s.repo.UpdateUserByID(user.ID, converted)
	} else {
		return errors.New("user cannot be updated because it is soft deleted")
	}
}

func UpdateConvertToUser(dto *entity.UpdateUserDTO) *entity.User {
	user := &entity.User{
		Name:      dto.Name,
		Surname:   dto.Surname,
		Age:       dto.Age,
		Gender:    dto.Gender,
		Education: dto.Education,
	}
	return user
}
