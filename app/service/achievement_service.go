package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"projectbase/app/model"
	"projectbase/app/repository"
)

type AchievementService struct {
	achievementRepo repository.AchievementRepository
	refRepo         repository.AchievementReferenceRepository
	studentRepo     repository.StudentRepository
	lecturerRepo    repository.LecturerRepository
}

func NewAchievementService(
	achievementRepo repository.AchievementRepository,
	refRepo repository.AchievementReferenceRepository,
	studentRepo repository.StudentRepository,
	lecturerRepo repository.LecturerRepository,
) *AchievementService {
	return &AchievementService{
		achievementRepo: achievementRepo,
		refRepo:         refRepo,
		studentRepo:     studentRepo,
		lecturerRepo:    lecturerRepo,
	}
}

func (s *AchievementService) CreateAchievement(
	ctx context.Context,
	userID string,
	ach *model.Achievement,
) error {

	student, err := s.studentRepo.GetByUserID(userID)
	if err != nil {
		return err
	}

	studentUUID, err := uuid.Parse(student.ID)
	if err != nil {
		return err
	}

	ach.StudentID = student.ID
	ach.CreatedAt = time.Now()
	ach.UpdatedAt = time.Now()

	// 🔥 INSERT KE MONGO
	if err := s.achievementRepo.CreateAchievement(ctx, ach); err != nil {
		return err
	}

	// 🔥 AMBIL ID MONGO
	mongoID := ach.ID

	// 🔥 SIMPAN KE POSTGRES
	ref := &model.AchievementReference{
		StudentID:          studentUUID,
		MongoAchievementID: mongoID,
		Status:             "draft",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	return s.refRepo.Create(ref)
}

func (s *AchievementService) GetAchievementsByStudent(ctx context.Context, studentID string) ([]model.Achievement, error) {
	return s.achievementRepo.FindByStudent(ctx, studentID)
}

// func (s *AchievementService) DeleteAchievement(ctx context.Context, id string) error {
// 	return s.achievementRepo.DeleteAchievement(ctx, id)
// }

func (s *AchievementService) DeleteAchievement(
	ctx context.Context,
	userID string,
	refID string,
) error {

	// 1️⃣ Ambil reference PostgreSQL
	ref, err := s.refRepo.GetByID(refID)
	if err != nil {
		return err
	}

	// 2️⃣ Validasi status
	if ref.Status != "draft" {
		return errors.New("only draft achievement can be deleted")
	}

	// 3️⃣ Soft delete MongoDB (PAKAI MONGO ID!)
	err = s.achievementRepo.UpdateAchievement(
		ctx,
		ref.MongoAchievementID, // ✅ INI YANG BENAR
		map[string]any{
			"deleted":   true,
			"deletedAt": time.Now(),
		},
	)
	if err != nil {
		return err
	}

	// 4️⃣ Update PostgreSQL
	return s.refRepo.MarkDeleted(refID)
}

// func (s *AchievementService) DeleteDraftAchievement(
// 	ctx context.Context,
// 	refID string,
// 	mongoID string,
// ) error {

// 	// hapus di Mongo
// 	if err := s.achievementRepo.DeleteAchievement(ctx, mongoID); err != nil {
// 		return err
// 	}

// 	// update status di PostgreSQL
// 	return s.refRepo.UpdateStatus(refID, "deleted")
// }

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

// func (s *AchievementService) UpdateAchievementDraft(
// 	ctx context.Context,
// 	refID string,
// 	update map[string]any,
// ) error {

// 	ref, err := s.refRepo.GetByID(refID)
// 	if err != nil {
// 		return err
// 	}

// 	if ref.Status != "draft" {
// 		return errors.New("only draft achievement can be updated")
// 	}

// 	update["updatedAt"] = time.Now()
// 	return s.achievementRepo.UpdateAchievement(ctx, ref.MongoAchievementID, update)
// }

func (s *AchievementService) UpdateAchievement(
	ctx context.Context,
	userID string,
	refID string,
	payload map[string]any,
) error {

	// 1️⃣ ambil reference
	ref, err := s.refRepo.GetByID(refID)
	if err != nil {
		return err
	}

	// 2️⃣ pastikan status masih draft
	if ref.Status != "draft" {
		return errors.New("only draft achievement can be updated")
	}

	// 3️⃣ pastikan achievement milik mahasiswa tsb
	student, err := s.studentRepo.GetByUserID(userID)
	if err != nil {
		return err
	}

	studentUUID, err := uuid.Parse(student.ID)
	if err != nil {
		return err
	}

	if ref.StudentID != studentUUID {
		return errors.New("forbidden: not your achievement")
	}

	// 4️⃣ set updatedAt
	payload["updatedAt"] = time.Now()

	// 5️⃣ update ke MongoDB
	return s.achievementRepo.UpdateAchievement(
		ctx,
		ref.MongoAchievementID,
		payload,
	)
}

// func (s *AchievementService) DeleteDraft(
// 	ctx context.Context,
// 	refID string,
// ) error {

// 	ref, err := s.refRepo.GetByID(refID)
// 	if err != nil {
// 		return err
// 	}

// 	if ref.Status != "draft" {
// 		return errors.New("only draft achievement can be deleted")
// 	}

// 	if err := s.achievementRepo.DeleteAchievement(ctx, ref.MongoAchievementID); err != nil {
// 		return err
// 	}

// 	return s.refRepo.UpdateStatus(refID, "deleted")
// }

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
	userID string,
) ([]model.Achievement, error) {

	// 1️⃣ ambil lecturer
	lecturer, err := s.lecturerRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	// 2️⃣ ambil mahasiswa bimbingan
	students, err := s.studentRepo.GetByAdvisorID(lecturer.ID)
	if err != nil {
		return nil, err
	}

	if len(students) == 0 {
		return []model.Achievement{}, nil
	}

	// 3️⃣ kumpulkan student UUIDs
	var studentUUIDs []uuid.UUID
	for _, st := range students {
		id, err := uuid.Parse(st.ID)
		if err != nil {
			continue
		}
		studentUUIDs = append(studentUUIDs, id)
	}

	// 4️⃣ ambil references dari Postgre
	refs, err := s.refRepo.GetByStudentIDs(studentUUIDs)
	if err != nil {
		return nil, err
	}

	// 5️⃣ ambil detail Mongo
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
) ([]model.AchievementResponse, error) {

	refs, err := s.refRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []model.AchievementResponse

	for _, ref := range refs {
		ach, err := s.achievementRepo.FindByID(ctx, ref.MongoAchievementID)
		if err != nil {
			continue
		}

		result = append(result, model.AchievementResponse{
			RefID:   ref.ID.String(),
			MongoID: ref.MongoAchievementID,
			Data:    *ach,
		})
	}

	return result, nil
}
