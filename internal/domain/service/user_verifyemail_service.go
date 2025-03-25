package service

import (
	"errors"
	"firstGoProject/internal/domain/enum"
	"time"
)

func (s *UserService) VerifyEmail(email, code string) error {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return err
	}

	if user.VerificationCode == nil {
		return errors.New("no verification code found")
	}

	if *user.VerificationCode != code {
		return errors.New("invalid verification code")
	}

	if user.VerificationCodeExpiry != nil && time.Now().After(*user.VerificationCodeExpiry) {
		return errors.New("verification code is expired")
	}

	user.Status = enum.Active
	user.VerificationCode = nil
	user.VerificationCodeExpiry = nil

	return s.repo.UpdateUserByID(user.ID, user)
}
