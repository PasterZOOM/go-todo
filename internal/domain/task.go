package domain

type Task struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type UpdateTaskInput struct {
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
	Done        *bool   `json:"done" db:"done"`
}

type TodoListTask struct {
	ID         int `json:"id" db:"id"`
	TodoListID int `json:"todoListId" db:"todo_list_id"`
	TaskID     int `json:"taskId" db:"task_id"`
}
