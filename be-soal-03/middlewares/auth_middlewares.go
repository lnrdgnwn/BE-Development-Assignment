package middlewares

import (
	"strings"

	"be-soal-03/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if auth == "" {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}

	token := strings.Replace(auth, "Bearer ", "", 1)

	claims, err := utils.ParseToken(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
	}

	c.Locals("user_id", claims["user_id"])
	c.Locals("role", claims["role"])

	return c.Next()
}
