package mysql

import (
	"database/sql"
	"errors"
	"log"

	"github.com/crisnlopez/social-media-bkend/internal/user/models"
  "github.com/crisnlopez/social-media-bkend/internal/user/repository"
)

type UserQueries struct{
  db *sql.DB
}

func NewRepository(db *sql.DB) repository.Repository{
  return &UserQueries{db:db}
}

func (r *UserQueries) CreateUser(u *user.UserRequest) (*user.User, error) {
  // Create User
  result, err := r.db.Exec("INSERT INTO users (email, pass, name, age, nick, created_at) VALUES (?, ?, ?, ?, ?, ?)", u.Email, u.Pass, u.Name, u.Age, u.Nick, u.CreatedAt)
  if err != nil {
    log.Printf("cannot save New User, %s\n",err.Error())
    return nil, err
  }

  // Get UserID
  id, err := result.LastInsertId()
  if err != nil {
    return nil, err
  }

  return &user.User{
    ID: id,
    Email: u.Email,
    Pass: u.Pass,
    Nick: u.Nick,
    Name: u.Name,
    Age: u.Age,
    CreatedAt: u.CreatedAt,
  }, nil
}

func (r *UserQueries) GetUser(id int64) (*user.User, error) {
  u := user.User{}
  // Getting User
  if err := r.db.QueryRow(`SELECT * FROM users WHERE id = ?`, id).Scan(&u.ID, &u.Email, &u.Pass, &u.Nick, &u.Name, &u.Age, &u.CreatedAt); err != nil {
    if err == sql.ErrNoRows {
      return &user.User{}, sql.ErrNoRows
    } else {
      return &user.User{}, err
    }
  }

  return &u, nil
}

func (r *UserQueries) GetUserEmail(email string) (bool, error){
  // Check if user already exists
  var col string
  row := r.db.QueryRow("SELECT email FROM users WHERE email = ?",email)
  if err := row.Scan(&col); err == sql.ErrNoRows {
    return false, errors.New("User doesn't exists")
  } else if err != nil {
    return false, err
  }
  return true, nil
}

func (r *UserQueries) UpdateUser(u *user.UserRequest, id int64) (*user.User, error) {
  result, err := r.db.Exec("UPDATE users SET email = ?, pass = ?, name= ?,  age= ?, nick= ? WHERE id = ?", u.Email, u.Pass, u.Name, u.Age, u.Nick, id) 
  if err != nil {
    return &user.User{}, err
  }

  rows, err := result.RowsAffected()
  if err != nil {
    return &user.User{}, err
  }
  log.Printf("Number of rows affected: %v\n", rows)

  usr, err := r.GetUser(id)
  if err != nil {
    return &user.User{}, err
  }
  
  return usr, nil
}

func (r *UserQueries) DeleteUser(id int) error {
  result, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
  if err != nil {
    return err
  }
  
  rows, err := result.RowsAffected()
  if err != nil {
    return err
  }
  log.Printf("Rows affected %v\n",rows)

  return nil
}
