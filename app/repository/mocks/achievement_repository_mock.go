package mocks

import (
	"context"
	"projectbase/app/model"

	"github.com/stretchr/testify/mock"
)

type AchievementRepositoryMock struct {
	mock.Mock
}

func (m *AchievementRepositoryMock) CreateAchievement(
	ctx context.Context,
	achievement *model.Achievement,
) error {
	args := m.Called(ctx, achievement)
	return args.Error(0)
}

func (m *AchievementRepositoryMock) FindByID(
	ctx context.Context,
	id string,
) (*model.Achievement, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Achievement), args.Error(1)
}

func (m *AchievementRepositoryMock) FindByStudent(
	ctx context.Context,
	studentID string,
) ([]model.Achievement, error) {
	args := m.Called(ctx, studentID)
	return args.Get(0).([]model.Achievement), args.Error(1)
}

func (m *AchievementRepositoryMock) UpdateAchievement(
	ctx context.Context,
	id string,
	update map[string]any,
) error {
	args := m.Called(ctx, id, update)
	return args.Error(0)
}

func (m *AchievementRepositoryMock) DeleteAchievement(
	ctx context.Context,
	id string,
) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *AchievementRepositoryMock) AddAttachment(
	ctx context.Context,
	id string,
	attachment model.AchievementAttachment,
) error {
	args := m.Called(ctx, id, attachment)
	return args.Error(0)
}