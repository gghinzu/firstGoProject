package service

import (
	"firstGoProject/internal/dto"
)

func (s *UserService) FilterUser(info dto.FilterDTO) (*[]dto.UserResponseDTO, error) {
	return s.repo.FilterUser(info)
}
