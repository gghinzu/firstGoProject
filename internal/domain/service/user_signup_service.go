package service

import (
	"errors"
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/dto"
	"firstGoProject/internal/helper"
)

func (s *UserService) SignUp(newUser *dto.SignUpDTO) error {
	hash, err := helper.EncryptPassword(newUser.Password)
	if err != nil {
		return err
	}

	newUser.Password = string(hash)

	converted := SignUpConvertToUser(newUser)
	if converted == nil {
		return errors.New("failed to convert DTO to entity")
	}

	return s.repo.SignUp(converted)
}

func SignUpConvertToUser(dto *dto.SignUpDTO) *entity.User {
	user := &entity.User{
		Email:     dto.Email,
		Password:  dto.Password,
		Name:      dto.Name,
		Surname:   dto.Surname,
		Age:       dto.Age,
		Gender:    dto.Gender,
		Education: dto.Education,
	}
	return user
}
