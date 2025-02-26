package token

import (
	"firstGoProject/internal/domain/entity"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func CreateToken(user *entity.User) (*entity.TokenUserDTO, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID.String(),
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	})

	/*	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    user.ID.String(),
			"email": user.Email,
			"exp":   time.Now().Add(time.Hour * 24 * 180).Unix(),
		})
	*/
	secretKey := []byte(GetSecretKey())

	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	/*	refreshTokenString, err := refreshToken.SignedString(secretKey)
		if err != nil {
			return nil, err
		}*/

	return &entity.TokenUserDTO{
		Email: &user.Email,
		Token: &accessTokenString,
		/*		RefreshToken: &refreshTokenString,*/
	}, nil
}

func GetSecretKey() string {
	return "wdjoasnwqdw2911204sd"
}
