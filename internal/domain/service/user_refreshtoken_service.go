package service

import (
	"firstGoProject/internal/dto"
	"firstGoProject/pkg/jwt"
)

func (s *UserService) RefreshToken(userID string) (*dto.TokenDTO, error) {
	userWithRole, role, err := s.GetUserWithRole(userID)

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
