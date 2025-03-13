package service

import (
	"crypto/rand"
	"errors"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	"firstGoProject/internal/helper"
	"fmt"
	"math/big"
)

func GenerateVerificationCode() string {
	num, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	return fmt.Sprintf("%06d", num)
}

func SendVerificationEmail(email, code string) error {
	// TODO: SMTP integration
	fmt.Printf("Verification email sent to %s with code: %s\n", email, code)
	return nil
}

func (s *UserService) SignUp(newUser *dto.SignUpDTO) error {
	hash, err := helper.EncryptPassword(newUser.Password)
	if err != nil {
		return fmt.Errorf("failed to encrypt password: %v", err)
	}
	newUser.Password = string(hash)

	convertedUser := newUser.SignUpConvertToUser(newUser)
	if convertedUser == nil {
		return errors.New("failed to convert DTO to entity")
	}

	userRole, err := s.repo.GetUserRoleByRoleName("user")
	if err != nil {
		return err
	}
	convertedUser.RoleID = userRole.RoleId

	existingUser, _ := s.repo.GetUserByEmail(convertedUser.Email)
	if existingUser != nil {
		return errors.New("email is taken")
	}

	verificationCode := GenerateVerificationCode()
	convertedUser.VerificationCode = verificationCode
	convertedUser.Status = enum.Passive

	err = s.repo.SignUp(convertedUser)
	if err != nil {
		return err
	}

	err = SendVerificationEmail(convertedUser.Email, verificationCode)
	if err != nil {
		return err
	}

	return nil
}
