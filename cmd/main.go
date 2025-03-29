package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello and Feysal Yesterday, what are you doing, How is your day, good?")
	})

	log.Fatal(app.Listen(":3000"))
}
