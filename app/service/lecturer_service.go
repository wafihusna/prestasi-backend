package service

import (
	"projectbase/app/model"
	"projectbase/app/repository"
)

type lecturerService struct {
	lecturerRepo repository.LecturerRepository
	studentRepo  repository.StudentRepository
}

func NewLecturerService(
	lecturerRepo repository.LecturerRepository,
	studentRepo repository.StudentRepository,
) LecturerService {
	return &lecturerService{
		lecturerRepo: lecturerRepo,
		studentRepo:  studentRepo,
	}
}

func (s *lecturerService) GetLecturers() ([]model.Lecturer, error) {
	return s.lecturerRepo.FindAll()
}

func (s *lecturerService) GetLecturerAdvisees(lecturerID string) ([]model.Student, error) {
	return s.studentRepo.FindByAdvisorID(lecturerID)
}
