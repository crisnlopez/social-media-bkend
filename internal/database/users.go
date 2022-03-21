package database

import (
	"database/sql"
	"fmt"
)

// User -
type User struct {
	Email     string    `json:"email"`
	Pass      string    `json:"pass"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

func (d DB) CreateUser(email, pass, name string, age int) (User, error) {
  // Need check if user exist
  var col string

  row := d.Db.QueryRow("select email from users where email = ?", email)
  err := row.Scan(&col)
  if err != nil {
    if err != sql.ErrNoRows {
      return User{}, fmt.Errorf("User with %s email already exists", email)
    }
  }

	// New user instance
	user := User{
		Email:     email,
		Pass:      pass,
		Name:      name,
		Age:       age,
	}

  // Create User in Database
  _, err = d.Db.Exec("insert into users (email, pass, name, age) values (?, ?, ?, ?)",user.Email, user.Pass, user.Name, user.Age)
  if err != nil {
    return User{}, fmt.Errorf("Error creating user: %v\n", err)
  }

	return user, nil
}

func (d DB) UpdateUser(email, pass, name string, age int) (User, error) {
 // Need check if user exist
  var col string

  row := d.Db.QueryRow("select email from users where email = ?", email)
  err := row.Scan(&col)
  if err != nil {
    if err == sql.ErrNoRows {
      return User{}, fmt.Errorf("User with %s email doesn't exists", email)
    } else {
      return User{}, err
    }
  }

  // Update user
  _, err = d.Db.Exec("update users set pass = ?, name = ?, age = ? where email = ?", pass, name, age, email)
  if err != nil {
    return User{}, err
  }

  return d.GetUser(email) // Need to return User update
}

func (d DB) GetUser(email string) (User, error) {
  user := User{}

  rows, err := d.Db.Query("select * from users where email = ?", email)
  if err != nil {
    return User{}, err
  }

  defer rows.Close()
  for rows.Next() {
    if err := rows.Scan(&user.Email, &user.Pass, &user.Name, &user.Age); err != nil {
      return User{}, fmt.Errorf("Get user: %q, %v\n", email, err)
    }

    if err = rows.Err(); err != nil {
      return User{}, fmt.Errorf("Get user: %q, %v\n", email, err)
    }
  }

  return user, nil
}

func (d DB) DeleteUser(email string) error {
  _, err := d.Db.Exec("delete from users where email = ?", email)
  if err != nil {
    return err
  }

  return nil
}
