package user

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	userReq := createUserRequest()
	id := RandomInt(1, 100)

	testCases := []struct {
		name          string
		input         []byte
		mockActions   func(gtwMock MockGateway)
		checkResponse func(rr *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			input: func() []byte {
				b, _ := json.Marshal(&userReq)
				return b
			}(),
			mockActions: func(gtwMock MockGateway) {
				gtwMock.EXPECT().
					CreateUser(userReq).
					Times(1).
					Return(id, nil)
			},
			checkResponse: func(rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, rr.Code)
			},
		},
		{
			name:        "Invalid Json",
			input:       []byte(`{"email":"test@test.com","pass":"12354","nick":"nicktest","name":"nametest","age":20`), // missing "}" to be a valid json
			mockActions: nil,
			checkResponse: func(rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, rr.Code)
			},
		},
		{
			name:        "Bad Request",
			input:       []byte(`{"email":"test@test.com"}`), // All files required
			mockActions: nil,
			checkResponse: func(rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rr.Code)
			},
		},
		{
			name: "Internal Error",
			input: func() []byte {
				b, _ := json.Marshal(&userReq)
				return b
			}(),
			mockActions: func(gtwMock MockGateway) {
				gtwMock.EXPECT().
					CreateUser(userReq).
					Return(int64(0), errors.New("Internal Server Error"))
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

			gtwMock := NewMockGateway(mockCrtl)
			if tc.mockActions != nil {
				tc.mockActions(*gtwMock)
			}

			url := "/users"

			req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(tc.input))
			require.NoError(t, err)

			rr := httptest.NewRecorder()

			testHandler := handler{gtw: gtwMock}
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
		mockActions   func(gtwMock MockGateway)
		checkResponse func(rr *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			mockActions: func(gtwMock MockGateway) {
				gtwMock.EXPECT().
					GetUser(userRes.ID).
					Times(1).
					Return(userRes, nil)
			},
			checkResponse: func(rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rr.Code)
			},
		},
		{
			name: "User Doesn't Exists",
			mockActions: func(gtwMock MockGateway) {
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
			mockActions: func(gtwMock MockGateway) {
				gtwMock.EXPECT().
					GetUser(userRes.ID).
					Times(1).
					Return(nil, errors.New("Internal Server Error"))
			},
			checkResponse: func(rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, rr.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			mockCrtl := gomock.NewController(t)
			defer mockCrtl.Finish()

			gtwMock := NewMockGateway(mockCrtl)
			tc.mockActions(*gtwMock)

			rr := httptest.NewRecorder()

			url := "/users/:id"

			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			testHandler := handler{gtw: gtwMock}
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
	id := RandomInt(1, 100)

	testCases := []struct {
		name          string
		mockActions   func(gtwMock MockGateway)
		checkResponse func(rr *httptest.ResponseRecorder)
	}{
		{
			name: "User Updated",
			mockActions: func(gtwMock MockGateway) {
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
			mockActions: func(gtwMock MockGateway) {
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

			gtwMock := NewMockGateway(mockCtrl)
			tc.mockActions(*gtwMock)

			data, err := json.Marshal(userReq)
			require.NoError(t, err)

			rr := httptest.NewRecorder()

			url := "/users/:id"

			req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			testHandler := handler{gtw: gtwMock}
			testHandler.UpdateUser(rr, req, []httprouter.Param{{
				Key:   "id",
				Value: strconv.Itoa(int(id)),
			}})

			tc.checkResponse(rr)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	id := RandomInt(0, 100)

	testCases := []struct {
		name          string
		mockActions   func(gtwMock MockGateway)
		checkResponse func(rr *httptest.ResponseRecorder)
	}{
		{
			name: "Delete User",
			mockActions: func(gtwMock MockGateway) {
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
			mockActions: func(gtwMock MockGateway) {
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

			gtwMock := NewMockGateway(mockCtrl)
			tc.mockActions(*gtwMock)

			rr := httptest.NewRecorder()

			url := "/users/:id"

			req, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			testHandler := handler{gtw: gtwMock}
			testHandler.DeleteUser(rr, req, []httprouter.Param{{
				Key:   "id",
				Value: strconv.Itoa(int(id)),
			}})

			tc.checkResponse(rr)
		})
	}
}

func createUserRequest() UserRequest {
	return UserRequest{
		Email: RandomEmail(),
		Pass:  RandomPass(),
		Nick:  RandomNick(),
		Name:  RandomName(),
		Age:   RandomAge(),
	}
}

func newUser(u UserRequest) User {
	return User{
		ID:    RandomInt(0, 100),
		Email: u.Email,
		Pass:  u.Pass,
		Nick:  u.Nick,
		Name:  u.Name,
		Age:   u.Age,
	}
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Gererate random Integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Gererate random String of length n
func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail() string {
	return RandomString(6) + "@testemail.com"
}

func RandomName() string {
	return RandomString(4) + "name"
}

func RandomNick() string {
	return RandomString(5) + "nick"
}

func RandomAge() int64 {
	return RandomInt(18, 90)
}

func RandomPass() string {
	return RandomString(8) + "pass"
}
