package repository

import (
	"context"
	"projectbase/app/model"

	"github.com/google/uuid"
)

// =====================
// PostgreSQL Repository
// =====================

// type UserRepository interface {
// 	FindByEmail(email string) (*model.User, error)
// 	FindByID(id string) (*model.User, error)
// 	Create(user *model.User) error
// }

type UserRepository interface {
	FindAll(limit, offset int) ([]model.User, int64, error)
	FindByID(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id string) error
	UpdateRole(userID, roleID string) error
}

type RoleRepository interface {
	GetByID(id string) (*model.Role, error)
	GetPermissions(roleID string) ([]model.Permission, error)
}

type StudentRepository interface {
	FindAll() ([]model.Student, error)
	FindByID(id string) (*model.Student, error)
	GetByUserID(userID string) (*model.Student, error)
	UpdateAdvisor(studentID, advisorID string) error
	FindByAdvisorID(advisorID string) ([]model.Student, error)
	GetByAdvisorID(advisorID string) ([]model.Student, error)
}

type LecturerRepository interface {
	FindAll() ([]model.Lecturer, error)
	FindByID(id string) (*model.Lecturer, error)
	GetByUserID(userID string) (*model.Lecturer, error)
}

type AchievementReferenceRepository interface {
	Create(ref *model.AchievementReference) error
	GetByStudentID(studentID string) ([]model.AchievementReference, error)
	GetByStudentIDs(studentIDs []uuid.UUID) ([]model.AchievementReference, error)
	UpdateStatus(id string, status string) error
	UpdateVerification(id string, lecturerID string) error
	Reject(id string, note string) error
	GetByID(id string) (*model.AchievementReference, error)
	GetHistory(id string) ([]model.AchievementReference, error)
	GetAll() ([]model.AchievementReference, error)
	MarkDeleted(refID string) error
}

type UserService interface {
	GetUsers(limit, page int) ([]model.User, int64, error)
	GetUserByID(id string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(id string, user *model.User) error
	DeleteUser(id string) error
	AssignRole(userID, roleID string) error
}

type AchievementRepository interface {
	CreateAchievement(ctx context.Context, achievement *model.Achievement) error
	FindByID(ctx context.Context, id string) (*model.Achievement, error)
	FindByStudent(ctx context.Context, studentID string) ([]model.Achievement, error)
	UpdateAchievement(ctx context.Context, id string, update map[string]any) error
	DeleteAchievement(ctx context.Context, id string) error
	AddAttachment(ctx context.Context, id string, attachment model.AchievementAttachment) error
}
