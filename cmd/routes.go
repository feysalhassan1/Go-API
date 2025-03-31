package main

import (
	"github.com/feysalhassan1/Go-API/handler"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handler.Home)

}
