package repository

import (
	"go-todo/internal/domain"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type TodoList interface {
	Create(userId int, todoList domain.TodoList) (int, error)
	GetAll(userId int) ([]domain.TodoList, error)
	GetById(userId int, todoListId int) (domain.TodoList, error)
	Update(userId int, todoListId int, input domain.UpdateTodoListInput) error
	Delete(userId int, todoListId int) error
}

type Task interface{}

type Repository struct {
	Authorization
	TodoList
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
