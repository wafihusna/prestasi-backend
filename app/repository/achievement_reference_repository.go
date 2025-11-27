package repository

import (
	"projectbase/app/model"

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
