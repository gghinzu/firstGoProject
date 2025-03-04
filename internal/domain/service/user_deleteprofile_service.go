package service

import (
	"github.com/google/uuid"
)

func (s *UserService) DeleteProfile(id string) error {
	uid, err := uuid.Parse(id)
	//TODO: uuid
	if err != nil {
		return err
	}
	err = s.repo.DeleteUserByID(uid)
	if err != nil {
		return err
	}
	return nil
}
