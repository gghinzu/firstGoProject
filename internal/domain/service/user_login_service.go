package service

import (
	"firstGoProject/internal/dto"
	"firstGoProject/internal/helper"
	"firstGoProject/pkg/jwt"
)

func (s *UserService) Login(user dto.LoginDTO) (*dto.TokenUserDTO, error) {
	storedUser, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	errCompare := helper.ComparePassword(user.Password, storedUser.Password)
	if errCompare != nil {
		return nil, errCompare
	}

	tokenUser, errToken := jwt.CreateAccessToken(storedUser)
	if errToken != nil {
		return nil, errToken
	}

	return tokenUser, nil
}
