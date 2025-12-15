package middleware

import (
	"github.com/gofiber/fiber/v2"
	"projectbase/app/service"
)

func RBACMiddleware(rbacService *service.RBACService, permission string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		role := c.Locals("role").(string)

		allowed, err := rbacService.CheckPermission(role, permission)
		if err != nil || !allowed {
			return c.Status(403).JSON(fiber.Map{
				"message": "access denied",
			})
		}

		return c.Next()
	}
}