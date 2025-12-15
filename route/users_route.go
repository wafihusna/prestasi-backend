package route

import (
	"projectbase/app/model"
	"projectbase/app/service"
	"projectbase/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(api fiber.Router, userService service.UserService) {

	users := api.Group("/users",
		middleware.JWTMiddleware(),
		middleware.RequireRole("admin"),
	)

	// =========================
	// GET /api/v1/users
	// =========================
	users.Get("/", func(c *fiber.Ctx) error {
		limit := c.QueryInt("limit", 10)
		page := c.QueryInt("page", 1)

		data, total, err := userService.GetUsers(limit, page)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data": data,
			"meta": fiber.Map{
				"total": total,
				"page":  page,
				"limit": limit,
			},
		})
	})

	// =========================
	// GET /api/v1/users/:id
	// =========================
	users.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		user, err := userService.GetUserByID(id)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"message": "user not found",
			})
		}

		return c.JSON(user)
	})

	// =========================
	// POST /api/v1/users
	// =========================
	users.Post("/", func(c *fiber.Ctx) error {
		var req struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
			FullName string `json:"full_name"`
			RoleID   string `json:"role_id"`
		}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "invalid request body",
			})
		}

		user := &model.User{
			Username: req.Username,
			Email:    req.Email,
			FullName: req.FullName,
			RoleID:   req.RoleID,
			// PasswordHash seharusnya di-hash di service / auth flow
			PasswordHash: req.Password,
		}

		if err := userService.CreateUser(user); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "user created successfully",
			"id":      user.ID,
		})
	})

	// =========================
	// PUT /api/v1/users/:id
	// =========================
	users.Put("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var req struct {
			Email    string `json:"email"`
			FullName string `json:"full_name"`
			IsActive bool   `json:"is_active"`
		}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "invalid request body",
			})
		}

		user := &model.User{
			Email:    req.Email,
			FullName: req.FullName,
			IsActive: req.IsActive,
		}

		if err := userService.UpdateUser(id, user); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "user updated successfully",
		})
	})

	// =========================
	// DELETE /api/v1/users/:id
	// =========================
	users.Delete("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		if err := userService.DeleteUser(id); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "user deleted successfully",
		})
	})

	// =========================
	// PUT /api/v1/users/:id/role
	// =========================
	users.Put("/:id/role", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var req struct {
			RoleID string `json:"role_id"`
		}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "invalid request body",
			})
		}

		if err := userService.AssignRole(id, req.RoleID); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "role assigned successfully",
		})
	})
}
