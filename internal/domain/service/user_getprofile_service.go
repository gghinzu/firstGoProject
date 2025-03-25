package service

import (
	"firstGoProject/internal/dto"
)

func (s *UserService) GetProfile(id string) (*dto.UserResponseDTO, error) {
	return s.repo.GetUserByID(id)
}
