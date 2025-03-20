package service

import (
	"crypto/rand"
	"errors"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	error2 "firstGoProject/internal/error"
	"firstGoProject/internal/helper"
	"firstGoProject/pkg/config"
	"fmt"
	"math/big"
	"net/smtp"
)

func (s *UserService) Register(newUser *dto.RegisterDTO) error {
	hash, err := helper.EncryptPassword(newUser.Password)

	if err != nil {
		return err
	}

	newUser.Password = string(hash)

	user := newUser.RegisterConvertToUser(newUser)
	if user == nil {
		return error2.ConversionError
	}

	userRole, err := s.repo.GetUserRoleByRoleName("user")
	if err != nil {
		return err
	}

	user.RoleID = userRole.RoleId

	existingUser, _ := s.repo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email is taken")
	}

	verificationCode := GenerateVerificationCode()
	user.VerificationCode = &verificationCode
	user.Status = enum.Passive

	err = s.repo.Register(user)
	if err != nil {
		return err
	}

	err = SendVerificationEmail(user.Email, verificationCode)
	if err != nil {
		return err
	}

	return nil
}

func GenerateVerificationCode() string {
	num, _ := rand.Int(rand.Reader, big.NewInt(1000000))

	return fmt.Sprintf("%06d", num)
}

func SendVerificationEmail(email, code string) error {
	configure, err := config.LoadConfig()

	if err != nil {
		return err
	}

	smtpHost := configure.SMTPHost
	smtpPort := configure.SMTPPort
	senderEmail := configure.SMTPSenderEmail
	senderPassword := configure.SMTPSenderPassword

	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
	to := []string{email}
	msg := []byte(fmt.Sprintf("Subject: Verification Code\r\n\r\nHi,\n\nYour verification code: %s\n\nThanks!", code))

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, to, msg)
	if err != nil {
		return err
	}

	fmt.Printf("Verification email sent to %s with code: %s\n", email, code)

	return nil
}
