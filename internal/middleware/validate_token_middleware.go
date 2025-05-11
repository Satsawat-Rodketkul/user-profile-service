package middleware

import (
	"fmt"
	"log"
	"user-profile-service/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(config.GetValue("JWT_SECRET"))

func ValidateToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return c.Status(fiber.StatusUnauthorized).JSON("Invalid token or missing")
	}

	tokenString := authHeader[7:]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON("Invalid or expired token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		log.Print("Decoded JWT claims")
		for key, value := range claims {
			log.Printf("  %s: %v\n", key, value)
		}
		c.Locals("userId", claims["userId"].(string))
	}

	return c.Next()
}
