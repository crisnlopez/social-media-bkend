package gateway

import (
  "database/sql"

  "github.com/crisnlopez/social-media-bkend/internal/user/models"
  "github.com/crisnlopez/social-media-bkend/internal/user/repository"
)

type UserGateway interface {
	CreateUser(u *user.UserRequest) (*user.User, error)
	GetUser(id int64) (*user.User, error)
	GetUserEmail(email string) (bool, error)
  UpdateUser(u *user.UserRequest, id int64) (*user.User, error)
  DeleteUser(id int) (error)
}

type UserInRepo struct {
	repo repository.Repository
}

func NewGateway(db *sql.DB) UserGateway {
	return &UserInRepo{repository.NewRepository(db)}
}

func (r *UserInRepo) CreateUser(newUser *user.UserRequest) (*user.User, error) {
	return r.repo.CreateUser(newUser)
}

func (r *UserInRepo) GetUser(id int64) (*user.User, error) {
	return r.repo.GetUser(id)
}

func (r *UserInRepo) GetUserEmail(email string) (bool, error) {
	return r.repo.GetUserEmail(email)
}

func (r *UserInRepo) UpdateUser(u *user.UserRequest, id int64) (*user.User, error) {
	return r.repo.UpdateUser(u, id)
}

func (r *UserInRepo) DeleteUser(id int) (error) {
  return r.repo.DeleteUser(id)
}
