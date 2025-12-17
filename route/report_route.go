package route

import (
	"context"

	"projectbase/app/service"
	"projectbase/middleware"

	"github.com/gofiber/fiber/v2"
)

func ReportRoute(
	api fiber.Router,
	svc *service.ReportService,
) {

	rep := api.Group("/reports", middleware.JWTMiddleware())

	// =========================
	// GET GLOBAL STATISTICS
	// =========================
	rep.Get("/statistics", func(c *fiber.Ctx) error {

		data, err := svc.GetGlobalStatistics(context.Background())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(data)
	})

	// =========================
	// GET STUDENT STATISTICS
	// =========================
	rep.Get("/student/:id", func(c *fiber.Ctx) error {

		studentID := c.Params("id")

		data, err := svc.GetStudentStatistics(
			context.Background(),
			studentID,
		)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(data)
	})
}