package repository

import (
	"github.com/crisnlopez/social-media-bkend/internal/user/models"
)

type Repository interface{
  CreateUser(u *user.UserRequest) (*user.User, error)
  GetUser(id int64) (*user.User, error)
  GetUserEmail(email string) (bool, error)
  UpdateUser(u *user.UserRequest, id int64) (*user.User, error)
  DeleteUser(id int) (error)
}
