package handlers

import (
	"log"
	"user-profile-service/internal/models"
	"user-profile-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SignupHandler(c *fiber.Ctx) error {
	request := models.UserSignupRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		log.Fatal("Error BodyParser")
		return c.Status(fiber.StatusInternalServerError).JSON("error")
	}

	err = service.SingupService(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("error")
	}

	return c.Status(fiber.StatusCreated).JSON("success")
}

func SigninHandler(c *fiber.Ctx) error {
	request := models.UserSigninRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		log.Fatal("Error BodyParser")
		return c.Status(fiber.StatusInternalServerError).JSON("error")
	}

	response, err := service.SinginService(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("error")
	}

	return c.Status(fiber.StatusFound).JSON(response)
}

func UserProfileHandler(c *fiber.Ctx) error {
	userId := c.Locals("userId")

	reponse, err := service.UserProfileService(userId.(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("error")
	}

	return c.Status(fiber.StatusFound).JSON(reponse)
}
