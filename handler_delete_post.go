package main

import (
	"net/http"
	"strings"
)

func (apiClnt apiClient) handlerDeletePost(w http.ResponseWriter, r *http.Request) {
  postId := strings.TrimPrefix(r.URL.Path, "/posts/") 

  err := apiClnt.dbClient.DeletePost(postId)
  if err != nil {
    responseWithError(w, http.StatusBadRequest, err)
    return
  }

  respondWithJSON(w, http.StatusOK, struct{}{})
}
