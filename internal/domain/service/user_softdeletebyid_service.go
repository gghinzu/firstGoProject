package service

import (
	"errors"
	"firstGoProject/internal/domain/enum"
	"fmt"
)

func (s *UserService) SoftDeleteUserByID(id int) error {
	user, err := s.GetUserByID(id)
	if err != nil {
		return fmt.Errorf("%v", err)
	} else {
		if user.Deleted != true {
			_ = s.repo.UpdateUserStatusByID(user.ID, enum.Passive)
			return s.repo.SoftDeleteUserByID(user.ID)
		} else {
			return errors.New("user is already soft deleted")
		}
	}
}
