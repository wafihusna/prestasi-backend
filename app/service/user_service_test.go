package service

import (
	"testing"

	"projectbase/app/model"
	"projectbase/app/repository/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserSuccess(t *testing.T) {
	// arrange
	mockUserRepo := new(mocks.UserRepositoryMock)

	user := &model.User{
		Email:        "test@mail.com",
		PasswordHash: "password123", 
		FullName:     "Test User",
	}

	// Create(user) 
	mockUserRepo.
		On("Create", mock.AnythingOfType("*model.User")).
		Return(nil)

	service := NewUserService(
		mockUserRepo,
		nil, 
	)

	// act
	err := service.CreateUser(user)

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.PasswordHash)
	assert.True(t, user.IsActive)

	mockUserRepo.AssertExpectations(t)
}
