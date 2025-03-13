package service

import (
	"errors"
	"firstGoProject/internal/domain/enum"
)

func (s *UserService) VerifyEmail(email, code string) error {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	if user.VerificationCode != code {
		return errors.New("invalid verification code")
	}

	user.Status = enum.Active
	user.VerificationCode = ""

	return s.repo.UpdateUserByID(user.ID, user)
}
