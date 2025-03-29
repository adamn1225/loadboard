package routes

import (
	"loadboard/controllers"
	"loadboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(app *fiber.App) {
	load := app.Group("/loads", middleware.JWTProtected())
	load.Post("/", controllers.CreateLoad)
	load.Get("/", controllers.GetAllLoads)
	load.Post("/:id/claim", controllers.ClaimLoad)
	load.Post("/:id/deliver", controllers.MarkDelivered)
}
