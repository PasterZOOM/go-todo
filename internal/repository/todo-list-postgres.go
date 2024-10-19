package repository

import (
	"fmt"
	"go-todo/internal/domain"

	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userID int, todoList domain.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createTodoListQuery := fmt.Sprintf(
		"INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id",
		todoListTable,
	)

	row := tx.QueryRow(createTodoListQuery, todoList.Title, todoList.Description)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersTodoListsQuery := fmt.Sprintf(
		"INSERT INTO %s (user_id, todo_list_id) VALUES ($1, $2)",
		usersTodoListsTable,
	)
	_, err = tx.Exec(createUsersTodoListsQuery, userID, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId int) ([]domain.TodoList, error) {
	var todoLists []domain.TodoList

	query := fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.todo_list_id WHERE ul.user_id = $1",
		todoListTable, usersTodoListsTable,
	)

	err := r.db.Select(&todoLists, query, userId)

	return todoLists, err
}

func (r *TodoListPostgres) GetById(userId int, todoListId int) (domain.TodoList, error) {
	var todoList domain.TodoList

	query := fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.todo_list_id WHERE ul.user_id = $1 AND ul.todo_list_id = $2",
		todoListTable, usersTodoListsTable,
	)

	err := r.db.Get(&todoList, query, userId, todoListId)

	return todoList, err
}
