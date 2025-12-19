package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "projectbase/docs"

	"projectbase/config"
	"projectbase/database"
	"projectbase/app/repository"
	"projectbase/app/service"
	"projectbase/route"
)

// @title ProjectBase API
// @version 1.0
// @description API Documentation for Student Achievement System
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.LoadConfig()

	pg := database.ConnectPostgres(
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
	)

	mongoClient := database.ConnectMongo(cfg.MongoURI)

	// Repository
	userRepo := repository.NewUserRepository(pg)
	roleRepo := repository.NewRoleRepository(pg)
	studentRepo := repository.NewStudentRepository(pg)
	lecturerRepo := repository.NewLecturerRepository(pg)

	achievementRepo := repository.NewAchievementRepository(
		mongoClient,
		cfg.MongoDB,
	)
	refRepo := repository.NewAchievementReferenceRepository(pg)

	// Service
	authService := service.NewAuthService(userRepo, roleRepo)
	userService := service.NewUserService(userRepo, roleRepo)
	studentService := service.NewStudentService(studentRepo, achievementRepo, lecturerRepo)
	lecturerService := service.NewLecturerService(lecturerRepo, studentRepo)
	achievementService := service.NewAchievementService(
		achievementRepo,
		refRepo,
		studentRepo,
		lecturerRepo,
	)
	reportService := service.NewReportService(
		achievementRepo,
		studentRepo,
		refRepo,
	)

	app := fiber.New()

	// Swagger endpoint
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api/v1")

	route.AuthRoute(api, authService)
	route.AchievementRoute(api, achievementService)
	route.UserRoute(api, userService)
	route.StudentLecturerRoute(api, studentService, lecturerService)
	route.ReportRoute(api, reportService)

	log.Println("🚀 Server running on port", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
