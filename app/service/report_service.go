package service

import (
	// "context"
	// "time"

	"projectbase/app/repository"

	// "go.mongodb.org/mongo-driver/bson"
)

type ReportService struct {
	achRepo    repository.AchievementRepository
	studentRepo repository.StudentRepository
	refRepo     repository.AchievementReferenceRepository
}

func NewReportService(
	achRepo repository.AchievementRepository,
	studentRepo repository.StudentRepository,
	refRepo repository.AchievementReferenceRepository,
) *ReportService {
	return &ReportService{
		achRepo:     achRepo,
		studentRepo: studentRepo,
		refRepo:     refRepo,
	}
}
