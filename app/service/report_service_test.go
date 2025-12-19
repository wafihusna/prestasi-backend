package service

import (
	"context"
	"testing"

	"projectbase/app/model"
	"projectbase/app/repository/mocks"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
)

func TestGetGlobalStatistics(t *testing.T) {
	// mock repository
	refRepo := new(mocks.AchievementReferenceRepositoryMock)

	refRepo.
		On("GetAll").
		Return([]model.AchievementReference{
			{Status: "draft"},
			{Status: "submitted"},
			{Status: "submitted"},
		}, nil)

	service := NewReportService(
		nil, // achRepo tidak dipakai
		nil, // studentRepo tidak dipakai
		refRepo,
	)

	stats, err := service.GetGlobalStatistics(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, stats)
	assert.Equal(t, 3, stats.TotalAchievements)
	assert.Equal(t, 1, stats.ByStatus["draft"])
	assert.Equal(t, 2, stats.ByStatus["submitted"])
}
