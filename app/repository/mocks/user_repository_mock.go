package mocks

import (
	"projectbase/app/model"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) FindAll(limit, offset int) ([]model.User, int64, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]model.User), args.Get(1).(int64), args.Error(2)
}

func (m *UserRepositoryMock) FindByID(id string) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *UserRepositoryMock) FindByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *UserRepositoryMock) Create(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) Update(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *UserRepositoryMock) UpdateRole(userID, roleID string) error {
	args := m.Called(userID, roleID)
	return args.Error(0)
}