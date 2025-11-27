package model

import "time"

type AchievementAttachment struct {
	FileName   string    `json:"fileName"`
	FileURL    string    `json:"fileUrl"`
	FileType   string    `json:"fileType"`
	UploadedAt time.Time `json:"uploadedAt"`
}