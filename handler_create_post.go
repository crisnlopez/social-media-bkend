package main

import (
	"encoding/json"
	"net/http"
)

type createPostDecoder struct {
  UserEmail string `json:"userEmail"`
  Text string `json:"text"`
}

func (apiClnt apiClient) handlerCreatePost(w http.ResponseWriter, r *http.Request) {
  decoder := json.NewDecoder(r.Body)
  params :=  createPostDecoder{}

  err := decoder.Decode(&params)
  if err != nil {
    responseWithError(w, http.StatusBadRequest, err)
    return
  }

  post, err := apiClnt.dbClient.CreatePost(params.UserEmail, params.Text)
  if err != nil {
    responseWithError(w, http.StatusBadRequest, err)
    return
  }

  respondWithJSON(w, http.StatusOK, post)
}
