package service

import (
	"projectbase/app/model"
	"projectbase/app/repository"
)

type LecturerService struct {
	lecturerRepo repository.LecturerRepository
}

func NewLecturerService(lecturerRepo repository.LecturerRepository) *LecturerService {
	return &LecturerService{lecturerRepo}
}

func (s *LecturerService) GetLecturerByUserID(userID string) (*model.Lecturer, error) {
	return s.lecturerRepo.GetByUserID(userID)
}
