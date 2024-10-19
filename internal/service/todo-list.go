package service

import (
	"go-todo/internal/domain"
	"go-todo/internal/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, todoList domain.TodoList) (int, error) {
	return s.repo.Create(userId, todoList)
}

func (s *TodoListService) GetAll(userId int) ([]domain.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId int, todoListId int) (domain.TodoList, error) {
	return s.repo.GetById(userId, todoListId)
}
