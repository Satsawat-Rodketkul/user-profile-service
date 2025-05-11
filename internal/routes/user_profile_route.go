package routes

import (
	"user-profile-service/internal/handlers"
	"user-profile-service/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserProfileRoute(app *fiber.App) {
	app.Post("/signup", handlers.SignupHandler)
	app.Post("/signin", handlers.SigninHandler)

	userGroup := app.Group("/user", middleware.ValidateToken)
	userGroup.Post("/profile", handlers.UserProfileHandler)
}
