package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saahil-mahato/edenmatrix-be/src/database"
	"github.com/saahil-mahato/edenmatrix-be/src/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8080")
}
