package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/crisnlopez/social-media-bkend/internal/user/handler"
	"github.com/crisnlopez/social-media-bkend/internal/user/models"
	"github.com/crisnlopez/social-media-bkend/internal/util"
	"github.com/julienschmidt/httprouter"

	"github.com/crisnlopez/social-media-bkend/internal/user/gateway/mocks"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/require"
)

func createUserRequest() *user.UserRequest {
  return &user.UserRequest{
    Email: util.RandomEmail(),
    Pass: util.RandomPass(),
    Nick: util.RandomNick(),
    Name: util.RandomName(),
    Age: util.RandomAge(),
    CreatedAt: time.Now().UTC(),
  }
}

func newUser(u *user.UserRequest) *user.User {
  return &user.User{
    ID: util.RandomInt(0,100),
    Email: u.Email,
    Pass: u.Pass,
    Nick: u.Nick,
    Name: u.Name,
    Age: u.Age,
    CreatedAt: u.CreatedAt,
  }
}

func TestCreateUserHandler(t *testing.T) {
  mockCrtl := gomock.NewController(t)
  defer mockCrtl.Finish()

  mockUserGateway := mocks.NewMockUserGateway(mockCrtl)  
  testHandler := &handler.UserHandler{Gtw: mockUserGateway}

  userReq := createUserRequest()
  userExpected := newUser(userReq)
  
  // Mock Calls
  mockUserGateway.EXPECT().GetUserEmail(userReq.Email).Return(false, nil).Times(1)
  mockUserGateway.EXPECT().CreateUser(userReq).Return(userExpected, nil).Times(1)

  // Prepare request
  data, err := json.Marshal(userReq)
  require.NoError(t, err)

  url := "/users"

  req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
  require.NoError(t, err)
  require.NotEmpty(t, req)

  recorder := httptest.NewRecorder()
  testHandler.CreateUser(recorder, req, httprouter.Params{})

  // Check response
  require.NotEqual(t, http.StatusInternalServerError, recorder.Code)
  require.NotEqual(t, http.StatusBadRequest, recorder.Code)
  require.Equal(t, http.StatusOK, recorder.Code)
}
