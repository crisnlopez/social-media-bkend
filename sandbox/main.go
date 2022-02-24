package main

import (
	"fmt"
	"log"

	db "github.com/crisnlopez/social-media-bkend/internal/database"
)

func main() {
	client := db.NewClient("db.json")
	err := client.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database ensured!")
}
