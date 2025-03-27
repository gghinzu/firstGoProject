package validation

import (
	"firstGoProject/internal/domain/enum"
	"regexp"
	"time"
)

var (
	emailRegex    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	nameRegex     = regexp.MustCompile(`^[a-zA-Z\s]{2,255}$`)
	passwordRegex = regexp.MustCompile(`^.{6,255}$`)
)

func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func ValidatePassword(password string) bool {
	return passwordRegex.MatchString(password)
}

func ValidateName(name string) bool {
	return nameRegex.MatchString(name)
}

func ValidateSurname(surname string) bool {
	return nameRegex.MatchString(surname)
}

func ValidateAge(age time.Time) bool {
	now := time.Now()
	return age.Before(now) && now.Sub(age).Hours() >= 24*365
}

func ValidateGender(gender enum.UserGender) bool {
	return gender == enum.Male || gender == enum.Female || gender == enum.NotSpecified
}

func ValidateEducation(education enum.UserEducation) bool {
	return education >= enum.None && education <= enum.Doctorate
}

func ValidateStatus(status enum.UserStatus) bool {
	return status == enum.Active || status == enum.Passive || status == enum.Deleted
}

func ValidateRole(role enum.UserRole) bool {
	return role == enum.User || role == enum.Admin
}
