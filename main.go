package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	db "github.com/crisnlopez/social-media-bkend/internal/database"
	"github.com/crisnlopez/social-media-bkend/internal/handler"
)

func main() { 
  db, err := db.OpenDB("social_media")
  if err != nil {
    log.Fatal(err)
  } 

  userHandler := handler.New(db)
  // Router
  router := httprouter.New()
  router.GET("/users/:id", userHandler.GetUser)
  router.PUT("/users/:id", userHandler.UpdateUser)
  router.POST("/users", userHandler.CreateUser)
  router.DELETE("/users/:id", userHandler.DeleteUser)
 
	const addr = "localhost:8080"
	srv := http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Blocks forever, until the server
	// has an unrecoverable error
	fmt.Println("server started on", addr)
  err = srv.ListenAndServe()
	log.Fatal(err)
}
