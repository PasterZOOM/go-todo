package domain

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" binding:"required"`
	UserName  string    `json:"userName" db:"user_name" binding:"required"`
	Email     string    `json:"email" db:"email" binding:"required,email"`
	Password  string    `json:"-" db:"password" binding:"required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type UserInput struct {
	Name     string `json:"name" db:"name" binding:"required"`
	UserName string `json:"userName" db:"user_name" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required,email"`
	Password string `json:"password" db:"password" binding:"required"`
}

type UserResponse struct {
	Data UserData `json:"data"`
	Meta UserMeta `json:"meta"`
}

type UserData struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

type UserMeta struct {
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
