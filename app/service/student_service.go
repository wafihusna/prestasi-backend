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

func (s *studentService) GetStudents() ([]model.Student, error) {
	return s.studentRepo.FindAll()
}

func (s *studentService) GetStudentByID(id string) (*model.Student, error) {
	return s.studentRepo.FindByID(id)
}

func (s *studentService) GetStudentAchievements(studentID string) ([]any, error) {
	// DIHANDLE DI IMPLEMENTASI (Mongo + reference)
	return nil, nil
}

func (s *studentService) AssignAdvisor(studentID, advisorID string) error {
	return s.studentRepo.UpdateAdvisor(studentID, advisorID)
}