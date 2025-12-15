package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"projectbase/config"
	"projectbase/database"

	"projectbase/app/repository"
	"projectbase/app/service"
	"projectbase/route"
)

func main() {
	cfg := config.LoadConfig()

	// =====================
	// DATABASE
	// =====================
	pg := database.ConnectPostgres(
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
	)

	mongoClient := database.ConnectMongo(cfg.MongoURI)

	// =====================
	// REPOSITORY
	// =====================
	userRepo := repository.NewUserRepository(pg)
	roleRepo := repository.NewRoleRepository(pg)
	studentRepo := repository.NewStudentRepository(pg)
	lecturerService := service.NewLecturerService(lecturerRepo, studentRepo)

	achievementRepo := repository.NewAchievementRepository(
		mongoClient,
		cfg.MongoDB,
	)
	refRepo := repository.NewAchievementReferenceRepository(pg)

	// =====================
	// SERVICE
	// =====================
	authService := service.NewAuthService(userRepo, roleRepo)

	userService := service.NewUserService(
		userRepo,
		roleRepo,
	)

	achievementService := service.NewAchievementService(
		achievementRepo,
		refRepo,
		studentRepo,
	)

	// =====================
	// FIBER
	// =====================
	app := fiber.New()
	api := app.Group("/api/v1")

	// =====================
	// ROUTES
	// =====================
	route.AuthRoute(api, authService)
	route.AchievementRoute(api, achievementService)
	route.UserRoute(api, userService)
	route.StudentLecturerRoute(api, studentService, lecturerService)

	log.Println("🚀 Server running on port", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
