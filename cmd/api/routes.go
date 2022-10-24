package api

import (
	"github.com/crisnlopez/social-media-bkend/internal/user"
	"github.com/julienschmidt/httprouter"
)

func mapRoutes() *httprouter.Router {
	r := httprouter.New()

	gtw := user.NewGateway()
	handler := user.NewHandler(gtw)

	r.GET("/users/:id", handler.GetUser)
	r.PUT("/users/:id", handler.UpdateUser)
	r.POST("/users", handler.CreateUser)
	r.DELETE("/users/:id", handler.DeleteUser)

	return r
}
