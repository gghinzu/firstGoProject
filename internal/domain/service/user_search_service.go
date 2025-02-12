package service

import "firstGoProject/internal/domain/entity"

func (s *UserService) SearchUser(name, status, gender string) (*[]entity.User, error) {
	return s.repo.SearchUser(name, status, gender)
}
