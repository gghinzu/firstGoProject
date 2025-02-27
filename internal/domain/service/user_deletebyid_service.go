package service

import (
	"fmt"
)

func (s *UserService) DeleteUserByID(id string) error {
	user, err := s.GetUserByID(id)
	if err != nil {
		return fmt.Errorf("%v", err)
	} else {
		return s.repo.DeleteUserByID(user.ID)
	}
}
