package main

import (
	"encoding/json"
	"net/http"
)

type newUserDecoder struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}

// Post handler reads the request body and creates a user with the given parameters
func (apiClnt apiClient) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := newUserDecoder{}
	err := decoder.Decode(&params) // Store decode JSON in params struct
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err)
		return
	}

  err = userIsElegible(params.Email,params.Password,params.Age)
  if err != nil {
    responseWithError(w, http.StatusBadRequest, err)
    return
  }
  
	user, err := apiClnt.dbClient.CreateUser(params.Email, params.Password, params.Name, params.Age)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err)
	}

	respondWithJSON(w, http.StatusCreated, user)
}
