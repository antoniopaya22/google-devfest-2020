package main

import "github.com/gofiber/fiber/v2"

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := User{
		Id:       id,
		Username: "Test",
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Product found", "data": user})
}

func GetUsers(c *fiber.Ctx) error {
	var users []User
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Product found", "data": users})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Created user", "data": user})
}