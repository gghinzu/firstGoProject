package service

import (
	"firstGoProject/internal/dto"
	"firstGoProject/pkg/jwt"
)

func (s *UserService) RefreshToken(userID string) (*dto.TokenUserDTO, error) {
	user, err := s.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	tokenUser, err := jwt.CreateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.CreateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}
	tokenUser.RefreshToken = refreshToken
	return tokenUser, nil
}
