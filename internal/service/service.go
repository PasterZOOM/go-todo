package service

import "go-todo/internal/repository"

type Authorization interface{}

type TodoList interface{}

type Task interface{}

type Service struct {
	Authorization
	TodoList
	Task
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
