package middleware

import "github.com/gofiber/fiber/v2"

func RequirePermission(permission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		perms, ok := c.Locals("permissions").([]string)
		if !ok {
			return c.Status(403).JSON(fiber.Map{
				"message": "permission denied",
			})
		}

		for _, p := range perms {
			if p == permission {
				return c.Next()
			}
		}

		return c.Status(403).JSON(fiber.Map{
			"message": "permission denied",
		})
	}
}
