package service

import (
	"errors"
	"go-todo/internal/domain"
	"go-todo/internal/repository/mocks"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthService_CreateUser(t *testing.T) {
	mockRepo := new(mocks.Authorization)
	authService := NewAuthService(mockRepo)

	userInput := domain.UserInput{
		Name:     "John Doe",
		UserName: "johndoe",
		Password: "password",
		Email:    "john@example.com",
	}

	hashedPassword := generatePasswordHash(userInput.Password)
	expectedUser := &domain.User{
		ID:        1,
		Name:      userInput.Name,
		UserName:  userInput.UserName,
		Email:     userInput.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("CreateUser", mock.AnythingOfType("domain.UserInput")).Return(expectedUser, nil).
		Run(
			func(args mock.Arguments) {
				user := args.Get(0).(domain.UserInput)
				assert.Equal(t, hashedPassword, user.Password)
			},
		)

	userResponse, err := authService.CreateUser(userInput)

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

func TestAuthService_GenerateToken(t *testing.T) {
	mockRepo := new(mocks.Authorization)
	authService := NewAuthService(mockRepo)

	userName := "johndoe"
	password := "password"
	hashedPassword := generatePasswordHash(password)

	expectedUser := domain.User{
		ID:       1,
		UserName: userName,
		Password: hashedPassword,
	}

	mockRepo.On("GetUser", userName, hashedPassword).Return(expectedUser, nil)

	token, err := authService.GenerateToken(userName, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedToken, err := jwt.ParseWithClaims(
		token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(signingKey), nil
		},
	)
	assert.NoError(t, err)

	claims, ok := parsedToken.Claims.(*TokenClaims)
	assert.True(t, ok)
	assert.Equal(t, expectedUser.ID, claims.UserID)
	mockRepo.AssertExpectations(t)
}

func TestAuthService_GenerateToken_InvalidCredentials(t *testing.T) {
	mockRepo := new(mocks.Authorization)
	authService := NewAuthService(mockRepo)

	userName := "johndoe"
	password := "wrongpassword"
	hashedPassword := generatePasswordHash(password)

	mockRepo.On("GetUser", userName, hashedPassword).Return(domain.User{}, errors.New("user not found"))

	token, err := authService.GenerateToken(userName, password)

	assert.Error(t, err)
	assert.Equal(t, "", token)
	mockRepo.AssertExpectations(t)
}

func TestAuthService_ParseToken(t *testing.T) {
	authService := NewAuthService(nil)

	userID := 1
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, &TokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(tokenTTL).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			userID,
		},
	)

	tokenString, err := token.SignedString([]byte(signingKey))
	assert.NoError(t, err)

	parsedUserID, err := authService.ParseToken(tokenString)
	assert.NoError(t, err)
	assert.Equal(t, userID, parsedUserID)
}

func TestAuthService_ParseToken_InvalidToken(t *testing.T) {
	authService := NewAuthService(nil)

	invalidToken := "invalidToken"

	userID, err := authService.ParseToken(invalidToken)

	assert.Error(t, err)
	assert.Equal(t, 0, userID)
}
