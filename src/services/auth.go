package services

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/saahil-mahato/edenmatrix-be/src/models"
	"github.com/saahil-mahato/edenmatrix-be/src/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo repositories.AuthRepository
}

func (s *AuthService) CreateUser(c *fiber.Ctx, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
	}
	user.Password = string(hashedPassword)

	if err := s.Repo.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": fmt.Sprintf("Successfully created user %s", user.Username)})
}

func (s *AuthService) LoginUser(c *fiber.Ctx, authPayload *models.AuthPayload) error {
	user, err := s.Repo.FindUserByEmail(authPayload.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User is not registered"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authPayload.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return createCookie(c, user)
}

func createCookie(c *fiber.Ctx, user *models.User) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	var jwtSecret = []byte("your_secret_key")
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteStrictMode,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Login successful"})
}
