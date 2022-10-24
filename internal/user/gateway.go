package user

import ()

//go:generate mockgen -destination=gateway_mock.go -package=user github.com/crisnlopez/social-media-bkend/internal/user Gateway

type Repository interface {
	CreateUser(u UserRequest) (int64, error)
	GetUser(id int64) (User, error)
	UpdateUser(u UserRequest, id int64) (int64, error)
	DeleteUser(id int) error
}

type gateway struct {
	repository Repository
}

func NewGateway() *gateway {
	return &gateway{}
}

func (r *gateway) CreateUser(newUser UserRequest) (int64, error) {
	return r.CreateUser(newUser)
}

func (r *gateway) GetUser(id int64) (User, error) {
	return r.GetUser(id)
}

func (r *gateway) UpdateUser(u UserRequest, id int64) (int64, error) {
	return r.UpdateUser(u, id)
}

func (r *gateway) DeleteUser(id int) error {
	return r.DeleteUser(id)
}
