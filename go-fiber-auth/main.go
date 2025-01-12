package main

import (
	"go-fiber-auth/config"
	"go-fiber-auth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//Loading environment variables
	config.LoadEnv()
	config.ConnectDB()

	//Creating a new fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",      // Allow the frontend origin
		AllowMethods: "GET,POST,PUT,DELETE",        // Allowed methods
		AllowHeaders: "Content-Type,Authorization", // Allowed headers
	}))
	//Setting up routes
	routes.SetupRoutes(app)

	//Starting the server
	app.Listen(":3000")
}
