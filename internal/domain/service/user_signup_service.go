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

	converted := newUser.SignUpConvertToUser(newUser)
	if converted == nil {
		return errors.New("failed to convert DTO to entity")
	}

	//create as user by default
	userRole, errU := s.repo.GetUserRoleByRoleName("user")
	if errU != nil {
		return errU
	}
	converted.RoleID = userRole.RoleId

	return s.repo.SignUp(converted)
}
