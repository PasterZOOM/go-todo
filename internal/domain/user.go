package domain

type User struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name" binding:"required"`
	UserName string `json:"userName" db:"user_name" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
