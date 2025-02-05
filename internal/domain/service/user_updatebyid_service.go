package service

import (
	"firstGoProject/internal/domain/entity"
)

func (s *UserService) UpdateUserByID(id int, updatedUser *entity.UpdateUserDTO) error {
	converted, err := UpdateConvertToUser(updatedUser)
	if err != nil {
		return err
	}
	var user, errID = s.GetUserByID(id)
	if errID != nil {
		return err
	}
	return s.repo.UpdateUserByID(user.ID, converted)
}

func UpdateConvertToUser(dto *entity.UpdateUserDTO) (*entity.User, error) {
	user := &entity.User{
		Name:      dto.Name,
		Surname:   dto.Surname,
		Age:       dto.Age,
		Gender:    dto.Gender,
		Education: dto.Education,
	}
	return user, nil
}
