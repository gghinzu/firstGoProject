package service

import (
	"errors"
	"firstGoProject/internal/domain/enum"
)

func (s *UserService) UpdateUserStatusByID(id string, userStatus enum.UserStatus) error {
	user, err := s.GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Status == userStatus {
		return errors.New("user is already in that status")
	}
	return s.repo.UpdateUserStatusByID(user.ID, userStatus)
}
