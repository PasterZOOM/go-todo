package repository

import (
	"fmt"
	"go-todo/internal/domain"
	"strings"

	"github.com/jmoiron/sqlx"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (r *TaskPostgres) Create(todoListId int, task domain.Task) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var taskId int
	createTaskQuery := fmt.Sprintf(
		"INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id",
		tasksTable,
	)

	row := tx.QueryRow(createTaskQuery, task.Title, task.Description)
	if err := row.Scan(&taskId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createTodoListsTasksQuery := fmt.Sprintf(
		"INSERT INTO %s (todo_list_id, task_id) VALUES ($1, $2)",
		todoListsTasksTable,
	)
	_, err = tx.Exec(createTodoListsTasksQuery, todoListId, taskId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return taskId, tx.Commit()
}

func (r *TaskPostgres) GetAllTasks(userId int, todoListId int) ([]domain.Task, error) {
	var tasks []domain.Task
	query := fmt.Sprintf(
		"SELECT t.id, t.title, t.description FROM %s t INNER JOIN %s tlt on tlt.task_id = t.id INNER JOIN %s utl on utl.todo_list_id = tlt.todo_list_id WHERE tlt.todo_list_id = $1 AND utl.user_id = $2",
		tasksTable,
		todoListsTasksTable,
		usersTodoListsTable,
	)

	if err := r.db.Select(&tasks, query, todoListId, userId); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskPostgres) GetById(userId int, taskId int) (domain.Task, error) {
	var task domain.Task
	query := fmt.Sprintf(
		"SELECT t.id, t.title, t.description, t.done FROM %s t INNER JOIN %s tlt on tlt.task_id = t.id INNER JOIN %s utl on utl.todo_list_id = tlt.todo_list_id WHERE t.id = $1 AND utl.user_id = $2",
		tasksTable,
		todoListsTasksTable,
		usersTodoListsTable,
	)

	if err := r.db.Get(&task, query, taskId, userId); err != nil {
		return task, err
	}

	return task, nil
}

func (r *TaskPostgres) Update(
	userId int,
	taskId int,
	input domain.UpdateTaskInput,
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

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	if len(setValues) == 0 {
		return fmt.Errorf("no values to update")
	}

	query := fmt.Sprintf(
		"UPDATE %s t SET %s FROM %s tlt, %s utl WHERE t.id = tlt.task_id AND tlt.todo_list_id = utl.todo_list_id AND utl.user_id = $%d AND t.id = $%d",
		tasksTable,
		setQuery,
		todoListsTasksTable,
		usersTodoListsTable,
		argId,
		argId+1,
	)
	args = append(args, userId, taskId)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TaskPostgres) Delete(userId int, taskId int) error {
	query := fmt.Sprintf(
		"DELETE FROM %s t USING %s tlt, %s utl WHERE t.id = tlt.task_id AND tlt.todo_list_id = utl.todo_list_id AND utl.user_id = $1 AND t.id = $2",
		tasksTable,
		todoListsTasksTable,
		usersTodoListsTable,
	)

	_, err := r.db.Exec(query, userId, taskId)
	return err
}
