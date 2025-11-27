package repository

import (
	"projectbase/app/model"

	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db}
}

func (r *studentRepository) GetByUserID(userID string) (*model.Student, error) {
	var s model.Student
	err := r.db.First(&s, "user_id = ?", userID).Error
	return &s, err
}