package handler_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/crisnlopez/social-media-bkend/internal/user/handler"
	user "github.com/crisnlopez/social-media-bkend/internal/user/models"
	"github.com/crisnlopez/social-media-bkend/internal/util"
	"github.com/julienschmidt/httprouter"

	"github.com/crisnlopez/social-media-bkend/internal/user/gateway/mocks"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/require"
)

func createUserRequest() *user.UserRequest {
	return &user.UserRequest{
		Email:     util.RandomEmail(),
		Pass:      util.RandomPass(),
		Nick:      util.RandomNick(),
		Name:      util.RandomName(),
		Age:       util.RandomAge(),
		CreatedAt: time.Now().UTC(),
	}
}

func newUser(u *user.UserRequest) *user.User {
	return &user.User{
		ID:        util.RandomInt(0, 100),
		Email:     u.Email,
		Pass:      u.Pass,
		Nick:      u.Nick,
		Name:      u.Name,
		Age:       u.Age,
		CreatedAt: u.CreatedAt,
	}
}

func TestCreateUser(t *testing.T) {
	userReq := createUserRequest()
	id := util.RandomInt(1, 100)

	testCases := []struct {
		name          string
		mockActions   func(gtwMock *mocks.MockUserGateway)
		checkResponse func(rr *httptest.ResponseRecorder)
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
					Return(id, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
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
		{
			name: "Internal Error",
			mockActions: func(gtwMock *mocks.MockUserGateway) {
				gtwMock.EXPECT().
					GetUserEmail(userReq.Email).
					Times(1).
					Return(false, errors.New("Internal Server Error"))
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
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

			data, err := json.Marshal(userReq)
			require.NoError(t, err)

			url := "/users"

			req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			rr := httptest.NewRecorder()

			testHandler := &handler.UserHandler{Gtw: gtwMock}
			testHandler.CreateUser(rr, req, httprouter.Params{})

			tc.checkResponse(rr)
		})
	}
}

func TestGetUser(t *testing.T) {
	userReq := createUserRequest()
	userRes := newUser(userReq)

	testCases := []struct {
		name          string
		mockActions   func(gtwMock *mocks.MockUserGateway)
		checkResponse func(rr *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			mockActions: func(gtwMock *mocks.MockUserGateway) {
				gtwMock.EXPECT().
					GetUser(userRes.ID).
					Times(1).
					Return(userRes, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "User Doesn't Exists",
			mockActions: func(gtwMock *mocks.MockUserGateway) {
				gtwMock.EXPECT().
					GetUser(gomock.Eq(userRes.ID)).
					Times(1).
					Return(nil, sql.ErrNoRows)
			},
			checkResponse: func(rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, rr.Code)
			},
		},
		{
			name: "Internal Error",
			mockActions: func(gtwMock *mocks.MockUserGateway) {
				gtwMock.EXPECT().
					GetUser(userRes.ID).
					Times(1).
					Return(nil, errors.New("Internal Server Error"))
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
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

			rr := httptest.NewRecorder()

			url := "/users/:id"

			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			testHandler := &handler.UserHandler{Gtw: gtwMock}
			testHandler.GetUser(rr, req, []httprouter.Param{{
				Key:   "id",
				Value: strconv.Itoa(int(userRes.ID)),
			}})

			tc.checkResponse(rr)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	userReq := createUserRequest()
	id := util.RandomInt(1,100)

	testCases := []struct {
		name          string
		mockActions   func(gtwMock *mocks.MockUserGateway)
		checkResponse func(rr *httptest.ResponseRecorder)
	}{
		{
			name: "User Updated",
			mockActions: func(gtwMock *mocks.MockUserGateway) {
				gtwMock.EXPECT().
				  UpdateUser(userReq, id).
					Times(1).
					Return(int64(1), nil)
			},
			checkResponse: func(rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rr.Code)
			},
		},
		{
			name: "User Doesn't Exists",
			mockActions: func(gtwMock *mocks.MockUserGateway) {
				gtwMock.EXPECT().
				  UpdateUser(userReq, id).
					Times(1).
					Return(int64(0), errors.New("Unexpected Error"))
			},
			checkResponse: func(rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, rr.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			gtwMock := mocks.NewMockUserGateway(mockCtrl)
			tc.mockActions(gtwMock)

			data, err := json.Marshal(userReq)
			require.NoError(t, err)

			rr := httptest.NewRecorder()

			url := "/users/:id"

			req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			testHandler := &handler.UserHandler{Gtw: gtwMock}
			testHandler.UpdateUser(rr, req, []httprouter.Param{{
				Key: "id",
				Value: strconv.Itoa(int(id)),
			}})

			tc.checkResponse(rr)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	id := util.RandomInt(0,100)

	testCases := []struct {
		name string
		mockActions func(gtwMock *mocks.MockUserGateway)
		checkResponse func(rr *httptest.ResponseRecorder)
	}{
		{
			name: "Delete User",
			mockActions: func(gtwMock *mocks.MockUserGateway) {
				gtwMock.EXPECT().
				  DeleteUser(int(id)).
					Times(1).
					Return(nil)
		  },
			checkResponse: func(rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rr.Code)
			},
	  },
		{
			name: "Internal Error",
			mockActions: func(gtwMock *mocks.MockUserGateway) {
				gtwMock.EXPECT().
				  DeleteUser(int(id)).
					Times(1).
					Return(errors.New("User doesn't exists"))
		  },
			checkResponse: func(rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, rr.Code)
			},
	  },
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			gtwMock := mocks.NewMockUserGateway(mockCtrl)
			tc.mockActions(gtwMock)

			rr := httptest.NewRecorder()

			url := "/users/:id"

			req, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			testHandler := &handler.UserHandler{Gtw: gtwMock}
			testHandler.DeleteUser(rr, req, []httprouter.Param{{
				Key: "id",
				Value: strconv.Itoa(int(id)),
			}})

			tc.checkResponse(rr)
		})
	}
}
