package service

import (
	"go-todo/internal/domain"
	"go-todo/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(userName, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list domain.TodoList) (int, error)
	GetAll(userId int) ([]domain.TodoList, error)
	GetById(userId int, todoListId int) (domain.TodoList, error)
	Update(userId int, todoListId int, input domain.UpdateTodoListInput) error
	Delete(userId int, todoListId int) error
}

type Task interface{}

type Service struct {
	Authorization
	TodoList
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
