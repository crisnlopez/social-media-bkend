package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}

	email := strings.TrimPrefix(r.URL.Path, "/users/")
	user, err := apiCfg.dbClient.UpdateUser(email, params.Password, params.Name, params.Age)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}
