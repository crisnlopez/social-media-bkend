package api

import (
	"github.com/crisnlopez/social-media-bkend/internal/user"
	"github.com/julienschmidt/httprouter"
)

func mapRoutes(handler *user.UserHandler) *httprouter.Router {
	r := httprouter.New()

	r.GET("/users/:id", handler.GetUser)
	r.PUT("/users/:id", handler.UpdateUser)
	r.POST("/users", handler.CreateUser)
	r.DELETE("/users/:id", handler.DeleteUser)

	return r
}
