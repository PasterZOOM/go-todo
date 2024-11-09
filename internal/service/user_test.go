package service

import (
	"errors"
	"go-todo/internal/domain"
	"go-todo/internal/repository/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserService_GetUserData(t *testing.T) {
	mockRepo := new(mocks.User)
	userService := NewUserService(mockRepo)

	userID := 1
	expectedUser := &domain.User{
		ID:        userID,
		Name:      "John Doe",
		UserName:  "johndoe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("GetUserData", userID).Return(expectedUser, nil)

	userResponse, err := userService.GetUserData(userID)

	assert.NoError(t, err)
	assert.NotNil(t, userResponse)
	assert.Equal(t, expectedUser.ID, userResponse.Data.ID)
	assert.Equal(t, expectedUser.Name, userResponse.Data.Name)
	assert.Equal(t, expectedUser.UserName, userResponse.Data.UserName)
	assert.Equal(t, expectedUser.Email, userResponse.Data.Email)
	assert.Equal(t, expectedUser.CreatedAt, userResponse.Meta.CreatedAt)
	assert.Equal(t, expectedUser.UpdatedAt, userResponse.Meta.UpdatedAt)

	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserData_NotFound(t *testing.T) {
	mockRepo := new(mocks.User)

	userID := 2
	mockRepo.On("GetUserData", userID).Return(nil, errors.New("user not found"))

	userService := NewUserService(mockRepo)

	userResponse, err := userService.GetUserData(userID)

	assert.Error(t, err)
	assert.Nil(t, userResponse)
	assert.EqualError(t, err, "user not found")

	mockRepo.AssertExpectations(t)
}
