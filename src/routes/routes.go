package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saahil-mahato/edenmatrix-be/src/controllers"
	"github.com/saahil-mahato/edenmatrix-be/src/repositories"
)

func Setup(app *fiber.App) {
	authController := controllers.AuthController{Repo: repositories.AuthRepository{}}

	app.Post("/register", authController.Register)
	app.Post("/login", authController.Login)
}
