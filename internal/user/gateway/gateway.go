package gateway

import (
	"database/sql"

	"github.com/crisnlopez/social-media-bkend/internal/user/models"
	repo "github.com/crisnlopez/social-media-bkend/internal/user/repository"
)

//go:generate mockgen -destination=mocks/mock_UserGateway.go -package=mocks github.com/crisnlopez/social-media-bkend/internal/user/gateway UserGateway

type UserGateway interface {
	CreateUser(u *user.UserRequest) (int64, error)
	GetUser(id int64) (*user.User, error)
	UpdateUser(u *user.UserRequest, id int64) (int64, error)
	DeleteUser(id int) error
}

type UserInRepo struct {
	repo.Repository
}

func NewGateway(db *sql.DB) UserGateway {
	return &UserInRepo{repo.NewRepository(db)}
}

func (r *UserInRepo) CreateUser(newUser *user.UserRequest) (int64, error) {
	return r.CreateUser(newUser)
}

func (r *UserInRepo) GetUser(id int64) (*user.User, error) {
	return r.GetUser(id)
}

func (r *UserInRepo) UpdateUser(u *user.UserRequest, id int64) (int64, error) {
	return r.UpdateUser(u, id)
}

func (r *UserInRepo) DeleteUser(id int) error {
	return r.DeleteUser(id)
}
