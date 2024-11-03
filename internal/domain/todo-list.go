package domain

type TodoList struct {
	ID          int    `json:"id" db:"id" example:"1"`
	Title       string `json:"title" db:"title" binding:"required" example:"Todo list 1"`
	Description string `json:"description" db:"description"  example:"Description for todo list 1"`
}

type UpdateTodoListInput struct {
	Title       *string `json:"title" db:"title" example:"New title"`
	Description *string `json:"description" db:"description" example:"New description"`
}

type UserTodolist struct {
	ID         int `json:"id" db:"id"`
	UserID     int `json:"userId" db:"user_id"`
	TodoListID int `json:"todoListId" db:"todo_list_id"`
}
