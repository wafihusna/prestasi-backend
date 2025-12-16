package service

import (
	"context"
	"time"
	"errors"

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

func (s *AchievementService) DeleteAchievement(ctx context.Context, id string) error {
	return s.achievementRepo.DeleteAchievement(ctx, id)
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

func (s *AchievementService) ListAchievementsByStudent(
	ctx context.Context,
	userID string,
) ([]model.Achievement, error) {

	student, err := s.studentRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	return s.achievementRepo.FindByStudent(ctx, student.ID)
}

func (s *AchievementService) GetAchievementDetail(
	ctx context.Context,
	refID string,
) (*model.Achievement, error) {

	ref, err := s.refRepo.GetByID(refID)
	if err != nil {
		return nil, err
	}

	return s.achievementRepo.FindByID(ctx, ref.MongoAchievementID)
}

func (s *AchievementService) UpdateAchievementDraft(
	ctx context.Context,
	refID string,
	update map[string]any,
) error {

	ref, err := s.refRepo.GetByID(refID)
	if err != nil {
		return err
	}

	if ref.Status != "draft" {
		return errors.New("only draft achievement can be updated")
	}

	update["updatedAt"] = time.Now()
	return s.achievementRepo.UpdateAchievement(ctx, ref.MongoAchievementID, update)
}

func (s *AchievementService) DeleteDraft(
	ctx context.Context,
	refID string,
) error {

	ref, err := s.refRepo.GetByID(refID)
	if err != nil {
		return err
	}

	if ref.Status != "draft" {
		return errors.New("only draft achievement can be deleted")
	}

	if err := s.achievementRepo.DeleteAchievement(ctx, ref.MongoAchievementID); err != nil {
		return err
	}

	return s.refRepo.UpdateStatus(refID, "deleted")
}

func (s *AchievementService) SubmitAchievement(refID string) error {
	return s.refRepo.UpdateStatus(refID, "submitted")
}

func (s *AchievementService) VerifyAchievement(
	refID string,
	lecturerID string,
) error {
	return s.refRepo.UpdateVerification(refID, lecturerID)
}

func (s *AchievementService) RejectAchievement(
	refID string,
	note string,
) error {
	return s.refRepo.Reject(refID, note)
}

func (s *AchievementService) GetAchievementHistory(
	refID string,
) ([]model.AchievementReference, error) {
	return s.refRepo.GetHistory(refID)
}

func (s *AchievementService) AddAttachment(
	ctx context.Context,
	refID string,
	attachment model.AchievementAttachment,
) error {

	ref, err := s.refRepo.GetByID(refID)
	if err != nil {
		return err
	}

	attachment.UploadedAt = time.Now()
	return s.achievementRepo.AddAttachment(ctx, ref.MongoAchievementID, attachment)
}

func (s *AchievementService) ListAchievementsByAdvisor(
	ctx context.Context,
	lecturerID string,
) ([]model.Achievement, error) {

	// 1️⃣ ambil mahasiswa bimbingan
	students, err := s.studentRepo.GetByAdvisorID(lecturerID)
	if err != nil {
		return nil, err
	}

	if len(students) == 0 {
		return []model.Achievement{}, nil
	}

	// 2️⃣ kumpulkan student IDs
	var studentIDs []string
	for _, s := range students {
		studentIDs = append(studentIDs, s.ID)
	}

	// 3️⃣ ambil reference
	refs, err := s.refRepo.GetByStudentIDs(studentIDs)
	if err != nil {
		return nil, err
	}

	// 4️⃣ ambil detail dari Mongo
	var result []model.Achievement
	for _, ref := range refs {
		ach, err := s.achievementRepo.FindByID(ctx, ref.MongoAchievementID)
		if err == nil {
			result = append(result, *ach)
		}
	}

	return result, nil
}

func (s *AchievementService) ListAllAchievements(
	ctx context.Context,
) ([]model.Achievement, error) {

	refs, err := s.refRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []model.Achievement
	for _, ref := range refs {
		ach, err := s.achievementRepo.FindByID(ctx, ref.MongoAchievementID)
		if err == nil {
			result = append(result, *ach)
		}
	}

	return result, nil
}