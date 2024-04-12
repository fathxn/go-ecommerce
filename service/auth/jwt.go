package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"go-ecommerce/config"
	"strconv"
	"time"
)

func CreateJWT(secret []byte, UserId int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpire)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    strconv.Itoa(UserId),
		"expired_at": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
