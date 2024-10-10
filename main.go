package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lpernett/godotenv"
	"github.com/saahil-mahato/edenmatrix-be/src/database"
	"github.com/saahil-mahato/edenmatrix-be/src/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("FRONTEND_URL"),
		AllowHeaders:     "Authorization, Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8080")
}
