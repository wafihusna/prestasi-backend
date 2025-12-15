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
			return nil // handled by controller
		},
	)

	students.Get("/:id",
		middleware.RequireRole("admin"),
		func(c *fiber.Ctx) error {
			return nil
		},
	)

	students.Get("/:id/achievements",
		middleware.RequireAnyRole("admin", "lecturer"),
		func(c *fiber.Ctx) error {
			return nil
		},
	)

	students.Put("/:id/advisor",
		middleware.RequireRole("admin"),
		func(c *fiber.Ctx) error {
			return nil
		},
	)

	lecturers := api.Group("/lecturers",
		middleware.JWTMiddleware(),
		middleware.RequireRole("admin"),
	)

	lecturers.Get("/", func(c *fiber.Ctx) error {
		return nil
	})

	lecturers.Get("/:id/advisees",
		middleware.RequireAnyRole("admin", "lecturer"),
		func(c *fiber.Ctx) error {
			return nil
		},
	)
}
