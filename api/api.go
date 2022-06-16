package api

import (
	"log"
	"os"

	"github.com/crisnlopez/social-media-bkend/internal/database"
	"github.com/julienschmidt/httprouter"
)

func Start(port string) {
	db, err := database.OpenDB(os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := httprouter.New()
	defineRoutes(r, db)

	server := newServer(port, r)
	server.Start(port)
}
