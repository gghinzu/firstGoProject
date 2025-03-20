package service

import (
	"errors"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	error2 "firstGoProject/internal/error"
)

func (s *UserService) CreateUser(newUser *dto.CreateDTO) error {
	user := newUser.CreateConvertToUser(newUser)
	if user == nil {
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

	user.RoleID = userRole.RoleId

	return s.repo.CreateUser(user)
}
