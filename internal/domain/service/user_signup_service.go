package service

import (
	"errors"
	"firstGoProject/internal/dto"
	"firstGoProject/internal/helper"
)

func (s *UserService) SignUp(newUser *dto.SignUpDTO) error {
	hash, err := helper.EncryptPassword(newUser.Password)
	if err != nil {
		return err
	}

	newUser.Password = string(hash)

	convertedUser := newUser.SignUpConvertToUser(newUser)
	if convertedUser == nil {
		return errors.New("failed to convert DTO to entity")
	}

	//create as user by default
	userRole, err := s.repo.GetUserRoleByRoleName("user")
	if err != nil {
		return err
	}
	convertedUser.RoleID = userRole.RoleId

	user, _ := s.repo.GetUserByEmail(convertedUser.Email)

	if user != nil {
		return errors.New("email is taken")
	}

	return s.repo.SignUp(convertedUser)
}
