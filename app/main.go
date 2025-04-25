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
		request := models.UserSignupRequest{}
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
		result := repositoies.UserSignup(&userProfile)
		if result != nil {
			return c.Status(fiber.StatusInternalServerError).JSON("error")
		}

		return c.Status(fiber.StatusCreated).JSON("success")
	})

	app.Post("/signin", func(c *fiber.Ctx) error {
		request := models.UserSigninRequest{}
		err := c.BodyParser(&request)
		if err != nil {
			return err
		}
		result, err := repositoies.UserSingin(request.Username, request.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}

		token, err := utils.GenerateJWT(result.UserId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}

		return c.Status(fiber.StatusFound).JSON(
			models.UserSigninResponse{
				Token: token,
			})
	})

	app.Get("/user/profile", func(c *fiber.Ctx) error {
		userId := ""

		//Get redis before

		result, err := repositoies.UserProfileGet(userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}

		return c.Status(fiber.StatusFound).JSON(
			models.UserProfileResponse{
				Username:     result.Username,
				Email:        result.Email,
				MobileNumber: result.MobileNumber,
			})
	})

	app.Listen(":8080")
}
