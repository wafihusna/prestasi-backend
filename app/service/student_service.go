package service

import (
	"projectbase/app/model"
	"projectbase/app/repository"
)

type StudentService struct {
	studentRepo repository.StudentRepository
}

func NewStudentService(studentRepo repository.StudentRepository) *StudentService {
	return &StudentService{studentRepo}
}

func (s *StudentService) GetStudentByUserID(userID string) (*model.Student, error) {
	return s.studentRepo.GetByUserID(userID)
}