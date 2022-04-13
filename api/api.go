package api

import (
	"log"
	"os"

	"github.com/crisnlopez/social-media-bkend/internal/database"
	handler "github.com/crisnlopez/social-media-bkend/internal/user/handler"
)

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
