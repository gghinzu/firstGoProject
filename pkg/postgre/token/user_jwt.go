package token

import (
	"firstGoProject/internal/domain/entity"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func CreateToken(user *entity.User) (*entity.TokenUserDTO, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID.String(),
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := []byte(GetSecretKey())
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &entity.TokenUserDTO{
		Email: &user.Email,
		Token: &tokenString,
	}, nil
}

func GetSecretKey() string {
	return "secret-key"
}
