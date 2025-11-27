package repository

import (
	"context"
	"projectbase/app/model"
)

// =====================
// PostgreSQL Repository
// =====================

type UserRepository interface {
	FindByEmail(email string) (*model.User, error)
	FindByID(id string) (*model.User, error)
	Create(user *model.User) error
}

type RoleRepository interface {
	GetByID(id string) (*model.Role, error)
	GetPermissions(roleID string) ([]model.Permission, error)
}

type StudentRepository interface {
	GetByUserID(userID string) (*model.Student, error)
}

type LecturerRepository interface {
	GetByUserID(userID string) (*model.Lecturer, error)
}

type AchievementReferenceRepository interface {
	Create(ref *model.AchievementReference) error
	GetByStudentID(studentID string) ([]model.AchievementReference, error)
	UpdateStatus(id string, status string) error
}

// ===============
// Mongo Repository
// ===============

type AchievementRepository interface {
	CreateAchievement(ctx context.Context, achievement *model.Achievement) error
	FindByID(ctx context.Context, id string) (*model.Achievement, error)
	FindByStudent(ctx context.Context, studentID string) ([]model.Achievement, error)
	UpdateAchievement(ctx context.Context, id string, update map[string]any) error
	DeleteAchievement(ctx context.Context, id string) error
}