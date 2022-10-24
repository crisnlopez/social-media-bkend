package sql

import (
	"database/sql"
	"log"

	"github.com/crisnlopez/social-media-bkend/internal/user"
)

type UserQueries struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserQueries {
	return &UserQueries{db: db}
}

func (r *UserQueries) CreateUser(u user.UserRequest) (int64, error) {
	// Create User
	result, err := r.db.Exec("INSERT INTO users (email, pass, name, age, nick) VALUES (?, ?, ?, ?, ?)", u.Email, u.Pass, u.Name, u.Age, u.Nick)
	if err != nil {
		log.Printf("cannot save New User, %s\n", err.Error())
		return 0, err
	}

	// Get UserID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserQueries) GetUser(id int64) (user.User, error) {
	var u user.User
	// Getting User
	if err := r.db.QueryRow(`SELECT * FROM users WHERE id = ?`, id).Scan(&u.ID, &u.Email, &u.Pass, &u.Nick, &u.Name, &u.Age, &u.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return user.User{}, sql.ErrNoRows
		} else {
			return user.User{}, err
		}
	}
	return u, nil
}

func (r *UserQueries) UpdateUser(u user.UserRequest, id int64) (int64, error) {
	result, err := r.db.Exec(`UPDATE users SET email = ?, pass = ?, name= ?,  age= ?, nick= ? WHERE id = ?`, u.Email, u.Pass, u.Name, u.Age, u.Nick, id)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return rows, err
	}

	return rows, nil
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
	log.Printf("Rows affected %v\n", rows)

	return nil
}
