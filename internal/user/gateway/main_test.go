package gateway

import (
	"log"
	"os"
	"testing"

	db "github.com/crisnlopez/social-media-bkend/internal/database"
  "github.com/crisnlopez/social-media-bkend/internal/user/repository"
)

var testQueries *repository.UserQueries 

func TestMain(m *testing.M) {
  conn, err := db.OpenDB("social_media")
  if err != nil {
    log.Fatal(err)
  }

  testQueries = repository.NewRepository(conn)

  os.Exit(m.Run())
}
