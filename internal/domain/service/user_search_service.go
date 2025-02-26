package service

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/dto"
)

func (s *UserService) SearchUser(info dto.SearchUserDTO) (*[]entity.User, error) {
	return s.repo.SearchUser(info)
}
