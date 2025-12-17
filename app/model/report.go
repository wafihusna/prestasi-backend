package model

type AchievementStatistic struct {
	TotalAchievements int                    `json:"totalAchievements"`
	ByStatus          map[string]int         `json:"byStatus"`
	ByStudent         map[string]int         `json:"byStudent,omitempty"`
}