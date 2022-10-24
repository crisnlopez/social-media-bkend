package user

import (
	"database/sql"
)

//go:generate mockgen -destination=mockgateway.go -package=user github.com/crisnlopez/social-media-bkend/internal/user Gateway

type Gateway interface {
	CreateUser(u UserRequest) (int64, error)
	GetUser(id int64) (User, error)
	UpdateUser(u UserRequest, id int64) (int64, error)
	DeleteUser(id int) error
}

type UserInRepo struct {
	Repository
}

func NewGateway(db *sql.DB) Gateway {
	return &UserInRepo{NewRepository(db)}
}

func (r *UserInRepo) CreateUser(newUser UserRequest) (int64, error) {
	return r.CreateUser(newUser)
}

func (r *UserInRepo) GetUser(id int64) (User, error) {
	return r.GetUser(id)
}

func (r *UserInRepo) UpdateUser(u UserRequest, id int64) (int64, error) {
	return r.UpdateUser(u, id)
}

func (r *UserInRepo) DeleteUser(id int) error {
	return r.DeleteUser(id)
}
