package service

import (
	"errors"
	"firstGoProject/internal/dto"
)

func (s *UserService) CreateUser(newUser *dto.CreateUserDTO) error {
	converted := newUser.CreateConvertToUser(newUser)
	if converted == nil {
		return errors.New("dto to entity conversion failed")
	}

	roleName := newUser.Role

	validRoles := map[string]bool{
		"user":  true,
		"admin": true,
	}

	if !validRoles[roleName] {
		return errors.New("role name is invalid")
	}

	userRole, err := s.repo.GetUserRoleByRoleName(roleName)
	if err != nil {
		return err
	}

	converted.RoleID = userRole.RoleId

	return s.repo.CreateUser(converted)
}
