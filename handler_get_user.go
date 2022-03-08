package main

import (
	"net/http"
	"strings"
)

func (apiClnt apiClient) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/users/")

	user, err := apiClnt.dbClient.GetUser(email)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}
