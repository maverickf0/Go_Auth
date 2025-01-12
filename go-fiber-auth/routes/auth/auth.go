package auth

import (
	"go-fiber-auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/auth/register", controllers.Register)
	app.Post("/auth/login", controllers.Login)
}
