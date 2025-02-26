package service

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/helper"
	"firstGoProject/pkg/postgre/token"
)

func (s *UserService) Login(user entity.LoginDTO) (*entity.TokenUserDTO, error) {
	storedUser, err := s.repo.GetUserByEmail(*user.Email)
	if err != nil {
		return nil, err
	}

	errCompare := helper.ComparePassword(*user.Password, storedUser.Password)
	if errCompare != nil {
		return nil, errCompare
	}

	tokenUser, errToken := token.CreateAccessToken(storedUser)
	if errToken != nil {
		return nil, errToken
	}

	/*storedUser.Token = *tokenUser.Token
	err = s.repo.UpdateUserToken(storedUser.ID, storedUser.Token)
	if err != nil {
		return nil, err
	}*/

	return tokenUser, nil
}
