package sql

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/crisnlopez/social-media-bkend/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	return db, mock
}

func TestCreateUserRepository(t *testing.T) {
	db, mock := NewMock()
	repo := NewUserRepository(db)
	defer repo.db.Close()

	query := `INSERT INTO users (email, pass, name, age, nick) VALUES (?, ?, ?, ?, ?)`

	userRequest := user.UserRequest{
		Email: "random@email.com",
		Pass:  "pass1234",
		Name:  "fulano",
		Age:   20,
		Nick:  "fulanoNick",
	}

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(userRequest.Email, userRequest.Pass, userRequest.Name, userRequest.Age, userRequest.Nick).WillReturnResult(sqlmock.NewResult(1, 1))

	user, err := repo.CreateUser(userRequest)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}

func TestGetUserRepository(t *testing.T) {
	db, mock := NewMock()
	repo := UserQueries{db}
	defer repo.db.Close()

	query := `SELECT * FROM users WHERE id = ?`

	user := user.User{
		ID:    1,
		Email: "random@email.com",
		Pass:  "pass1234",
		Name:  "fulano",
		Age:   20,
		Nick:  "fulanoNick",
	}

	rows := sqlmock.NewRows([]string{"id", "email", "pass", "nick", "name", "age", "created_at"}).AddRow(user.ID, user.Email, user.Pass, user.Nick, user.Name, user.Age, user.CreatedAt)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(user.ID).WillReturnRows(rows)

	user, err := repo.GetUser(user.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}

func TestUpdateUserRepository(t *testing.T) {
	db, mock := NewMock()
	repo := UserQueries{db}
	defer repo.db.Close()

	query := `UPDATE users SET email = ?, pass = ?, name= ?,  age= ?, nick= ? WHERE id = ?`

	userRequest := user.UserRequest{
		Email: "random@email.com",
		Pass:  "pass1234",
		Name:  "fulano",
		Age:   20,
		Nick:  "fulanoNick",
	}

	id := 1

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(userRequest.Email, userRequest.Pass, userRequest.Name, userRequest.Age, userRequest.Nick, id).WillReturnResult(sqlmock.NewResult(0, 1)).WillReturnError(nil)

	row, err := repo.UpdateUser(userRequest, int64(id))
	require.NoError(t, err)
	require.Equal(t, int64(1), row)
}

func TestDeleteUserRepository(t *testing.T) {
	db, mock := NewMock()
	repo := UserQueries{db}
	defer func() {
		repo.db.Close()
	}()

	query := `DELETE FROM users WHERE id = ?`

	id := 1

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 1)).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.DeleteUser(id)
	require.NoError(t, err)
}
