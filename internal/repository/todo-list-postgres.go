package repository

import (
	"fmt"
	"go-todo/internal/domain"
	"strings"

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
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	createUsersTodoListsQuery := fmt.Sprintf(
		"INSERT INTO %s (user_id, todo_list_id) VALUES ($1, $2)",
		usersTodoListsTable,
	)
	_, err = tx.Exec(createUsersTodoListsQuery, userID, id)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId int) ([]domain.TodoList, error) {
	var todoLists []domain.TodoList

	query := fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s utl on tl.id = utl.todo_list_id WHERE utl.user_id = $1",
		todoListTable, usersTodoListsTable,
	)

	err := r.db.Select(&todoLists, query, userId)

	return todoLists, err
}

func (r *TodoListPostgres) GetById(userId int, todoListId int) (domain.TodoList, error) {
	var todoList domain.TodoList

	query := fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s utl on tl.id = utl.todo_list_id WHERE utl.user_id = $1 AND utl.todo_list_id = $2",
		todoListTable, usersTodoListsTable,
	)

	err := r.db.Get(&todoList, query, userId, todoListId)

	return todoList, err
}

func (r *TodoListPostgres) Update(
	userId int,
	todoListId int,
	input domain.UpdateTodoListInput,
) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	if len(setValues) == 0 {
		return fmt.Errorf("no values to update")
	}

	query := fmt.Sprintf(
		"UPDATE %s tl SET %s FROM %s utl WHERE tl.id = utl.todo_list_id AND utl.user_id = $%d AND utl.todo_list_id = $%d",
		todoListTable,
		setQuery,
		usersTodoListsTable,
		argId,
		argId+1,
	)
	args = append(args, userId, todoListId)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TodoListPostgres) Delete(userId int, todoListId int) error {
	query := fmt.Sprintf(
		"DELETE FROM %s tl USING %s utl WHERE tl.id = utl.todo_list_id AND utl.user_id = $1 AND utl.todo_list_id = $2",
		todoListTable,
		usersTodoListsTable,
	)

	_, err := r.db.Exec(query, userId, todoListId)
	return err
}
