package model

import (
	"time"
	"github.com/google/uuid"
)

type AchievementReference struct {
	ID                  uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	StudentID           uuid.UUID   `gorm:"type:uuid;not null"`
	MongoAchievementID string `gorm:"type:text;not null"`
	RefID   string              `json:"ref_id"`
	Status              string
	SubmittedAt         *time.Time
	VerifiedAt          *time.Time
	VerifiedBy          *uuid.UUID  `gorm:"type:uuid"`
	RejectionNote       *string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

