package service

import (
	"firstGoProject/internal/domain/entity"
)

func (s *UserService) CreateUser(newUser *entity.CreateUserDTO) error {
	converted, err := CreateConvertToUser(newUser)
	if err != nil {
		return err
	}
	return s.repo.CreateUser(converted)
}

func CreateConvertToUser(dto *entity.CreateUserDTO) (*entity.User, error) {
	user := &entity.User{
		Name:      dto.Name,
		Surname:   dto.Surname,
		Age:       dto.Age,
		Gender:    dto.Gender,
		Education: dto.Education,
	}
	return user, nil
}
