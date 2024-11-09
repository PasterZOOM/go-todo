package repository

import (
	"errors"
	"go-todo/internal/domain"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func setupDBMock(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock db: %s", err)
	}
	return sqlx.NewDb(db, "postgres"), mock
}

func TestAuthPostgres_CreateUser(t *testing.T) {
	db, mock := setupDBMock(t)
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("error closing db: %s", err)
		}
	}(db)

	repo := NewAuthPostgres(db)

	userInput := domain.UserInput{
		Name:     "John Doe",
		UserName: "johndoe",
		Password: "hashed_password",
		Email:    "john@example.com",
	}

	expectedUser := domain.User{
		ID:        1,
		Name:      userInput.Name,
		UserName:  userInput.UserName,
		Email:     userInput.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := regexp.QuoteMeta(
		"INSERT INTO users (name, user_name, password_hash, email, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id, name, user_name, email, created_at, updated_at",
	)

	rows := sqlmock.NewRows([]string{"id", "name", "user_name", "email", "created_at", "updated_at"}).
		AddRow(expectedUser.ID, expectedUser.Name, expectedUser.UserName, expectedUser.Email, expectedUser.CreatedAt, expectedUser.UpdatedAt)

	mock.ExpectQuery(query).
		WithArgs(userInput.Name, userInput.UserName, userInput.Password, userInput.Email).
		WillReturnRows(rows)

	createdUser, err := repo.CreateUser(userInput)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, createdUser.ID)
	assert.Equal(t, expectedUser.Name, createdUser.Name)
	assert.Equal(t, expectedUser.UserName, createdUser.UserName)
	assert.Equal(t, expectedUser.Email, createdUser.Email)
}

func TestAuthPostgres_GetUser(t *testing.T) {
	db, mock := setupDBMock(t)
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("error closing db: %s", err)
		}
	}(db)

	repo := NewAuthPostgres(db)

	userName := "johndoe"
	password := "hashed_password"

	expectedUser := domain.User{
		ID:        1,
		Name:      "John Doe",
		UserName:  userName,
		Email:     "john@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := regexp.QuoteMeta("SELECT id FROM users WHERE user_name = $1 AND password_hash = $2")

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(expectedUser.ID)

	mock.ExpectQuery(query).
		WithArgs(userName, password).
		WillReturnRows(rows)

	user, err := repo.GetUser(userName, password)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, user.ID)
}

func TestAuthPostgres_CreateUser_DuplicateUserName(t *testing.T) {
	db, mock := setupDBMock(t)
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("error closing db: %s", err)
		}
	}(db)

	repo := NewAuthPostgres(db)

	userInput := domain.UserInput{
		Name:     "John Doe",
		UserName: "johndoe",
		Password: "hashed_password",
		Email:    "john@example.com",
	}

	query := regexp.QuoteMeta(
		"INSERT INTO users (name, user_name, password_hash, email, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id, name, user_name, email, created_at, updated_at",
	)

	mock.ExpectQuery(query).
		WithArgs(userInput.Name, userInput.UserName, userInput.Password, userInput.Email).
		WillReturnError(errors.New("pq: duplicate key value violates unique constraint \"users_user_name_key\""))

	createdUser, err := repo.CreateUser(userInput)

	assert.Nil(t, createdUser)
	assert.EqualError(t, err, "username johndoe is already taken")
}

func TestAuthPostgres_CreateUser_DuplicateEmail(t *testing.T) {
	db, mock := setupDBMock(t)
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("error closing db: %s", err)
		}
	}(db)

	repo := NewAuthPostgres(db)

	userInput := domain.UserInput{
		Name:     "John Doe",
		UserName: "johndoe",
		Password: "hashed_password",
		Email:    "john@example.com",
	}

	query := regexp.QuoteMeta(
		"INSERT INTO users (name, user_name, password_hash, email, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id, name, user_name, email, created_at, updated_at",
	)

	mock.ExpectQuery(query).
		WithArgs(userInput.Name, userInput.UserName, userInput.Password, userInput.Email).
		WillReturnError(errors.New("pq: duplicate key value violates unique constraint \"users_email_key\""))

	createdUser, err := repo.CreateUser(userInput)

	assert.Nil(t, createdUser)
	assert.EqualError(t, err, "email john@example.com is already in use")
}
