package user

import "database/sql"

type UserGateway interface {
	CreateUser(u *UserRequest) (*User, error)
	GetUser(id int) (*User, error)
	GetUserEmail(email string) (bool, error)
	UpdateUser(u *UserRequest, id int) (*User, error)
  DeleteUser(id int) (error)
}

type UserInRepo struct {
	repo Repository
}

func NewGateway(db *sql.DB) UserGateway {
	return &UserInRepo{NewRepository(db)}
}

func (r *UserInRepo) CreateUser(newUser *UserRequest) (*User, error) {
	return r.repo.createUser(newUser)
}

func (r *UserInRepo) GetUser(id int) (*User, error) {
	return r.repo.getUser(id)
}

func (r *UserInRepo) GetUserEmail(email string) (bool, error) {
	return r.repo.getUserEmail(email)
}

func (r *UserInRepo) UpdateUser(u *UserRequest, id int) (*User, error) {
	return r.repo.updateUser(u, id)
}

func (r *UserInRepo) DeleteUser(id int) (error) {
  return r.repo.deleteUser(id)
}
