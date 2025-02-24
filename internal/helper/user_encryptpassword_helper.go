package helper

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) ([]byte, error) {
	hash, errCreation := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if errCreation != nil {
		return nil, errCreation
	}
	return hash, nil
}
