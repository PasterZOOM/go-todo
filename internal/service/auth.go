package service

import (
	"crypto/sha1"
	"fmt"
	"go-todo/internal/domain"
	"go-todo/internal/repository"
)

const salt = "aksdjhl;csxjvkjbf"

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func (s *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)

	return s.repos.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
