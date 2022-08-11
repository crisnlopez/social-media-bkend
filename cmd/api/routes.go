package api

import (
	"net/http"

	"github.com/crisnlopez/social-media-bkend/internal/response"
	handler "github.com/crisnlopez/social-media-bkend/internal/user/handler"
	"github.com/julienschmidt/httprouter"
)

func routes(services *handler.UserHandler) *httprouter.Router {
	r := httprouter.New()

  r.GET("/ping", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    response.RespondWithJSON(w, http.StatusOK, "pong")
  })
	r.GET("/users/:id", services.GetUser)
	r.PUT("/users/:id", services.UpdateUser)
	r.POST("/users", services.CreateUser)
	r.DELETE("/users/:id", services.DeleteUser)

	return r
}
