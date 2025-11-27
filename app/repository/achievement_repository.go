package repository

import (
	"context"
	"projectbase/app/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type achievementRepository struct {
	col *mongo.Collection
}

func NewAchievementRepository(db *mongo.Client, dbName string) AchievementRepository {
	return &achievementRepository{
		col: db.Database(dbName).Collection("achievements"),
	}
}

func (r *achievementRepository) CreateAchievement(ctx context.Context, achievement *model.Achievement) error {
	_, err := r.col.InsertOne(ctx, achievement)
	return err
}

func (r *achievementRepository) FindByID(ctx context.Context, id string) (*model.Achievement, error) {
	var res model.Achievement
	err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&res)
	return &res, err
}

func (r *achievementRepository) FindByStudent(ctx context.Context, studentID string) ([]model.Achievement, error) {
	cursor, err := r.col.Find(ctx, bson.M{"studentId": studentID})
	if err != nil {
		return nil, err
	}

	var list []model.Achievement
	err = cursor.All(ctx, &list)
	return list, err
}

func (r *achievementRepository) UpdateAchievement(ctx context.Context, id string, update map[string]any) error {
	_, err := r.col.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

func (r *achievementRepository) DeleteAchievement(ctx context.Context, id string) error {
	_, err := r.col.DeleteOne(ctx, bson.M{"_id": id})
	return err
}