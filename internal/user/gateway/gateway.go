package user

import (
  "database/sql"

  "github.com/crisnlopez/social-media-bkend/internal/user/models"
)

type UserGateway interface {
	CreateUser(u *user.UserRequest) (*user.User, error)
	GetUser(id int64) (*user.User, error)
	GetUserEmail(email string) (bool, error)
  UpdateUser(u *user.UserRequest, id int64) (*user.User, error)
  DeleteUser(id int) (error)
}

type UserInRepo struct {
	repo Repository
}

func NewGateway(db *sql.DB) UserGateway {
	return &UserInRepo{NewRepository(db)}
}

func (r *UserInRepo) CreateUser(newUser *user.UserRequest) (*user.User, error) {
	return r.repo.createUser(newUser)
}

func (r *UserInRepo) GetUser(id int64) (*user.User, error) {
	return r.repo.getUser(id)
}

func (r *UserInRepo) GetUserEmail(email string) (bool, error) {
	return r.repo.getUserEmail(email)
}

func (r *UserInRepo) UpdateUser(u *user.UserRequest, id int64) (*user.User, error) {
	return r.repo.updateUser(u, id)
}

func (r *UserInRepo) DeleteUser(id int) (error) {
  return r.repo.deleteUser(id)
}
