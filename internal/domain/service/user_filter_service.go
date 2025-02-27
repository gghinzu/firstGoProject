package service

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/dto"
)

func (s *UserService) FilterUser(info dto.FilterDTO) (*[]entity.User, error) {
	return s.repo.FilterUser(info)
}
