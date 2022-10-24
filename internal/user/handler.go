package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type Gateway interface {
	CreateUser(u UserRequest) (int64, error)
	GetUser(id int64) (User, error)
	UpdateUser(u UserRequest, id int64) (int64, error)
	DeleteUser(id int) error
}

type handler struct {
	gtw Gateway
}

func NewHandler(gtw Gateway) *handler {
	return &handler{
		gtw: gtw,
	}
}

func (h handler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Decode request
	decoder := json.NewDecoder(r.Body)
	var newUser UserRequest
	err := decoder.Decode(&newUser)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	err = validateRequest(newUser)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Create User
	id, err := h.gtw.CreateUser(newUser)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, 201, &id)
}

func (h handler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Getting id from Request
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	if id == 0 {
		respondWithError(w, http.StatusBadRequest, errors.New("no userID provided to get user!"))
		return
	}

	// Call GetUser
	user, err := h.gtw.GetUser(int64(id))
	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, err)
		return
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (h handler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Getting id from Request
	id, err := strconv.Atoi(ps.ByName("id"))
	if id == 0 {
		respondWithError(w, http.StatusBadRequest, errors.New("no userID provided to update user!"))
		return
	}

	// Decode JSON from request
	decoder := json.NewDecoder(r.Body)
	var updateUser UserRequest
	err = decoder.Decode(&updateUser)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	// Updating user
	rows, err := h.gtw.UpdateUser(updateUser, int64(id))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusOK, rows)
}

func (h handler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Getting userID from Request
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
	}
	if id == 0 {
		respondWithError(w, http.StatusBadRequest, errors.New("no userID provided to delete user!"))
		return
	}

	// Deleting User
	err = h.gtw.DeleteUser(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusOK, "User Deleted")
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

type errorBody struct {
	Error string `json:"error"`
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content Type", "application/json")
	w.Header().Set("Access Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if payload != nil {
		response, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling", err)

			w.WriteHeader(500)
			response, _ := json.Marshal(errorBody{
				Error: "Error Marshaling",
			})

			w.Write(response)
			return
		}

		w.WriteHeader(code)
		w.Write(response)
	}
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	if err == nil {
		log.Println("Don't call RespondWithError with nil err!")
		return
	}

	log.Println(err)
	respondWithJSON(w, code, errorBody{
		Error: err.Error(),
	})
}
