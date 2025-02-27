package helper

import (
	a "firstGoProject/pkg/jwt"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.GetSecretKey()), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
