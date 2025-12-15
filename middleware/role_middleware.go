package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RequireRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := strings.ToLower(c.Locals("role").(string))

		if userRole != strings.ToLower(role) {
			return c.Status(403).JSON(fiber.Map{
				"message": "forbidden",
			})
		}
		return c.Next()
	}
}
