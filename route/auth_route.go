package route

import (
	"github.com/gofiber/fiber/v2"
	"projectbase/app/service"
	"projectbase/utils"
	"projectbase/middleware"
)

func AuthRoute(api fiber.Router, authService *service.AuthService) {

	auth := api.Group("/auth")

	// =======================
	// LOGIN
	// =======================
	auth.Post("/login", func(c *fiber.Ctx) error {

		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "invalid request",
			})
		}

		token, user, err := authService.Login(req.Email, req.Password)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"token": token,
			"user": fiber.Map{
				"id":    user.ID,
				"email": user.Email,
				"name":  user.FullName,
				"role":  user.RoleID,
			},
		})
	})

	// =======================
	// PROFILE
	// =======================
	auth.Get("/profile",
		middleware.JWTMiddleware(),
		func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"user_id": c.Locals("user_id"),
				"role":    c.Locals("role"),
			})
		},
	)

	// =======================
	// REFRESH TOKEN
	// =======================
	auth.Post("/refresh",
		middleware.JWTMiddleware(),
		func(c *fiber.Ctx) error {

			userID := c.Locals("user_id").(string)
			role := c.Locals("role").(string)

			permissions := c.Locals("permissions").([]string)

			token, _ := utils.GenerateJWT(userID, role, permissions)

			return c.JSON(fiber.Map{
				"token": token,
			})
		},
	)

	// =======================
	// LOGOUT (STATELESS)
	// =======================
	auth.Post("/logout", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "logout success (client side)",
		})
	})
}