package main

import (
	"log"

	"loadboard/database"
	"loadboard/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	app := fiber.New()

	app.Use(cors.New()) // <-- ADD THIS BEFORE ROUTES

	routes.SetupRoutes(app)

log.Fatal(app.Listen(":8080"))
}
