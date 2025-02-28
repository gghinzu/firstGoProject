package service

import (
	"firstGoProject/internal/dto"
	"firstGoProject/pkg/jwt"
)

func (s *UserService) RefreshToken() (*dto.TokenUserDTO, error) {
	tokenUser, errToken := jwt.CreateRefreshToken()
	/*	if uid == uuid.Nil {
		return nil, errors.New("invalid id")
	}*/
	if errToken != nil {
		return nil, errToken
	}
	return tokenUser, nil
}
