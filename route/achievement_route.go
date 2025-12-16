package route

import (
	"context"
	"strings"

	"projectbase/app/model"
	"projectbase/app/service"
	"projectbase/middleware"

	"github.com/gofiber/fiber/v2"
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

	ach.Get("/",
		middleware.JWTMiddleware(),
		func(c *fiber.Ctx) error {

			role := strings.ToLower(c.Locals("role").(string))
			userID := c.Locals("user_id").(string)
			ctx := context.Background()

			switch role {

			case "mahasiswa":
				data, err := svc.ListAchievementsByStudent(ctx, userID)
				if err != nil {
					return c.Status(500).JSON(fiber.Map{"message": err.Error()})
				}
				return c.JSON(data)

			case "dosen wali":
				data, err := svc.ListAchievementsByAdvisor(ctx, userID)
				if err != nil {
					return c.Status(500).JSON(fiber.Map{"message": err.Error()})
				}
				return c.JSON(data)

			case "admin":
				data, err := svc.ListAllAchievements(ctx)
				if err != nil {
					return c.Status(500).JSON(fiber.Map{"message": err.Error()})
				}
				return c.JSON(data)

			default:
				return c.Status(403).JSON(fiber.Map{
					"message": "forbidden",
				})
			}
		},
	)

	// =========================
	// GET DETAIL ACHIEVEMENT
	// =========================
	ach.Get("/:id",
		middleware.JWTMiddleware(),
		func(c *fiber.Ctx) error {
			refID := c.Params("id")

			data, err := svc.GetAchievementDetail(
				context.Background(),
				refID, // UUID
			)
			if err != nil {
				return c.Status(404).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			return c.JSON(data)
		},
	)
}
