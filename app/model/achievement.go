package model

import "time"

type Achievement struct {
	ID              string                  `json:"id" bson:"_id,omitempty"`
	StudentID       string                  `json:"studentId"`
	AchievementType string                  `json:"achievementType"`
	Title           string                  `json:"title"`
	Description     string                  `json:"description"`
	Details         AchievementDetails      `json:"details"`
	Attachments     []AchievementAttachment `json:"attachments"`
	Tags            []string                `json:"tags"`
	Points          int                     `json:"points"`
	CreatedAt       time.Time               `json:"createdAt"`
	UpdatedAt       time.Time               `json:"updatedAt"`
}
