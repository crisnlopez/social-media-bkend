package user

import (
	"database/sql"
)

type Repository interface{
 GetUserEmail(u User) (bool, error)
}

type repository struct{
  db *sql.DB
}

func newRepository(db *sql.DB) Repository{
  return &repository{
    db:db,
  }
}

func (r *repository) GetUserEmail(u User) (bool, error){
  // Check if user already exists
  var email string
  row := r.db.QueryRow("SELECT email FROM users WHERE email = ?", u.Email)
  if err := row.Scan(&email); err != sql.ErrNoRows {
    return false, err
  }
  return true, nil
}
