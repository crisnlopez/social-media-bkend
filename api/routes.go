package api

import (
	"database/sql"

	handler "github.com/crisnlopez/social-media-bkend/internal/user/handler"
	"github.com/julienschmidt/httprouter"
)

// Does the routes
func defineRoutes(r *httprouter.Router, db *sql.DB) {
	u := handler.New(db)

	r.GET("/users/:id", u.GetUser)
	r.PUT("/users/:id", u.UpdateUser)
	r.POST("/users", u.CreateUser)
	r.DELETE("/users/:id", u.DeleteUser)
}
