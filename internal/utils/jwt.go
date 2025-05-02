package utils

import (
	"fmt"
	"log"
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

func ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		log.Fatal("Error decoding token:", err)
		return "", err
	}

	var userId string
	var exp int64

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Print("Decoded JWT claims")
		for key, value := range claims {
			log.Printf("  %s: %v\n", key, value)
		}

		exp = claims["exp"].(int64)
		if exp < time.Now().Unix() {
			return "", fmt.Errorf("token expired")
		}

		userId = claims["userId"].(string)
	} else {
		return "", fmt.Errorf("Invalid token")
	}

	return userId, nil
}
