package routes

import (
	"go-fiber-auth/routes/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	auth.SetupAuthRoutes(app)
}
