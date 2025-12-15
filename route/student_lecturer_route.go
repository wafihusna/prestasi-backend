package route

import (
	"projectbase/app/service"
	"projectbase/middleware"

	"github.com/gofiber/fiber/v2"
)

func StudentLecturerRoute(
	api fiber.Router,
	studentService service.StudentService,
	lecturerService service.LecturerService,
) {

	students := api.Group("/students",
		middleware.JWTMiddleware(),
	)

	students.Get("/",
		middleware.RequireRole("admin"),
		func(c *fiber.Ctx) error {
			data, err := studentService.GetStudents()
			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": err.Error(),
				})
			}
			return c.JSON(data)
		},
	)

	students.Get("/:id",
		middleware.RequireRole("admin"),
		func(c *fiber.Ctx) error {
			id := c.Params("id")

			data, err := studentService.GetStudentByID(id)
			if err != nil {
				return c.Status(404).JSON(fiber.Map{
					"message": "student not found",
				})
			}

			return c.JSON(data)
		},
	)

	students.Get("/:id/achievements",
		middleware.RequireAnyRole("admin", "lecturer"),
		func(c *fiber.Ctx) error {
			return c.Status(501).JSON(fiber.Map{
				"message": "Get student achievements not implemented yet",
			})
		},
	)

	students.Put("/:id/advisor",
		middleware.RequireRole("admin"),
		func(c *fiber.Ctx) error {
			id := c.Params("id")

			var req struct {
				AdvisorID string `json:"advisor_id"`
			}

			if err := c.BodyParser(&req); err != nil {
				return c.Status(400).JSON(fiber.Map{
					"message": "invalid request body",
				})
			}

			if err := studentService.AssignAdvisor(id, req.AdvisorID); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			return c.JSON(fiber.Map{
				"message": "advisor assigned successfully",
			})
		},
	)

	lecturers := api.Group("/lecturers",
		middleware.JWTMiddleware(),
		middleware.RequireRole("admin"),
	)

	lecturers.Get("/", func(c *fiber.Ctx) error {
		data, err := lecturerService.GetLecturers()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(data)
	})

	lecturers.Get("/:id/advisees",
		middleware.RequireAnyRole("admin", "lecturer"),
		func(c *fiber.Ctx) error {
			lecturerID := c.Params("id")

			data, err := lecturerService.GetLecturerAdvisees(lecturerID)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			return c.JSON(data)
		},
	)
}
