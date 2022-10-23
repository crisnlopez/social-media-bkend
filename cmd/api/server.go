package api

import (
	"log"
	"net/http"
	"time"

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

func (srv *server) Start(port string) error {
	log.Printf("Starting Server on port: %v\n", port)

	return srv.ListenAndServe()
}
