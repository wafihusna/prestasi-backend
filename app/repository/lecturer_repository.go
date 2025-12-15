package repository

import (
	"projectbase/app/model"

	"gorm.io/gorm"
)

type lecturerRepository struct {
	db *gorm.DB
}

func NewLecturerRepository(db *gorm.DB) LecturerRepository {
	return &lecturerRepository{db}
}

func (r *lecturerRepository) FindAll() ([]model.Lecturer, error) {
	var lecturers []model.Lecturer
	err := r.db.Find(&lecturers).Error
	return lecturers, err
}

func (r *lecturerRepository) FindByID(id string) (*model.Lecturer, error) {
	var l model.Lecturer
	err := r.db.First(&l, "id = ?", id).Error
	return &l, err
}


func (r *lecturerRepository) GetByUserID(userID string) (*model.Lecturer, error) {
	var l model.Lecturer
	err := r.db.First(&l, "user_id = ?", userID).Error
	return &l, err
}