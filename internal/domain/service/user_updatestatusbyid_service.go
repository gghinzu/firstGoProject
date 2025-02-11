package service

import "firstGoProject/internal/domain/enum"

func (s *UserService) UpdateUserStatusByID(id int, userStatus enum.UserStatus) error {
	_, err := s.GetUserByID(id)
	if err != nil {
		return err
	}
	return s.repo.UpdateUserStatusByID(id, userStatus)
}
