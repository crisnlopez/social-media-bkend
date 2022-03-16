package main

import (
	"encoding/json"
	"net/http"
	"strings"
)
  
type updateUserDecoder struct {
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}
  
func (apiClnt apiClient) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/users/")
	
	decoder := json.NewDecoder(r.Body)
	params := updateUserDecoder{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}

  err = userIsEligible(email, params.Password, params.Age)
  if err != nil {
    responseWithError(w, http.StatusBadRequest, err)
    return
  }

	user, err := apiClnt.dbClient.UpdateUser(email, params.Password, params.Name, params.Age)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}
