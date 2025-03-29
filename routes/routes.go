package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupAuthenticationRoutes(app *fiber.App) {
	AuthRoutes(app) // Setup authentication-related routes
	// Later we'll add: LoadRoutes(app)
}

func SetupRoutes(app *fiber.App) {
	AuthRoutes(app)
	LoadRoutes(app)
}
