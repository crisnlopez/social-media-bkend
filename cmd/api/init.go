package api

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/crisnlopez/social-media-bkend/internal/database"
	"github.com/crisnlopez/social-media-bkend/internal/user/handler"
	"github.com/julienschmidt/httprouter"
)

type server struct {
	*http.Server
}

func newServer(listening string, router *httprouter.Router) *server {
	srv := &http.Server{
		Addr:         ":" + listening,
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return &server{srv}
}

func (srv *server) Start(port string) {
	log.Printf("Starting Server on port: %v\n", port)

	err := srv.ListenAndServe()
	log.Fatal(err)
}

func Start(port string) {
	db, err := database.OpenDB(os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := routes(handler.New(db))
	server := newServer(port, r)

	server.Start(port)
}
