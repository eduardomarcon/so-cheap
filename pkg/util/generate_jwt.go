package util

import (
	"encoding/json"
	"so-cheap/internal/config"
	"so-cheap/internal/user/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateNewAccessToken(user entity.User) (string, error) {
	secret := config.GetJWT().SecretKey

	claims := jwt.MapClaims{}

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(config.GetJWT().ExpireMinutes)).Unix()

	userJson, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	claims["user"] = string(userJson)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
