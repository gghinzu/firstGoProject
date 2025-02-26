package token

import (
	"errors"
	"firstGoProject/internal/domain/entity"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func CreateRefreshToken(email *string) (*entity.TokenUserDTO, error) {
	if email == nil {
		return nil, errors.New("email is required")
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24 * 180).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &entity.TokenUserDTO{
		Email:        email,
		RefreshToken: &refreshTokenString,
	}, nil
}
