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
	GetStudentByUserID(userID string) (*model.Student, error)
	GetStudentAchievements(studentUUID string) ([]model.Achievement, error)
	GetStudentAchievementsForLecturer(
		studentUUID string,
		lecturerUserID string,
	) ([]model.Achievement, error)
	AssignAdvisor(studentID, advisorID string) error
}

type LecturerService interface {
	GetLecturers() ([]model.Lecturer, error)
	GetLecturerAdvisees(lecturerID string) ([]model.Student, error)
}
