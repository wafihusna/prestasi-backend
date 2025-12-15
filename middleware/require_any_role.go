package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RequireAnyRole(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := strings.ToLower(c.Locals("role").(string))

		for _, r := range roles {
			if userRole == strings.ToLower(r) {
				return c.Next()
			}
		}

		return c.Status(403).JSON(fiber.Map{
			"message": "forbidden",
		})
	}
}