package middleware

import "github.com/gofiber/fiber/v2"

func RequirePermission(permission string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		role := c.Locals("role").(string)
		perms := rolePermissions[role]

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
