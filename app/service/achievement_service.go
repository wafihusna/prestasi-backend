package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"projectbase/app/model"
	"projectbase/app/repository"
)

type AchievementService struct {
	achievementRepo repository.AchievementRepository
	refRepo         repository.AchievementReferenceRepository
	studentRepo     repository.StudentRepository
}

func NewAchievementService(
	achievementRepo repository.AchievementRepository,
	refRepo repository.AchievementReferenceRepository,
	studentRepo repository.StudentRepository,
) *AchievementService {
	return &AchievementService{
		achievementRepo: achievementRepo,
		refRepo:         refRepo,
		studentRepo:     studentRepo,
	}
}

// CreateAchievement = insert ke Mongo + insert ref ke PostgreSQL
// func (s *AchievementService) CreateAchievement(
// 	ctx context.Context,
// 	userID string,
// 	ach *model.Achievement,
// ) error {

// 	// 🔑 ambil student dari user_id
// 	student, err := s.studentRepo.GetByUserID(userID)
// 	if err != nil {
// 		return err
// 	}

// 	ach.StudentID = student.ID
// 	ach.CreatedAt = time.Now()
// 	ach.UpdatedAt = time.Now()

// 	// MongoDB
// 	if err := s.achievementRepo.CreateAchievement(ctx, ach); err != nil {
// 		return err
// 	}

// 	// PostgreSQL reference
// 	ref := &model.AchievementReference{
// 		StudentID:          student.ID,
// 		MongoAchievementID: ach.ID,
// 		Status:             "draft",
// 		CreatedAt:          time.Now(),
// 	}

// 	return s.refRepo.Create(ref)
// }

// CreateAchievement = insert ke Mongo + insert ref ke PostgreSQL
func (s *AchievementService) CreateAchievement(
	ctx context.Context,
	userID string,
	ach *model.Achievement,
) error {

	// 🔑 ambil student dari user_id
	student, err := s.studentRepo.GetByUserID(userID)
	if err != nil {
		return err
	}

	// 🔁 convert student.ID (string) → uuid.UUID
	studentUUID, err := uuid.Parse(student.ID)
	if err != nil {
		return err
	}

	// set data achievement (Mongo)
	ach.StudentID = student.ID
	ach.CreatedAt = time.Now()
	ach.UpdatedAt = time.Now()

	// MongoDB insert
	if err := s.achievementRepo.CreateAchievement(ctx, ach); err != nil {
		return err
	}

	// simpan Mongo ID sebagai string
	mongoID := ach.ID

	// PostgreSQL reference
	ref := &model.AchievementReference{
		StudentID:          studentUUID, // ✅ UUID
		MongoAchievementID: mongoID,     // ✅ string pointer
		Status:             "draft",
		CreatedAt:          time.Now(),
	}

	return s.refRepo.Create(ref)
}

func (s *AchievementService) GetAchievementsByStudent(ctx context.Context, studentID string) ([]model.Achievement, error) {
	return s.achievementRepo.FindByStudent(ctx, studentID)
}

func (s *AchievementService) UpdateAchievement(ctx context.Context, id string, update map[string]any) error {
	update["updatedAt"] = time.Now()
	return s.achievementRepo.UpdateAchievement(ctx, id, update)
}

func (s *AchievementService) DeleteAchievement(ctx context.Context, id string) error {
	return s.achievementRepo.DeleteAchievement(ctx, id)
}

func (s *AchievementService) SubmitAchievement(refID string) error {
	return s.refRepo.UpdateStatus(refID, "submitted")
}

func (s *AchievementService) VerifyAchievement(refID, lecturerID string) error {
	return s.refRepo.UpdateVerification(
		refID,
		"verified",
		lecturerID,
	)
}

func (s *AchievementService) RejectAchievement(refID, note string) error {
	return s.refRepo.Reject(refID, note)
}

func (s *AchievementService) DeleteDraftAchievement(
	ctx context.Context,
	refID string,
	mongoID string,
) error {

	// hapus di Mongo
	if err := s.achievementRepo.DeleteAchievement(ctx, mongoID); err != nil {
		return err
	}

	// update status di PostgreSQL
	return s.refRepo.UpdateStatus(refID, "deleted")
}
