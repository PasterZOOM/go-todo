package repository

import (
	"database/sql"
	"log"
	"regexp"
	"testing"
	"time"

	"go-todo/internal/domain"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestUserPostgres_GetUserData(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("error closing db: %s", err)
		}
	}(db)

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	userRepo := NewUserPostgres(sqlxDB)

	userId := 1
	user := &domain.User{
		ID:        userId,
		Name:      "John Doe",
		UserName:  "johndoe",
		Email:     "john.doe@example.com",
		CreatedAt: time.Now().Truncate(time.Second),
		UpdatedAt: time.Now().Truncate(time.Second),
	}

	query := regexp.QuoteMeta("SELECT id, name, user_name, email, created_at, updated_at FROM users WHERE id = $1")

	t.Run(
		"GetUserData_Success", func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"id", "name", "user_name", "email", "created_at", "updated_at"}).
				AddRow(user.ID, user.Name, user.UserName, user.Email, user.CreatedAt, user.UpdatedAt)
			mock.ExpectQuery(query).WithArgs(userId).WillReturnRows(rows)

			result, err := userRepo.GetUserData(userId)

			assert.NoError(t, err)
			assert.Equal(t, user.ID, result.ID)
			assert.Equal(t, user.Name, result.Name)
			assert.Equal(t, user.UserName, result.UserName)
			assert.Equal(t, user.Email, result.Email)
			assert.WithinDuration(t, user.CreatedAt, result.CreatedAt, time.Second)
			assert.WithinDuration(t, user.UpdatedAt, result.UpdatedAt, time.Second)
			assert.NoError(t, mock.ExpectationsWereMet())
		},
	)

	t.Run(
		"GetUserData_NotFound", func(t *testing.T) {
			mock.ExpectQuery(query).WithArgs(userId).WillReturnError(sql.ErrNoRows)

			result, err := userRepo.GetUserData(userId)
			logrus.Info(result)
			logrus.Info(err)
			assert.Error(t, err)
			assert.Nil(t, result)
			assert.NoError(t, mock.ExpectationsWereMet())
		},
	)
}
