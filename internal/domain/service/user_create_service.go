package service

import (
	"errors"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	error2 "firstGoProject/internal/error"
)

func (s *UserService) CreateUser(newUser *dto.CreateDTO) error {
	convertedUser := newUser.CreateConvertToUser(newUser)
	if convertedUser == nil {
		return error2.ConversionError
	}

	roleName := newUser.Role

	validRoles := map[string]bool{
		string(enum.User):  true,
		string(enum.Admin): true,
	}

	if !validRoles[roleName] {
		return errors.New("invalid role")
	}

	userRole, err := s.repo.GetUserRoleByRoleName(roleName)
	if err != nil {
		return err
	}

	convertedUser.RoleID = userRole.RoleId

	return s.repo.CreateUser(convertedUser)
}
