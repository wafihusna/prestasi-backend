package mocks

import (
	"projectbase/app/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type AchievementReferenceRepositoryMock struct {
	mock.Mock
}

func (m *AchievementReferenceRepositoryMock) Create(
	ref *model.AchievementReference,
) error {
	args := m.Called(ref)
	return args.Error(0)
}

func (m *AchievementReferenceRepositoryMock) GetByStudentID(
	studentID string,
) ([]model.AchievementReference, error) {
	args := m.Called(studentID)
	return args.Get(0).([]model.AchievementReference), args.Error(1)
}

func (m *AchievementReferenceRepositoryMock) GetByStudentIDs(
	studentIDs []uuid.UUID,
) ([]model.AchievementReference, error) {
	args := m.Called(studentIDs)
	return args.Get(0).([]model.AchievementReference), args.Error(1)
}

func (m *AchievementReferenceRepositoryMock) UpdateStatus(
	id string,
	status string,
) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *AchievementReferenceRepositoryMock) UpdateVerification(
	id string,
	lecturerID string,
) error {
	args := m.Called(id, lecturerID)
	return args.Error(0)
}

func (m *AchievementReferenceRepositoryMock) Reject(
	id string,
	note string,
) error {
	args := m.Called(id, note)
	return args.Error(0)
}

func (m *AchievementReferenceRepositoryMock) GetByID(
	id string,
) (*model.AchievementReference, error) {
	args := m.Called(id)
	return args.Get(0).(*model.AchievementReference), args.Error(1)
}

func (m *AchievementReferenceRepositoryMock) GetHistory(
	id string,
) ([]model.AchievementReference, error) {
	args := m.Called(id)
	return args.Get(0).([]model.AchievementReference), args.Error(1)
}

func (m *AchievementReferenceRepositoryMock) GetAll() ([]model.AchievementReference, error) {
	args := m.Called()
	return args.Get(0).([]model.AchievementReference), args.Error(1)
}

func (m *AchievementReferenceRepositoryMock) MarkDeleted(
	refID string,
) error {
	args := m.Called(refID)
	return args.Error(0)
}