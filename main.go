package main

import (
	"log"
	"os"

	"github.com/crisnlopez/social-media-bkend/api"
)

const defaultPort = "8080"

func main() {
	log.Println("starting API")
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	api.Start(port)
}
