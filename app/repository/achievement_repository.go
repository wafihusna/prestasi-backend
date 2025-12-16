package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"projectbase/app/model"
)

type achievementRepository struct {
	col *mongo.Collection
}

type achievementMongo struct {
	ID        primitive.ObjectID `bson:"_id"`
	StudentID string             `bson:"studentId"`
	Title     string             `bson:"title"`
	Details   any                `bson:"details"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

func NewAchievementRepository(db *mongo.Client, dbName string) AchievementRepository {
	return &achievementRepository{
		col: db.Database(dbName).Collection("achievements"),
	}
}

// ✅ CREATE
func (r *achievementRepository) CreateAchievement(
	ctx context.Context,
	ach *model.Achievement,
) error {
	res, err := r.col.InsertOne(ctx, ach)
	if err != nil {
		return err
	}

	objID := res.InsertedID.(primitive.ObjectID)
	ach.ID = objID.Hex()
	return nil
}

// ✅ FIND BY ID
func (r *achievementRepository) FindByID(
	ctx context.Context,
	id string,
) (*model.Achievement, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var res model.Achievement
	err = r.col.FindOne(ctx, bson.M{"_id": objID}).Decode(&res)
	return &res, err
}

// ✅ FIND BY STUDENT
func (r *achievementRepository) FindByStudent(
	ctx context.Context,
	studentID string,
) ([]model.Achievement, error) {

	var list []model.Achievement

	cursor, err := r.col.Find(ctx, bson.M{"studentId": studentID})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &list)
	return list, err
}

// ✅ UPDATE
func (r *achievementRepository) UpdateAchievement(
	ctx context.Context,
	id string,
	update map[string]any,
) error {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.col.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": update},
	)
	return err
}

// ✅ DELETE
func (r *achievementRepository) DeleteAchievement(
	ctx context.Context,
	id string,
) error {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.col.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (r *achievementRepository) AddAttachment(
	ctx context.Context,
	mongoID string,
	attachment model.AchievementAttachment,
) error {

	oid, err := primitive.ObjectIDFromHex(mongoID)
	if err != nil {
		return err
	}

	_, err = r.col.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{
			"$push": bson.M{
				"attachments": attachment,
			},
		},
	)

	return err
}
