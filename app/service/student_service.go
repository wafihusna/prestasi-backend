package service

import (
	"context"
	"projectbase/app/model"
	"projectbase/app/repository"
)

type studentService struct {
	studentRepo     repository.StudentRepository
	achievementRepo repository.AchievementRepository
}

func NewStudentService(
	studentRepo repository.StudentRepository,
	achievementRepo repository.AchievementRepository,
) StudentService {
	return &studentService{
		studentRepo:     studentRepo,
		achievementRepo: achievementRepo,
	}
}

func (s *studentService) GetStudentByUserID(userID string) (*model.Student, error) {
	return s.studentRepo.GetByUserID(userID)
}

func (s *studentService) GetStudents() ([]model.Student, error) {
	return s.studentRepo.FindAll()
}

func (s *studentService) GetStudentByID(id string) (*model.Student, error) {
	return s.studentRepo.FindByID(id)
}

func (s *studentService) AssignAdvisor(studentID, advisorID string) error {
	return s.studentRepo.UpdateAdvisor(studentID, advisorID)
}

func (s *studentService) GetStudentAchievements(
	studentUUID string,
) ([]model.Achievement, error) {

	// 1️⃣ Ambil student dari PostgreSQL
	student, err := s.studentRepo.FindByID(studentUUID)
	if err != nil {
		return nil, err
	}

	// 2️⃣ Pakai NIM untuk Mongo
	return s.achievementRepo.FindByStudent(
		context.Background(),
		student.StudentID,
	)
}
