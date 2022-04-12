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
	userReq := createUserRequest()
	userExpected := newUser(userReq)

	testCases := []struct {
		name string
		mockActions func(gtwMock *mocks.MockUserGateway)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			mockActions: func(gtwMock *mocks.MockUserGateway) {
				gtwMock.EXPECT().
				GetUserEmail(userReq.Email).
				Times(1).
				Return(false, nil)

				gtwMock.EXPECT().
				CreateUser(userReq).
				Times(1).
				Return(userExpected, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Bad Request",
			mockActions: func(gtwMock *mocks.MockUserGateway) {
				gtwMock.EXPECT().
				GetUserEmail(userReq.Email).
				Times(1).
				Return(true, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			mockCrtl := gomock.NewController(t)
			defer mockCrtl.Finish()

			gtwMock := mocks.NewMockUserGateway(mockCrtl)
			tc.mockActions(gtwMock)

			recorder := httptest.NewRecorder()

			data, err := json.Marshal(userReq)
			require.NoError(t, err)

			url := "/users"

			req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			testHandler := &handler.UserHandler{Gtw: gtwMock}
			testHandler.CreateUser(recorder, req, httprouter.Params{})

			tc.checkResponse(recorder)
		})
	}
}
