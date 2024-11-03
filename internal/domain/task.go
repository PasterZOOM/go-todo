package domain

type Task struct {
	ID          int    `json:"id" db:"id" example:"1"`
	Title       string `json:"title" db:"title" binding:"required" example:"Task 1"`
	Description string `json:"description" db:"description" example:"Description for task 1"`
	Done        bool   `json:"done" db:"done" example:"false"`
}

type UpdateTaskInput struct {
	Title       *string `json:"title" db:"title" example:"New task title"`
	Description *string `json:"description" db:"description" example:"New task description"`
	Done        *bool   `json:"done" db:"done" example:"true"`
}

type TodoListTask struct {
	ID         int `json:"id" db:"id"`
	TodoListID int `json:"todoListId" db:"todo_list_id"`
	TaskID     int `json:"taskId" db:"task_id"`
}
