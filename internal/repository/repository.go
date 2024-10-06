package repository

type Authorization interface{}

type TodoList interface{}

type Task interface{}

type Repository struct {
	Authorization
	TodoList
	Task
}

func NewRepository() *Repository {
	return &Repository{}
}
