package service

import (
	"errors"
	"firstGoProject/internal/domain/enum"
)

func (s *UserService) UpdateUserStatusByID(id int, userStatus enum.UserStatus) error {
	user, err := s.GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Status == enum.Active && userStatus == enum.Active {
		return errors.New("user is already active")
	} else if user.Status == enum.Passive && userStatus == enum.Passive {
		return errors.New("user is already passive")
	} else {
		if user.Deleted != true {
			return s.repo.UpdateUserStatusByID(user.ID, userStatus)
		} else {
			return errors.New("user cannot be updated because it is soft deleted")
		}
	}
}
