package main

import (
	"encoding/json"
	"net/http"
)

// Post handler reads the request body and creates a user with the given parameters
func (apiCfg apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params) // Store decode JSON in params struct
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := apiCfg.dbClient.CreateUser(params.Email, params.Password, params.Name, params.Age)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
	}
	respondWithJSON(w, http.StatusCreated, user)
}