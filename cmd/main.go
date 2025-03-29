package main

import (
	"log"

	"github.com/feysalhassan1/Go-API/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)
	//
	log.Fatal(app.Listen(":3000"))
}
