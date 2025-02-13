package service

import (
	"firstGoProject/internal/domain/entity"
)

func (s *UserService) SearchUser(info entity.SearchUserDTO) (*[]entity.User, error) {
	return s.repo.SearchUser(info)
}
