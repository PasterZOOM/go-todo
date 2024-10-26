package service

import (
	"go-todo/internal/domain"
	"go-todo/internal/repository"
)

type TaskService struct {
	repo         repository.Task
	todoListRepo repository.TodoList
}

func NewTaskService(repo repository.Task, todoListRepo repository.TodoList) *TaskService {
	return &TaskService{repo: repo, todoListRepo: todoListRepo}
}

func (s *TaskService) Create(userId int, todoListId int, task domain.Task) (int, error) {
	_, err := s.todoListRepo.GetById(userId, todoListId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(todoListId, task)
}

func (s *TaskService) GetAllTasks(userId int, todoListId int) ([]domain.Task, error) {
	return s.repo.GetAllTasks(userId, todoListId)
}

func (s *TaskService) GetById(userId int, taskId int) (domain.Task, error) {
	return s.repo.GetById(userId, taskId)
}
func (s *TaskService) Update(
	userId int,
	taskId int,
	input domain.UpdateTaskInput,
) error {
	return s.repo.Update(userId, taskId, input)
}

func (s *TaskService) Delete(userId int, taskId int) error {
	return s.repo.Delete(userId, taskId)
}
