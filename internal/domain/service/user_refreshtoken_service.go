package service

import (
	"firstGoProject/internal/dto"
	"firstGoProject/pkg/jwt"
)

func (s *UserService) RefreshToken() (*dto.TokenUserDTO, error) {
	tokenUser, errToken := jwt.CreateRefreshToken()
	if errToken != nil {
		return nil, errToken
	}
	return tokenUser, nil
}
