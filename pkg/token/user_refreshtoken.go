package token

import (
	"firstGoProject/internal/dto"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func CreateRefreshToken() (*dto.TokenUserDTO, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24 * 180).Unix(),
	})

	secretKey := []byte(GetSecretKey())

	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &dto.TokenUserDTO{
		RefreshToken: refreshTokenString,
	}, nil
}
