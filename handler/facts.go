package handlers

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello and Feysal Yesterday, what are you doing, How is your day, good?")
}
