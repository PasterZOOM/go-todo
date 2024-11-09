package repository

import (
	"fmt"
	"go-todo/internal/domain"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user domain.UserInput) (*domain.User, error) {
	var createdUser domain.User
	query := fmt.Sprintf(
		"INSERT INTO %s (name, user_name, password_hash, email, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id, name, user_name, email, created_at, updated_at",
		usersTable,
	)

	row := r.db.QueryRow(query, user.Name, user.UserName, user.Password, user.Email)

	if err := row.Scan(
		&createdUser.ID,
		&createdUser.Name,
		&createdUser.UserName,
		&createdUser.Email,
		&createdUser.CreatedAt,
		&createdUser.UpdatedAt,
	); err != nil {
		logrus.Info(err.Error())
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_user_name_key\"" {
			return nil, fmt.Errorf("username %s is already taken", user.UserName)
		}
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			return nil, fmt.Errorf("email %s is already in use", user.Email)
		}
		return nil, fmt.Errorf("unable to insert user: %w", err)
	}

	return &createdUser, nil
}

func (r *AuthPostgres) GetUser(userName, password string) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf(
		"SELECT id FROM %s WHERE user_name = $1 AND password_hash = $2",
		usersTable,
	)

	err := r.db.Get(&user, query, userName, password)

	return user, err
}
