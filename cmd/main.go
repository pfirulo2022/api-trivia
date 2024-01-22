package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/pfirulo2022/api-trivia/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
