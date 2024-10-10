package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/saahil-mahato/edenmatrix-be/src/controllers"
	"github.com/saahil-mahato/edenmatrix-be/src/repositories"
)

func Setup(app *fiber.App) {
	apiVersion := "/v1"
	authController := controllers.AuthController{Repo: repositories.AuthRepository{}}

	app.Post(fmt.Sprintf("%s/register", apiVersion), authController.Register)
	app.Post(fmt.Sprintf("%s/login", apiVersion), authController.Login)
}
