package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"projectbase/utils"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "missing token",
			})
		}

		token := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": "invalid token",
			})
		}

		c.Locals("user_id", claims.UserID)
		c.Locals("role", claims.Role)
		c.Locals("permissions", claims.Permissions)

		return c.Next()
	}
}