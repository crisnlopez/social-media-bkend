package gateway

import (
  "database/sql"

  "github.com/crisnlopez/social-media-bkend/internal/user/models"
  repo "github.com/crisnlopez/social-media-bkend/internal/user/repository"
)

type UserGateway interface {
	CreateUser(u *user.UserRequest) (*user.User, error)
	GetUser(id int64) (*user.User, error)
	GetUserEmail(email string) (bool, error)
  UpdateUser(u *user.UserRequest, id int64) (int64, error)
  DeleteUser(id int) (error)
}

type UserInRepo struct {
	repo.Repository
}

func NewGateway(db *sql.DB) UserGateway {
	return &UserInRepo{repo.NewRepository(db)}
}

func (r *UserInRepo) CreateUser(newUser *user.UserRequest) (*user.User, error) {
	return r.CreateUser(newUser)
}

func (r *UserInRepo) GetUser(id int64) (*user.User, error) {
	return r.GetUser(id)
}

func (r *UserInRepo) GetUserEmail(email string) (bool, error) {
	return r.GetUserEmail(email)
}

func (r *UserInRepo) UpdateUser(u *user.UserRequest, id int64) (int64, error) {
	return r.UpdateUser(u, id)
}

func (r *UserInRepo) DeleteUser(id int) (error) {
  return r.DeleteUser(id)
}
