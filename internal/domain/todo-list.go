package domain

type TodoList struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UpdateTodoListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type UserTodolist struct {
	ID         int `json:"id" db:"id"`
	UserID     int `json:"userId" db:"user_id"`
	TodoListID int `json:"todoListId" db:"todo_list_id"`
}
