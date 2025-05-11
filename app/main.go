package main

import (
	"user-profile-service/internal/config"
	"user-profile-service/internal/database"
	"user-profile-service/internal/redis"
	"user-profile-service/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	database.DBconnection()
	redis.RedisConnection()

	app := fiber.New()
	routes.UserProfileRoute(app)
	app.Listen(":8080")
}
