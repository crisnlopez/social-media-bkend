package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"

	"github.com/crisnlopez/social-media-bkend/internal/response"
)

type UserHandler struct {
	Gtw Gateway
}

func NewHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		Gtw: NewGateway(db),
	}
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Decode request
	decoder := json.NewDecoder(r.Body)
  var newUser UserRequest
	err := decoder.Decode(&newUser)
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	err = validateRequest(newUser)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Create User
	id, err := h.Gtw.CreateUser(newUser)
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
		return
	}
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err)
		return
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
  var updateUser UserRequest
	err = decoder.Decode(&updateUser)
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	// Updating user
	rows, err := h.Gtw.UpdateUser(updateUser, int64(id))
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

func validateRequest(u UserRequest) error {
	err := validator.New().Struct(u)

	if err != nil {
		if invalid, ok := err.(*validator.InvalidValidationError); ok {
			return invalid
		}

		return err.(validator.ValidationErrors)
	}

	return nil
}
