package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(loginPass, storedUserPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(storedUserPass), []byte(loginPass))
}
