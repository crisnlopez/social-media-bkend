package mysql

import (
	"database/sql"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/crisnlopez/social-media-bkend/internal/user/models"
	"github.com/crisnlopez/social-media-bkend/internal/util"
)

var u = &user.User{
  ID: util.RandomInt(0,1000),
  Email: util.RandomEmail(),
  Pass: util.RandomPass(),
  Nick: util.RandomNick(),
  Name: util.RandomName(),
  Age: util.RandomAge(),
  CreatedAt: time.Now().UTC().Local(),
}

var uRequest = &user.UserRequest{
  Email: util.RandomEmail(),
  Pass: util.RandomPass(),
  Nick: util.RandomNick(),
  Name: util.RandomEmail(),
  Age: util.RandomAge(),
  CreatedAt: time.Now().UTC().Local(),
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
  db, mock, err := sqlmock.New()
  if err != nil {
    log.Fatalf("Error: %s", err)
  }

  return db, mock
}

func TestCreateUser(t *testing.T) {
  db, mock := NewMock()
  repo := &UserQueries{db}
  defer func() {
    repo.db.Close()
  }()

  query := `INSERT INTO users (email, pass, name, age, nick, created_at) VALUES (?, ?, ?, ?, ?, ?)`

  mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(uRequest.Email,uRequest.Pass,uRequest.Name,uRequest.Age,uRequest.Nick,uRequest.CreatedAt).WillReturnResult(sqlmock.NewResult(1,1))

  user, err := repo.CreateUser(uRequest)
  assert.NoError(t,err)
  assert.NotEmpty(t,user)
} 

func TestGetUser(t *testing.T) {
  db, mock := NewMock()
  repo := &UserQueries{db}
  defer func() {
    repo.db.Close()
  }()

  query :=`SELECT * FROM users WHERE id = ?` 

  rows := sqlmock.NewRows([]string{"id","email","pass","nick","name","age","created_at"}).AddRow(u.ID,u.Email,u.Pass,u.Nick,u.Name,u.Age,u.CreatedAt)

  mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(u.ID).WillReturnRows(rows)

  user, err := repo.GetUser(u.ID)
  assert.NoError(t, err)
  assert.NotEmpty(t, user)
}

func TestGetUserEmail(t *testing.T) {
  db, mock := NewMock()
  repo := &UserQueries{db}
  defer func() {
    repo.db.Close()
  }()

  query := `SELECT email FROM users WHERE email = ?`

  row := mock.NewRows([]string{"email"}).AddRow(u.Email)

  mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(u.Email).WillReturnRows(row)

  ok, err := repo.GetUserEmail(u.Email)
  require.NoError(t, err)
  require.True(t, ok)

  ok, err = repo.GetUserEmail("falseemail@email.com")
  require.Error(t, err)
  require.False(t, ok)
}

func TestUpdateUser(t *testing.T) {
  db, mock := NewMock()
  repo := &UserQueries{db}
  defer func() {
    repo.db.Close()
  }()

  query := `UPDATE users SET email = ?, pass = ?, nick= ?,  name= ?, age= ? WHERE id = ?`

  mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(uRequest.Email,uRequest.Pass,uRequest.Nick,uRequest.Name,uRequest.Age,u.ID).WillReturnResult(sqlmock.NewResult(0,1))
}

func TestDeleteUser(t *testing.T) {
  db, mock := NewMock()
  repo := &UserQueries{db}
  defer func() {
    repo.db.Close()
  }()

  query := `DELETE FROM users WHERE id = ?`

  mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(u.ID).WillReturnResult(sqlmock.NewResult(0, 1)).WillReturnResult(sqlmock.NewResult(0,1))
}
