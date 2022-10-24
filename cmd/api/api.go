package api

import (
	"log"
	"os"

	"github.com/crisnlopez/social-media-bkend/internal/database"
	"github.com/crisnlopez/social-media-bkend/internal/user"
)

func Start(port string) {
	db, err := database.OpenDB(os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mapRoutes(user.NewHandler(db))
	server := newServer(port, r)

	server.Start(port)
}
