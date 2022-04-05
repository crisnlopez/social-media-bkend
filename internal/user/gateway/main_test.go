package user

import (
	"log"
	"os"
	"testing"

	db "github.com/crisnlopez/social-media-bkend/internal/database"
)

var testQueries *UserQueries 

func TestMain(m *testing.M) {
  conn, err := db.OpenDB("social_media")
  if err != nil {
    log.Fatal(err)
  }

  testQueries = NewRepository(conn)

  os.Exit(m.Run())
}
