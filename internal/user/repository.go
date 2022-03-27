package user

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type Repository interface{
  createUser(u *UserRequest) (*User, error)
  getUser(id int) (*User, error)
  getUserEmail(email string) (bool, error)
  updateUser(u *UserRequest, id int) (*User, error)
  deleteUser(id int) (error)
}

type UserRepository struct{
  db *sql.DB
}

func NewRepository(db *sql.DB) Repository{
  return &UserRepository{db:db}
}

func (r *UserRepository) createUser(u *UserRequest) (*User, error) {
// Create User
  result, err := r.db.Exec("INSERT INTO users (email, pass, user_nick, user_name, age) VALUES (?, ?, ?, ?, ?)", u.Email, u.Pass, u.Nick, u.Name, u.Age)
  if err != nil {
    log.Printf("cannot save New User, %s\n",err.Error())
    return nil, err
  }

  // Get UserID
  id, err := result.LastInsertId()
  if err != nil {
    return nil, err
  }

  return &User{
    ID: int(id),
    Email: u.Email,
    Pass: u.Pass,
    Nick: u.Nick,
    Name: u.Name,
    Age: u.Age,
    CreatedAt: time.Now(),
  }, nil
}

func (r *UserRepository) getUser(id int) (*User, error) {
  user := User{}
  // Getting User
  if err := r.db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Email, &user.Pass, &user.Nick, &user.Name, &user.Age, &user.CreatedAt); err != nil {
    if err == sql.ErrNoRows { // If user doesn't exist
      return &User{}, errors.New("User doesn't exists")
    } else {
      return &User{}, err
    }
  }

  return &user, nil
}

func (r *UserRepository) getUserEmail(email string) (bool, error){
  // Check if user already exists
  var col string
  row := r.db.QueryRow("SELECT email FROM users WHERE email = ?",email)
  if err := row.Scan(&col); err != sql.ErrNoRows {
    return false, err
  }
  return true, nil
}

func (r *UserRepository) updateUser(u *UserRequest, id int) (*User, error) {
  result, err := r.db.Exec("UPDATE users SET email = ?, pass = ?, nick = ?,  name= ?, age = ? WHERE id = ?", u.Pass, u.Name, u.Age, u.Nick, u.Email, id) 
  if err != nil {
    return &User{}, err
  }

  rows, err := result.RowsAffected()
  if err != nil {
    return &User{}, err
  }
  log.Printf("Number of rows affected: %v\n", rows)

  user, err := r.getUser(id) // Return user updated??
  if err != nil {
    return &User{}, err
  }
  
  return user, nil
}

func (r *UserRepository) deleteUser(id int) (error) {
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
