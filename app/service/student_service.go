package service

import (
	"context"
	"projectbase/app/model"
	"projectbase/app/repository"

	"github.com/gofiber/fiber/v2"
)

type studentService struct {
	studentRepo     repository.StudentRepository
	achievementRepo repository.AchievementRepository
	lecturerRepo    repository.LecturerRepository
}

func NewStudentService(
	studentRepo repository.StudentRepository,
	achievementRepo repository.AchievementRepository,
	lecturerRepo repository.LecturerRepository,
) StudentService {
	return &studentService{
		studentRepo:     studentRepo,
		achievementRepo: achievementRepo,
		lecturerRepo:    lecturerRepo,
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

func (s *studentService) GetStudentAchievementsForLecturer(
	studentUUID string,
	lecturerUserID string,
) ([]model.Achievement, error) {

	// 1️⃣ Ambil student target
	student, err := s.studentRepo.FindByID(studentUUID)
	if err != nil {
		return nil, err
	}

	// 2️⃣ Ambil lecturer dari USER ID
	lecturer, err := s.lecturerRepo.GetByUserID(lecturerUserID)
	if err != nil {
		return nil, fiber.ErrForbidden
	}

	// 3️⃣ Cek apakah dosen adalah dosen wali
	if student.AdvisorID != lecturer.ID {
		return nil, fiber.ErrForbidden
	}

	// 4️⃣ Ambil achievement
	return s.achievementRepo.FindByStudent(
		context.Background(),
		student.StudentID,
	)
}
