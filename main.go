package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
  
	db "github.com/crisnlopez/social-media-bkend/internal/database"
  hd "github.com/crisnlopez/social-media-bkend/internal/handler"
)

func main() { 
  db, err := db.OpenDB("social_media")
  if err != nil {
  log.Fatal(err)
  } 

  // Router
  router := httprouter.New()
  router.GET("/users/:id", hd.UserHandler{Db:db}.GetUser)
  router.PUT("/users/:id", hd.UserHandler{Db: db}.UpdateUser)
  router.POST("/users", hd.UserHandler{Db: db}.CreateUser)
  router.DELETE("/users/:id", hd.UserHandler{Db: db}.DeleteUser)
 
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
