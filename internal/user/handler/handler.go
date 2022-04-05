package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/crisnlopez/social-media-bkend/internal/response"
  "github.com/crisnlopez/social-media-bkend/internal/user/models"
  gtw "github.com/crisnlopez/social-media-bkend/internal/user/gateway"
)

type UserHandler struct {
  gtw  gtw.UserGateway
}

func New(db *sql.DB) *UserHandler{
  return &UserHandler{
    gtw: gtw.NewGateway(db),
  }
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  // Decode request
  decoder := json.NewDecoder(r.Body)
  newUser := user.UserRequest{}
  err := decoder.Decode(&newUser)
  if err != nil {
    http.Error(w, err.Error(), 400)
    return
  }

  // Check if user already exists
  exists, err := h.gtw.GetUserEmail(newUser.Email)
  if err != nil{
    response.RespondWithError(w, http.StatusInternalServerError, err)
    return 
  }
  if exists{
    response.RespondWithError(w, http.StatusBadRequest, errors.New("User provided already exists"))
    return 
  }

  // Create User
  user, err := h.gtw.CreateUser(&newUser)
  if err != nil {
    response.RespondWithError(w, http.StatusInternalServerError, err)
    return
  }
     
  log.Println(user)
  response.RespondWithJSON(w, 200, &user)
  }

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  user := user.User{}

  // Getting userID from Request
  userID, err := strconv.Atoi(ps.ByName("id"))
  if err != nil {
    response.RespondWithError(w, http.StatusInternalServerError, err)
    return
  }
  if userID == 0 {
    response.RespondWithError(w, http.StatusBadRequest, errors.New("no userID provided to get user!"))
    return
  }

  response.RespondWithJSON(w, http.StatusOK, user)
}

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  // Getting id from Request
  id, err := strconv.Atoi(ps.ByName("id"))
  if id == 0 {
    response.RespondWithError(w, http.StatusBadRequest, errors.New("no userID provided to update user!"))
    return
  }

  // Decode JSON from request
  decoder := json.NewDecoder(r.Body)
  user := user.UserRequest{}
  err = decoder.Decode(&user)
  if err != nil {
    response.RespondWithError(w, http.StatusInternalServerError, err)
    return
  }

  // Updating user
  userUpdate, err := h.gtw.UpdateUser(&user, int64(id))
  if err != nil {
    response.RespondWithError(w, http.StatusInternalServerError, err)
    return
  }

  response.RespondWithJSON(w, http.StatusOK, userUpdate)
}

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  // Getting userID from Request
  id, err := strconv.Atoi(ps.ByName("id"))
  if err != nil {
    response.RespondWithError(w, http.StatusInternalServerError, err)
  }
  if id == 0 {
    response.RespondWithError(w, http.StatusBadRequest, errors.New("no userID provided to delete user!"))
    return
  }

  // Deleting User
  err = h.gtw.DeleteUser(id)
  if err != nil {
    response.RespondWithError(w, http.StatusInternalServerError, err)
    return
  }

  response.RespondWithJSON(w, http.StatusOK, "User Deleted")
}
