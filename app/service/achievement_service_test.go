package service

import (
	"context"
	"testing"

	"projectbase/app/model"
	"projectbase/app/repository/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	validStudentUUID     = "550e8400-e29b-41d4-a716-446655440000"
	validAchievementUUID = "111e8400-e29b-41d4-a716-446655440000"
)

func TestCreateAchievementDraft(t *testing.T) {
	mockAchRepo := new(mocks.AchievementRepositoryMock)
	mockStudentRepo := new(mocks.StudentRepositoryMock)

	mockStudentRepo.
		On("GetByUserID", validStudentUUID).
		Return(&model.Student{
			ID: validStudentUUID,
		}, nil)

	mockAchRepo.
		On("CreateAchievement", mock.Anything, mock.Anything).
		Return(nil)

	service := AchievementService{
		achievementRepo: mockAchRepo,
		studentRepo:     mockStudentRepo,
	}

	err := service.CreateAchievement(
		context.Background(),
		validStudentUUID,
		&model.Achievement{
			Title: "Juara 1 AI",
		},
	)

	assert.NoError(t, err)
}

func TestSubmitAchievement(t *testing.T) {
	mockAchRepo := new(mocks.AchievementRepositoryMock)
	mockRefRepo := new(mocks.AchievementReferenceRepositoryMock)

	mockAchRepo.
		On("UpdateStatus", mock.Anything, validAchievementUUID, "submitted").
		Return(nil)

	mockRefRepo.
		On("UpdateStatus", mock.Anything, validAchievementUUID, "submitted").
		Return(nil)

	service := AchievementService{
		achievementRepo: mockAchRepo,
		refRepo:         mockRefRepo,
	}

	err := service.SubmitAchievement(
		validAchievementUUID,
		validStudentUUID,
	)

	assert.NoError(t, err)
}

func TestDeleteDraftAchievement(t *testing.T) {
	mockAchRepo := new(mocks.AchievementRepositoryMock)

	mockAchRepo.
		On("DeleteAchievement", mock.Anything, validAchievementUUID).
		Return(nil)

	service := AchievementService{
		achievementRepo: mockAchRepo,
	}

	err := service.DeleteAchievement(
		context.Background(),
		validAchievementUUID,
		validStudentUUID,
	)

	assert.NoError(t, err)
}
