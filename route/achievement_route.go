package route

import (
	"context"
	"strings"
	"time"

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
			userID := c.Locals("user_id").(string)

			if err := svc.SubmitAchievement(userID, refID); err != nil {
				return c.Status(400).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			return c.JSON(fiber.Map{
				"status": "submitted",
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

	// =========================
	// UPDATE ACHIEVEMENT (DRAFT)
	// =========================
	ach.Put("/:id",
		middleware.JWTMiddleware(),
		middleware.RequireRole("mahasiswa"),
		func(c *fiber.Ctx) error {

			refID := c.Params("id")
			userID := c.Locals("user_id").(string)

			var payload map[string]any
			if err := c.BodyParser(&payload); err != nil {
				return c.Status(400).JSON(fiber.Map{
					"message": "invalid request body",
				})
			}

			if err := svc.UpdateAchievement(
				context.Background(),
				userID,
				refID,
				payload,
			); err != nil {
				return c.Status(400).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			return c.JSON(fiber.Map{
				"message": "achievement updated",
			})
		},
	)

	// =========================
	// DELETE ACHIEVEMENT (DRAFT)
	// =========================
	ach.Delete("/:id",
		middleware.JWTMiddleware(),
		middleware.RequireRole("mahasiswa"),
		func(c *fiber.Ctx) error {

			refID := c.Params("id")
			userID := c.Locals("user_id").(string)

			if err := svc.DeleteAchievement(
				context.Background(),
				userID,
				refID,
			); err != nil {
				return c.Status(400).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			return c.JSON(fiber.Map{
				"message": "achievement deleted successfully",
			})
		},
	)

	// =========================
	// FR-007 VERIFY
	// =========================
	ach.Post("/:id/verify",
		middleware.JWTMiddleware(),
		middleware.RequireRole("dosen wali"),
		func(c *fiber.Ctx) error {

			refID := c.Params("id")
			userID := c.Locals("user_id").(string)

			if err := svc.VerifyAchievement(userID, refID); err != nil {
				return c.Status(400).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			return c.JSON(fiber.Map{
				"status": "verified",
			})
		},
	)

	// =========================
	// FR-008 REJECT
	// =========================
	ach.Post("/:id/reject",
		middleware.JWTMiddleware(),
		middleware.RequireRole("dosen wali"),
		func(c *fiber.Ctx) error {

			refID := c.Params("id")

			var body struct {
				Note string `json:"note"`
			}

			if err := c.BodyParser(&body); err != nil {
				return c.Status(400).JSON(fiber.Map{
					"message": "invalid request body",
				})
			}

			if body.Note == "" {
				return c.Status(400).JSON(fiber.Map{
					"message": "rejection note is required",
				})
			}

			if err := svc.RejectAchievement(refID, body.Note); err != nil {
				return c.Status(400).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			return c.JSON(fiber.Map{
				"status": "rejected",
			})
		},
	)

	// =========================
	// GET ACHIEVEMENT HISTORY
	// =========================
	ach.Get("/:id/history",
		middleware.JWTMiddleware(),
		func(c *fiber.Ctx) error {

			refID := c.Params("id")

			data, err := svc.GetAchievementHistory(refID)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			return c.JSON(data)
		},
	)

	// =========================
	// UPLOAD ATTACHMENT
	// =========================
	ach.Post("/:id/attachments",
		middleware.JWTMiddleware(),
		middleware.RequireRole("mahasiswa"),
		func(c *fiber.Ctx) error {

			refID := c.Params("id")

			file, err := c.FormFile("file")
			if err != nil {
				return c.Status(400).JSON(fiber.Map{
					"message": "file is required",
				})
			}

			// 📁 Simpan file (local storage)
			savePath := "./uploads/" + file.Filename
			if err := c.SaveFile(file, savePath); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": "failed to save file",
				})
			}

			attachment := model.AchievementAttachment{
				FileName:   file.Filename,
				FileURL:    savePath,
				FileType:   file.Header.Get("Content-Type"),
				UploadedAt: time.Now(),
			}

			if err := svc.AddAttachment(
				context.Background(),
				refID,
				attachment,
			); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			return c.JSON(fiber.Map{
				"message": "attachment uploaded successfully",
			})
		},
	)
}
