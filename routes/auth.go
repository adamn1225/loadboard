package routes

import (
	"loadboard/controllers"

	"loadboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	auth.Get("/me", middleware.JWTProtected(), controllers.Me)

}
