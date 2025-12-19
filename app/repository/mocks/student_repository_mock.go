package mocks

import (
	"projectbase/app/model"

	"github.com/stretchr/testify/mock"
)

type StudentRepositoryMock struct {
	mock.Mock
}

func (m *StudentRepositoryMock) FindAll() ([]model.Student, error) {
	args := m.Called()
	return args.Get(0).([]model.Student), args.Error(1)
}

func (m *StudentRepositoryMock) FindByID(id string) (*model.Student, error) {
	args := m.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Student), args.Error(1)
}

func (m *StudentRepositoryMock) GetByUserID(userID string) (*model.Student, error) {
	args := m.Called(userID)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Student), args.Error(1)
}

func (m *StudentRepositoryMock) UpdateAdvisor(studentID, advisorID string) error {
	args := m.Called(studentID, advisorID)
	return args.Error(0)
}

func (m *StudentRepositoryMock) FindByAdvisorID(advisorID string) ([]model.Student, error) {
	args := m.Called(advisorID)
	return args.Get(0).([]model.Student), args.Error(1)
}

func (m *StudentRepositoryMock) GetByAdvisorID(advisorID string) ([]model.Student, error) {
	args := m.Called(advisorID)
	return args.Get(0).([]model.Student), args.Error(1)
}