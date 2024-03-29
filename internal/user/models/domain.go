package user

import (
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Pass      string    `json:"pass"`
	Nick      string    `json:"nick"`
	Name      string    `json:"name"`
	Age       int64     `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRequest struct {
	Email string `json:"email" validate:"required,email"`
	Pass  string `json:"pass" validate:"required"`
	Nick  string `json:"nick" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Age   int64  `json:"age" validate:"required"`
}
