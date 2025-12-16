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

func (r *studentRepository) FindAll() ([]model.Student, error) {
	var students []model.Student
	err := r.db.Find(&students).Error
	return students, err
}

func (r *studentRepository) FindByID(id string) (*model.Student, error) {
	var s model.Student
	err := r.db.First(&s, "id = ?", id).Error
	return &s, err
}

func (r *studentRepository) UpdateAdvisor(studentID, advisorID string) error {
	return r.db.Model(&model.Student{}).
		Where("id = ?", studentID).
		Update("advisor_id", advisorID).Error
}

func (r *studentRepository) FindByAdvisorID(advisorID string) ([]model.Student, error) {
	var students []model.Student
	err := r.db.Where("advisor_id = ?", advisorID).Find(&students).Error
	return students, err
}

func (r *studentRepository) GetByUserID(userID string) (*model.Student, error) {
	var s model.Student
	err := r.db.First(&s, "user_id = ?", userID).Error
	return &s, err
}

func (r *studentRepository) GetByAdvisorID(advisorID string) ([]model.Student, error) {
	var students []model.Student
	err := r.db.
		Where("advisor_id = ?", advisorID).
		Find(&students).Error
	return students, err
}