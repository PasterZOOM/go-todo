package repository

import "github.com/jmoiron/sqlx"

type Authorization interface{}

type TodoList interface{}

type Task interface{}

type Repository struct {
	Authorization
	TodoList
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
