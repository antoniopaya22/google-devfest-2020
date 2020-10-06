package main

import "github.com/gofiber/fiber/v2"

func SetRoutes(app *fiber.App) {
	api := app.Group("/api")

	user := api.Group("/users")
	user.Get("/:id", GetUser)
	user.Get("/", GetUsers)
	user.Post("/", CreateUser)
}