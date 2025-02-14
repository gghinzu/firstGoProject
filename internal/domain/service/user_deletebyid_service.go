package service

import (
	"fmt"
)

// DeleteUserByID gets an id and deletes the user with this id, from database
// uses the instance of UserService (it connects the service with the repo)
func (s *UserService) DeleteUserByID(id string) error {
	user, err := s.GetUserByID(id)
	if err != nil {
		return fmt.Errorf("%v", err)
	} else {
		return s.repo.DeleteUserByID(user.ID)
	}
}
