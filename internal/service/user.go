package service

import (
	"go-todo/internal/domain"
	"go-todo/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserData(userId int) (*domain.UserResponse, error) {
	user, err := s.repo.GetUserData(userId)
	if err != nil {
		return nil, err
	}

	return generateUserResponse(user), nil
}

func generateUserResponse(user *domain.User) *domain.UserResponse {
	response := &domain.UserResponse{
		Data: domain.UserData{
			ID:       user.ID,
			Name:     user.Name,
			UserName: user.UserName,
			Email:    user.Email,
		},
		Meta: domain.UserMeta{
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}

	return response
}
