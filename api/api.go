package api

import (
	"log"

	"github.com/crisnlopez/social-media-bkend/internal/database"
	handler "github.com/crisnlopez/social-media-bkend/internal/user/handler"
)

func Start(port string) {
  db, err := database.OpenDB("social_media")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  r := routes(handler.New(db))
  server := newServer(port, r)

  server.Start(port)
}
