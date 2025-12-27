package middlewares

import "github.com/gofiber/fiber/v2"

func AdminMiddleware(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	}

	if role != "ADMIN" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"message": "Forbidden: Admin access only",
		})
	}

	return c.Next()
}
