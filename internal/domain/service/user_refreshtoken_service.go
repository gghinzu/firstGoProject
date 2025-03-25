package service

import (
	"firstGoProject/internal/dto"
	"firstGoProject/internal/helper"
	"firstGoProject/pkg/jwt"
	"fmt"
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

	if _, err := helper.VerifyToken(refreshToken, "refresh"); err != nil {
		return nil, fmt.Errorf("invalid refresh token: %v", err)
	}

	tokenUser.RefreshToken = refreshToken

	return tokenUser, nil
}
