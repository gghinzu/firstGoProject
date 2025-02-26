package service

import (
	"errors"
	"firstGoProject/internal/domain/entity"
	"firstGoProject/pkg/postgre/token"
)

func (s *UserService) RefreshToken(email *string) (*entity.TokenUserDTO, error) {
	if email == nil {
		return nil, errors.New("email cannot be nil")
	}

	_, err := s.repo.GetUserByEmail(*email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	tokenUser, errToken := token.CreateRefreshToken(email)
	if errToken != nil {
		return nil, errToken
	}
	return tokenUser, nil
}
