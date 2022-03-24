package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	res "github.com/crisnlopez/social-media-bkend/internal/response"
)

type userHandler struct {
  // we have a user gateway between handler and repository
  repo Repository
}

func New(db *sql.DB) userHandler{
  return userHandler{
    repo: newRepository(db),
  }
}

func (h userHandler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  // Decode request
  decoder := json.NewDecoder(r.Body)
  user := User{}
  err := decoder.Decode(&user)
  if err != nil {
    http.Error(w, err.Error(), 400)
    return
  }
  // Check if user already exists
  exists, err := h.repo.GetUserEmail(user)

  if err != nil{
    return 
  }
  if exists{
    return 
  }


  // Execute Query
  result, err := h.db.Exec("INSERT INTO users (email, pass, user_nick, user_name, age) VALUES (?, ?, ?, ?, ?)", user.Email, user.Pass, user.Nick, user.Name, user.Age)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  // Get UserID
  id, err := result.LastInsertId()
  if err != nil {
    res.RespondWithError(w, http.StatusInternalServerError, err)
    return
  }
  user.ID = int(id)

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

func (h userHandler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  user := User{}

  // Getting userID from Request
  userID, err := strconv.Atoi(ps.ByName("id"))
  if err != nil {
    res.RespondWithError(w, http.StatusInternalServerError, err)
    return
  }
  if userID == 0 {
    res.RespondWithError(w, http.StatusBadRequest, errors.New("no userID provided to get user!"))
    return
  }

  // Getting User
  if err := h.db.QueryRow("SELECT * FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Email, &user.Pass, &user.Nick, &user.Name, &user.Age); err != nil {
    if err == sql.ErrNoRows { // If user doesn't exist
      res.RespondWithError(w, http.StatusNotFound, err)
      return
    } else {
      res.RespondWithError(w, http.StatusInternalServerError, err)
      return
    }
  }

  res.RespondWithJSON(w, http.StatusOK, user)
}

func (h userHandler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  // Getting userID from Request
  userID, err := strconv.Atoi(ps.ByName("id"))
  if userID == 0 {
    res.RespondWithError(w, http.StatusBadRequest, errors.New("no userID provided to update user!"))
    return
  }

  // Decode JSON from request
  decoder := json.NewDecoder(r.Body)
  user := UserUpdated{}
  err = decoder.Decode(&user)
  if err != nil {
    res.RespondWithError(w, http.StatusInternalServerError, err)
    return
  }

  // Check if user exists
  var idCheck string
  row := h.db.QueryRow("SELECT id FROM users WHERE id = ?", userID)
  err = row.Scan(&idCheck)
  if err != nil {
    res.RespondWithError(w, http.StatusInternalServerError, err)
    return
  }

  // Updating user
  _, err = h.db.Exec("UPDATE users SET pass = ?, user_name = ?, age = ?, user_nick = ?, email = ? WHERE id = ?", user.Pass, user.Name, user.Age, user.Nick, user.Email, userID)
  if err != nil {
    res.RespondWithError(w, http.StatusInternalServerError, err)
    return
  }

  res.RespondWithJSON(w, http.StatusOK, user)
}

func (h userHandler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  // Getting userID from Request
  userID, err := strconv.Atoi(ps.ByName("id"))
  if userID == 0 {
    res.RespondWithError(w, http.StatusBadRequest, errors.New("no userID provided to delete user!"))
    return
  }

  // Deleting User
  _, err = h.db.Exec("DELETE FROM users WHERE id = ?", userID)
  if err != nil {
    res.RespondWithError(w, http.StatusInternalServerError, err)
    return
  }

  res.RespondWithJSON(w, http.StatusOK, "User Deleted")
}
