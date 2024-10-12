package service

import (
	"go-todo/internal/domain"
	"go-todo/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
}

type TodoList interface{}

type Task interface{}

type Service struct {
	Authorization
	TodoList
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
