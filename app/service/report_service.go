package service

import (
	"context"

	"projectbase/app/model"
	"projectbase/app/repository"
)

type ReportService struct {
	achRepo     repository.AchievementRepository
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

// =========================
// GLOBAL STATISTICS
// =========================
func (s *ReportService) GetGlobalStatistics(
	ctx context.Context,
) (*model.AchievementStatistic, error) {

	refs, err := s.refRepo.GetAll()
	if err != nil {
		return nil, err
	}

	stat := &model.AchievementStatistic{
		TotalAchievements: len(refs),
		ByStatus:          map[string]int{},
	}

	for _, ref := range refs {
		stat.ByStatus[ref.Status]++
	}

	return stat, nil
}

// =========================
// STUDENT STATISTICS
// =========================
func (s *ReportService) GetStudentStatistics(
	ctx context.Context,
	studentID string,
) (*model.AchievementStatistic, error) {

	refs, err := s.refRepo.GetByStudentID(studentID)
	if err != nil {
		return nil, err
	}

	stat := &model.AchievementStatistic{
		TotalAchievements: len(refs),
		ByStatus:          map[string]int{},
	}

	for _, ref := range refs {
		stat.ByStatus[ref.Status]++
	}

	return stat, nil
}
