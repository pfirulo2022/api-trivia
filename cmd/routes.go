package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/pfirulo2022/api-trivia/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)

}
