package service

import (
	"firstGoProject/internal/dto"
	"firstGoProject/pkg/token"
)

func (s *UserService) RefreshToken() (*dto.TokenUserDTO, error) {
	tokenUser, errToken := token.CreateRefreshToken()
	if errToken != nil {
		return nil, errToken
	}
	return tokenUser, nil
}
