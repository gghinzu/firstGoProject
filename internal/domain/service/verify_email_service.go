package service

import (
	"errors"
	"firstGoProject/internal/domain/enum"
	error2 "firstGoProject/internal/error"
)

func (s *UserService) VerifyEmail(email, code string) error {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return err
	}

	if user.VerificationCode != &code {
		return errors.New(error2.InvalidInput)
	}

	user.Status = enum.Active
	user.VerificationCode = nil

	return s.repo.UpdateUserByID(user.ID, user)
}
