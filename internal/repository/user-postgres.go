package repository

import (
	"fmt"
	"go-todo/internal/domain"

	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetUserData(userId int) (*domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT id, name, user_name, email, created_at, updated_at FROM %s WHERE id = $1", usersTable)

	err := r.db.QueryRow(query, userId).Scan(
		&user.ID,
		&user.Name,
		&user.UserName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
