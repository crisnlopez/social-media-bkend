package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/crisnlopez/social-media-bkend/internal/user/models"
)

type Repository interface {
	createUser(u *user.UserRequest) (int64, error)
	getUser(id int64) (*user.User, error)
	getUserEmail(email string) (bool, error)
	updateUser(u *user.UserRequest, id int64) (int64, error)
	deleteUser(id int) error
}

type userQueries struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &userQueries{db: db}
}

func (r *userQueries) createUser(u *user.UserRequest) (int64, error) {
	// Create User
	result, err := r.db.Exec("INSERT INTO users (email, pass, name, age, nick, created_at) VALUES (?, ?, ?, ?, ?, ?)", u.Email, u.Pass, u.Name, u.Age, u.Nick, u.CreatedAt)
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

func (r *userQueries) getUser(id int64) (*user.User, error) {
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

func (r *userQueries) getUserEmail(email string) (bool, error) {
	// Check if user already exists
	var col string
	row := r.db.QueryRow("SELECT email FROM users WHERE email = ?", email)
	if err := row.Scan(&col); err == sql.ErrNoRows {
		return false, errors.New("User doesn't exists")
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (r *userQueries) updateUser(u *user.UserRequest, id int64) (int64, error) {
	result, err := r.db.Exec(`UPDATE users SET email = ?, pass = ?, name= ?,  age= ?, nick= ? WHERE id = ?`, u.Email, u.Pass, u.Name, u.Age, u.Nick, id)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (r *userQueries) deleteUser(id int) error {
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
