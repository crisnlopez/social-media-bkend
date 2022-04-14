package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/crisnlopez/social-media-bkend/internal/response"
	gtw "github.com/crisnlopez/social-media-bkend/internal/user/gateway"
	"github.com/crisnlopez/social-media-bkend/internal/user/models"
)

type UserHandler struct {
	Gtw gtw.UserGateway
}

func New(db *sql.DB) *UserHandler {
	return &UserHandler{
		Gtw: gtw.NewGateway(db),
	}
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Decode request
	decoder := json.NewDecoder(r.Body)
	newUser := user.UserRequest{}
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if user already exists
	exists, err := h.Gtw.GetUserEmail(newUser.Email)
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}
	if exists {
		response.RespondWithError(w, http.StatusBadRequest, errors.New("User provided already exists"))
		return
	}

	// Create User
	id, err := h.Gtw.CreateUser(&newUser)
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	response.RespondWithJSON(w, 201, &id)
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Getting id from Request
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}
	if id == 0 {
		response.RespondWithError(w, http.StatusBadRequest, errors.New("no userID provided to get user!"))
		return
	}

	// Call GetUser
	user, err := h.Gtw.GetUser(int64(id))
	if err == sql.ErrNoRows {
 	response.RespondWithError(w, http.StatusNotFound, err)
	}
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err)
	}

	response.RespondWithJSON(w, http.StatusOK, user)
}

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Getting id from Request
	id, err := strconv.Atoi(ps.ByName("id"))
	if id == 0 {
		response.RespondWithError(w, http.StatusBadRequest, errors.New("no userID provided to update user!"))
		return
	}

	// Decode JSON from request
	decoder := json.NewDecoder(r.Body)
	updateUser := user.UserRequest{}
	err = decoder.Decode(&updateUser)
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	// Updating user
	rows, err := h.Gtw.UpdateUser(&updateUser, int64(id))
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	response.RespondWithJSON(w, http.StatusOK, rows)
}

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Getting userID from Request
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err)
	}
	if id == 0 {
		response.RespondWithError(w, http.StatusBadRequest, errors.New("no userID provided to delete user!"))
		return
	}

	// Deleting User
	err = h.Gtw.DeleteUser(id)
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	response.RespondWithJSON(w, http.StatusOK, "User Deleted")
}
