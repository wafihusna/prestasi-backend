package model

type AchievementResponse struct {
	RefID   string       `json:"ref_id"`
	MongoID string       `json:"mongo_id"`
	Data    Achievement `json:"data"`
}