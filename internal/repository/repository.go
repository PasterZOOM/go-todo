package repository

import (
	"go-todo/internal/domain"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
}

type TodoList interface{}

type Task interface{}

type Repository struct {
	Authorization
	TodoList
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
