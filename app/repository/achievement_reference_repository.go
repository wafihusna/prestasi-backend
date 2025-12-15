package repository

import (
	"projectbase/app/model"
	"time"

	"gorm.io/gorm"
)

type achievementRefRepo struct {
	db *gorm.DB
}

func NewAchievementReferenceRepository(db *gorm.DB) AchievementReferenceRepository {
	return &achievementRefRepo{db}
}

func (r *achievementRefRepo) Create(ref *model.AchievementReference) error {
	return r.db.Create(ref).Error
}

// func (r *achievementRefRepo) GetByUserID(userID string) ([]model.AchievementReference, error) {
// 	var refs []model.AchievementReference

// 	err := r.db.
// 		Joins("JOIN students ON students.id = achievement_references.student_id").
// 		Where("students.user_id = ?", userID).
// 		Find(&refs).Error

// 	return refs, err
// }

func (r *achievementRefRepo) GetByStudentID(studentID string) ([]model.AchievementReference, error) {
	var refs []model.AchievementReference
	err := r.db.Where("student_id = ?", studentID).Find(&refs).Error
	return refs, err
}

func (r *achievementRefRepo) UpdateStatus(id string, status string) error {
	return r.db.Model(&model.AchievementReference{}).
		Where("id = ?", id).
		Update("status", status).Error
}

func (r *achievementRefRepo) UpdateVerification(
	id string,
	status string,
	lecturerID string,
) error {
	return r.db.Model(&model.AchievementReference{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"status":      status,
			"verified_by": lecturerID,
			"verified_at": time.Now(),
		}).Error
}

func (r *achievementRefRepo) Reject(id string, note string) error {
	return r.db.Model(&model.AchievementReference{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"status":         "rejected",
			"rejection_note": note,
		}).Error
}
