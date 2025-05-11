package utils

import (
	"time"
	"user-profile-service/internal/config"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(config.GetValue("JWT_SECRET"))

func GenerateJWT(userId string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Minute * 15).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
