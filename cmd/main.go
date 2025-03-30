package main

import (
	"time"
	"user-profile-service/internal/config"
	"user-profile-service/internal/database"
	"user-profile-service/internal/models"
	"user-profile-service/internal/repositoies"
	"user-profile-service/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	database.DBconnection()

	app := fiber.New()

	app.Post("/signup", func(c *fiber.Ctx) error {
		request := models.UsersignupRequest{}
		err := c.BodyParser(&request)
		if err != nil {
			return err
		}
		userProfile := models.UserProfile{
			UserId:       utils.GenerateUserId(),
			Username:     request.Username,
			Password:     request.Password,
			Email:        request.Email,
			MobileNumber: request.MobileNumber,
			CreateDate:   time.Now(),
			UpdateDate:   time.Now(),
		}
		result := repositoies.Usersignup(&userProfile)
		if result != nil {
			return c.Status(fiber.StatusInternalServerError).JSON("error")
		}
		return c.Status(fiber.StatusCreated).JSON("success")
	})

	app.Listen(":8080")
}
