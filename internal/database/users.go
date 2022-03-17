package database

import (
	"fmt"
	"time"
)

// User -
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Pass  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

func (c Client) CreateUser(email, password, name string, age int) (User, error) {
  // Need check if user exist

	// New user instance
	user := User{
		CreatedAt: time.Now().UTC().Local(),
		Email:     email,
		Pass:  password,
		Name:      name,
		Age:       age,
	}

  _, err := c.DB.Exec("insert into users (email, password, name, age) values (?, ?, ?, ?)", user.Email, user.Pass, user.Name, user.Age)
  if err != nil {
    return User{}, fmt.Errorf("Error creating user: %v", err)
  }

	return user, nil
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
  result, err := c.DB.Exec("update users set pass = ?, name = ?, age = ? where email = ?", password, name, age, email)
  if err != nil {
    return User{}, err
  }

  result.RowsAffected() 

  return c.GetUser(email) // Need to return User update
}

func (c Client) GetUser(email string) (User, error) {
  user := User{}

  rows, err := c.DB.Query("select * from users where email = ?", email)
  if err != nil {
    return User{}, err
  }

  defer rows.Close()
  for rows.Next() {
    if err := rows.Scan(&user.CreatedAt, &user.Email, &user.Pass, &user.Name, &user.Age); err != nil {
      return User{}, fmt.Errorf("Get user: %q, %v", email, err)
    }

    if err = rows.Err(); err != nil {
      return User{}, fmt.Errorf("Get user: %q, %v", email, err)
    }
  }

  return user, nil
}

func (c Client) DeleteUser(email string) error {
  result, err := c.DB.Exec("delete from users where email = ?", email)
  if err != nil {
    return err
  }

  fmt.Printf("User deleted, printing result: %s", result)

  return nil
}
