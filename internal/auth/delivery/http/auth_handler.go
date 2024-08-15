package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/thitiphongD/go-backend/internal/auth/usecase"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func (h *AuthHandler) SignIn(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	user, token, err := h.authUsecase.SignIn(data["email"], data["password"])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
	}
	c.Cookie(&cookie)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message":      "Login successful",
		"code":         200,
		"email":        user.Email,
		"name":         user.Name,
		"profileImage": user.Image,
		"role":         user.Role,
		"token":        cookie.Value,
	})
}
