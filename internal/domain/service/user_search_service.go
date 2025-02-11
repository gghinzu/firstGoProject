package service

import "firstGoProject/internal/domain/entity"

func (s *UserService) SearchUser(searchStr string) (*[]entity.User, error) {
	return s.repo.SearchUser(searchStr)
}
