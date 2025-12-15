package route

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"projectbase/app/model"
	"projectbase/app/service"
	"projectbase/middleware"
)

func AchievementRoute(
	api fiber.Router,
	svc *service.AchievementService,
) {
	ach := api.Group("/achievements")

	// =========================
	// FR-003 CREATE PRESTASI
	// =========================
	ach.Post("/",
		middleware.JWTMiddleware(),
		middleware.RequireRole("mahasiswa"),
		func(c *fiber.Ctx) error {

			var req model.Achievement
			if err := c.BodyParser(&req); err != nil {
				return c.Status(400).JSON(fiber.Map{
					"message": "invalid request",
				})
			}

			studentID := c.Locals("user_id").(string)

			if err := svc.CreateAchievement(
				context.Background(),
				studentID,
				&req,
			); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			return c.JSON(req)
		},
	)

	// =========================
	// FR-004 SUBMIT
	// =========================
	ach.Post("/:id/submit",
		middleware.JWTMiddleware(),
		middleware.RequireRole("mahasiswa"),
		func(c *fiber.Ctx) error {

			refID := c.Params("id")

			if err := svc.SubmitAchievement(refID); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			return c.JSON(fiber.Map{
				"status": "submitted",
			})
		},
	)

	// =========================
	// FR-005 DELETE DRAFT
	// =========================
	ach.Delete("/:id",
		middleware.JWTMiddleware(),
		middleware.RequireRole("mahasiswa"),
		func(c *fiber.Ctx) error {

			refID := c.Params("id")
			mongoID := c.Query("mongo_id")

			if err := svc.DeleteDraftAchievement(
				context.Background(),
				refID,
				mongoID,
			); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			return c.JSON(fiber.Map{
				"message": "achievement deleted",
			})
		},
	)
}
