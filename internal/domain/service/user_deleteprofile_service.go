package service

import (
	"github.com/google/uuid"
)

func (s *UserService) DeleteProfile(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	errRepo := s.repo.DeleteUserByID(uid)
	if errRepo != nil {
		return errRepo
	}
	return nil
}
