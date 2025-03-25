package service

import (
	"firstGoProject/internal/dto"
)

func (s *UserService) GetUserByID(id string) (*dto.UserResponseDTO, error) {
	return s.repo.GetUserByID(id)
}
