package main

import (
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/users/")

	err := apiCfg.dbClient.DeleteUser(email)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, struct{}{})
}
