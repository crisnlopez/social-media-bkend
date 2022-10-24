package sql

import (
	"database/sql"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/crisnlopez/social-media-bkend/internal/util"

	"github.com/crisnlopez/social-media-bkend/internal/user"
)

var u = user.User{
	ID:        util.RandomInt(0, 1000),
	Email:     util.RandomEmail(),
	Pass:      util.RandomPass(),
	Nick:      util.RandomNick(),
	Name:      util.RandomName(),
	Age:       util.RandomAge(),
	CreatedAt: time.Now().UTC().Local(),
}

var uRequest = user.UserRequest{
	Email: util.RandomEmail(),
	Pass:  util.RandomPass(),
	Nick:  util.RandomNick(),
	Name:  util.RandomName(),
	Age:   util.RandomAge(),
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	return db, mock
}

func TestCreateUserRepository(t *testing.T) {
	db, mock := NewMock()
	repo := UserQueries{db}
	defer func() {
		repo.db.Close()
	}()

	query := `INSERT INTO users (email, pass, name, age, nick) VALUES (?, ?, ?, ?, ?)`

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(uRequest.Email, uRequest.Pass, uRequest.Name, uRequest.Age, uRequest.Nick).WillReturnResult(sqlmock.NewResult(1, 1))

	user, err := repo.CreateUser(uRequest)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}

func TestGetUserRepository(t *testing.T) {
	db, mock := NewMock()
	repo := UserQueries{db}
	defer func() {
		repo.db.Close()
	}()

	query := `SELECT * FROM users WHERE id = ?`

	rows := sqlmock.NewRows([]string{"id", "email", "pass", "nick", "name", "age", "created_at"}).AddRow(u.ID, u.Email, u.Pass, u.Nick, u.Name, u.Age, u.CreatedAt)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(u.ID).WillReturnRows(rows)

	user, err := repo.GetUser(u.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}

func TestUpdateUserRepository(t *testing.T) {
	db, mock := NewMock()
	repo := UserQueries{db}
	defer func() {
		repo.db.Close()
	}()

	query := `UPDATE users SET email = ?, pass = ?, name= ?,  age= ?, nick= ? WHERE id = ?`

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(uRequest.Email, uRequest.Pass, uRequest.Name, uRequest.Age, uRequest.Nick, u.ID).WillReturnResult(sqlmock.NewResult(0, 1)).WillReturnError(nil)

	row, err := repo.UpdateUser(uRequest, u.ID)
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

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(u.ID).WillReturnResult(sqlmock.NewResult(0, 1)).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.DeleteUser(int(u.ID))
	require.NoError(t, err)
}
