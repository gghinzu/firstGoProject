package helper

import (
	"errors"
	error2 "firstGoProject/internal/error"
	a "firstGoProject/pkg/jwt"
	"github.com/golang-jwt/jwt/v4"
)

type UserCustomClaims struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	Type    string `json:"type"`
	jwt.RegisteredClaims
}

func VerifyToken(tokenString, expectedType string) (UserCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserCustomClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(a.GetSecretKey()), nil
	})

	if err != nil {
		return UserCustomClaims{}, err
	}

	customClaims, ok := token.Claims.(*UserCustomClaims)
	if !ok || !token.Valid {
		return UserCustomClaims{}, errors.New(error2.NotAuthorized)
	}

	if customClaims.Type != expectedType {
		return UserCustomClaims{}, errors.New(error2.InvalidInput)
	}

	return *customClaims, nil
}
