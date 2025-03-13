package service

import (
	"firstGoProject/internal/dto"
	"firstGoProject/internal/helper"
	"firstGoProject/pkg/jwt"
)

func (s *UserService) Login(user dto.LoginDTO) (*dto.TokenDTO, error) {
	storedUser, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	err = helper.ComparePassword(user.Password, storedUser.Password)
	if err != nil {
		return nil, err
	}

	userWithRole, role, err := s.GetUserWithRole(storedUser.ID)
	if err != nil {
		return nil, err
	}

	tokenUser, err := jwt.CreateAccessToken(userWithRole, role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.CreateRefreshToken(userWithRole, role)
	if err != nil {
		return nil, err
	}

	tokenUser.RefreshToken = refreshToken

	return tokenUser, nil
}
