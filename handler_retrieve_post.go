package main

import (
	"errors"
	"net/http"
)

func (apiCltn apiClient) handlerRetrievePost(w http.ResponseWriter, r *http.Request) {
  usrEmail := r.URL.Query().Get("userEmail")
  if usrEmail == "" {
    responseWithError(w, http.StatusBadRequest, errors.New("no userEmail"))
    return
  }

  posts, err := apiCltn.dbClient.GetPosts(usrEmail)
  if err != nil {
    responseWithError(w, http.StatusBadRequest, err)
    return
  } 

  respondWithJSON(w, http.StatusOK, posts)
}
