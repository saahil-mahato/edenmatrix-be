package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saahil-mahato/edenmatrix-be/src/models"
	"github.com/saahil-mahato/edenmatrix-be/src/services"
)

type AuthController struct {
	Service services.AuthService
}

func (ac *AuthController) Register(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	return ac.Service.CreateUser(c, user)
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	authPayload := new(models.AuthPayload)

	if err := c.BodyParser(authPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	return ac.Service.LoginUser(c, authPayload)

}
