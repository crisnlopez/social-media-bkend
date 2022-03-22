package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	db "github.com/crisnlopez/social-media-bkend/internal/database"
	res "github.com/crisnlopez/social-media-bkend/internal/response"
)

type UserHandler struct {
  Db *sql.DB
}

type UserUpdated struct {
  Pass string `json:"pass"`
  Name string `json:"name"`
  Age  int    `json:"age"`
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  // Decode request
  decoder := json.NewDecoder(r.Body)
  user := db.User{}
  err := decoder.Decode(&user)
  if err != nil {
    http.Error(w, err.Error(), 400)
    return
  }

  // Check if user already exists
  var email string
  row := h.Db.QueryRow("SELECT email FROM users WHERE email = ?", user.Email)
  if err = row.Scan(&email); err != nil {
    if err != sql.ErrNoRows { // If already exist respond with error
      http.Error(w, err.Error(), 500)
      return
    }
  }

  // Execute Query
  _, err = h.Db.Exec("INSERT INTO users (email, pass, name, age) VALUES (?, ?, ?, ?)", user.Email, user.Pass, user.Name, user.Age)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }

  // Response
  w.WriteHeader(201)
  w.Header().Set("Content-Type", "application/json")

  response, err := json.Marshal(user)
  // If response error
  if err != nil {
    log.Println("Error Marshalling", err)
    w.WriteHeader(500)
    w.Write(response)
    return
  }

  w.Write(response)
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  user := db.User{}

  // Getting email user Request
  userEmail := ps.ByName("userEmail")
  if userEmail == "" {
    res.RespondWithError(w, http.StatusBadRequest, errors.New("no userEmail provided to get user!"))
    return
  }

  // Getting User
  if err := h.Db.QueryRow("SELECT * FROM users WHERE email = ?", userEmail).Scan(&user.Email, &user.Pass, &user.Name, &user.Age); err != nil {
    if err == sql.ErrNoRows {
      res.RespondWithError(w, 404, err)
    }
  }

  res.RespondWithJSON(w, http.StatusOK, user)
}

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

  // Getting email user Request
  userEmail := ps.ByName("userEmail")
  if userEmail == "" {
    res.RespondWithError(w, http.StatusBadRequest, errors.New("no userEmail provided to update user!"))
    return
  }

  // Decode JSON from request
  decoder := json.NewDecoder(r.Body)
  user := UserUpdated{}
  err := decoder.Decode(&user)
  if err != nil {
    res.RespondWithError(w, http.StatusInternalServerError, err)
  }

  // Check if user exists
  var col string
  row := h.Db.QueryRow("SELECT email FROM users WHERE email = ?", userEmail)
  err = row.Scan(&col)
  if err != nil {
    if err == sql.ErrNoRows {
      res.RespondWithError(w, http.StatusBadRequest, errors.New("the email provided doesn't exist"))
      return
    } else {
      res.RespondWithError(w, http.StatusInternalServerError, err)
    }
  }

  // Updating user
  _, err = h.Db.Exec("UPDATE users SET pass = ?, name = ?, age = ? WHERE email = ?", user.Pass, user.Name, user.Age, userEmail)
  if err != nil {
    res.RespondWithError(w, http.StatusInternalServerError, err)
    return
  }

  res.RespondWithJSON(w, http.StatusOK, user)
}
