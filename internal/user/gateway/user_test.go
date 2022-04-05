package user

import (
	"database/sql"
	"testing"
	"time"

	"github.com/crisnlopez/social-media-bkend/internal/user/models"
	"github.com/crisnlopez/social-media-bkend/internal/util"

	"github.com/stretchr/testify/require"
)

var userRequest = user.UserRequest{
   Email: util.RandomEmail(),
   Pass: util.RandomPass(),
   Nick: util.RandomNick(),
   Name: util.RandomName(),
   Age: util.RandomAge(),
}

func createUser(t *testing.T) *user.User{
   user, err := testQueries.createUser(&userRequest)

   require.NoError(t, err)
   require.NotEmpty(t, user)

   require.Equal(t, userRequest.Email, user.Email)
   require.Equal(t, userRequest.Pass, user.Pass)
   require.Equal(t, userRequest.Nick, user.Nick)
   require.Equal(t, userRequest.Name, user.Name)
   require.Equal(t, userRequest.Age, user.Age)

   require.NotZero(t, user.ID)
   require.NotZero(t, user.CreatedAt)

   return user
}

func TestCreateUser(t *testing.T) {
   createUser(t)
}

func TestGetUser(t *testing.T) {
   newUser := createUser(t)
   getUser, err := testQueries.getUser(newUser.ID)
  
   require.NoError(t, err)
   require.NotEmpty(t, getUser)

   require.Equal(t, newUser.ID, getUser.ID)
   require.Equal(t, newUser.Email, getUser.Email)
   require.Equal(t, newUser.Nick, getUser.Nick)
   require.Equal(t, newUser.Age, getUser.Age)
   require.WithinDuration(t, newUser.CreatedAt, getUser.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
   newUser := createUser(t)

   arg := user.UserRequest{
     Email: util.RandomEmail(),
     Pass: util.RandomPass(),
     Nick: util.RandomNick(),
     Name: util.RandomName(),
     Age: util.RandomAge(),
   }

   updateUser, err := testQueries.updateUser(&arg, newUser.ID)

   require.NoError(t,err)
   require.NotEmpty(t,updateUser)

   require.Equal(t, newUser.ID, updateUser.ID)
   require.Equal(t, arg.Email, updateUser.Email)
   require.Equal(t, arg.Pass, updateUser.Pass)
   require.Equal(t, arg.Nick, updateUser.Nick)
   require.Equal(t, arg.Name, updateUser.Name)
   require.Equal(t, arg.Age, updateUser.Age)
   require.WithinDuration(t, newUser.CreatedAt, updateUser.CreatedAt, time.Second)
}

func TestGetUserEmail(t *testing.T) {
   newUser := createUser(t)
   okEmail, err := testQueries.getUserEmail(newUser.Email)

   require.NoError(t, err)
   require.True(t, okEmail)

   dontExist, err := testQueries.getUserEmail("Email_dontExist@dontExist.com")
   require.Error(t, err)
   require.False(t, dontExist)
}

func TestDeleteUser(t *testing.T) {
   newUser := createUser(t)

   err := testQueries.deleteUser(int(newUser.ID))
   require.NoError(t, err)

   dontExist, err := testQueries.getUser(newUser.ID)
   require.Error(t, err)
   require.EqualError(t, err, sql.ErrNoRows.Error())
   require.Empty(t, dontExist)
}
