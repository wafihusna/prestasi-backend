package service

import "projectbase/app/model"

type UserService interface {
	GetUsers(limit, page int) ([]model.User, int64, error)
	GetUserByID(id string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(id string, user *model.User) error
	DeleteUser(id string) error
	AssignRole(userID, roleID string) error
}

type StudentService interface {
	GetStudents() ([]model.Student, error)
	GetStudentByID(id string) (*model.Student, error)
	GetStudentAchievements(studentID string) ([]any, error) // Mongo handled di impl
	AssignAdvisor(studentID, advisorID string) error
}

type LecturerService interface {
	GetLecturers() ([]model.Lecturer, error)
	GetLecturerAdvisees(lecturerID string) ([]model.Student, error)
}