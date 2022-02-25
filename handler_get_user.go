package main

import (
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/users/")

	user, err := apiCfg.dbClient.GetUser(email)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}
