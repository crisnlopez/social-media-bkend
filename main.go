package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	db "github.com/crisnlopez/social-media-bkend/internal/database"
)

type errorBody struct {
	Error string `json:"error"`
}

func main() {
  dbClient, err := db.NewClient("social_media")
  if err != nil {
    log.Fatal(err)
  }

  user, err := dbClient.CreateUser("text@example.com", "12345", "Pepe", 20)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("User created: %v\n", user)

  user, err = dbClient.UpdateUser("text@example.com", "123455555", "Carlos", 30)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("User update: %v\n", user)

  user, err = dbClient.UpdateUser("text@example_false.com", "12345", "Robert", 12)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Printf("User update, but it should't: %+v\n", user)

  user, err = dbClient.GetUser("text@example.com")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Get user: %v\n", user)

  err = dbClient.DeleteUser("text@example.com")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("User deleted")

  user, err = dbClient.GetUser("text@example.com")
  if err != nil {
    fmt.Println(err)
  }

  defer dbClient.Close()
	// Router
	mux := http.NewServeMux()

	const addr = "localhost:8080"
	srv := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Blocks forever, until the server
	// has an unrecoverable error
	fmt.Println("server started on", addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func responseWithError(w http.ResponseWriter, code int, err error) {
	log.Println(err)
	respondWithJSON(w, code, errorBody{
		Error: err.Error(),
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if payload != nil {
		response, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling", err)
			response, _ := json.Marshal(errorBody{
				Error: "error marshalling",
			})
			w.Write(response)
			w.WriteHeader(500)
			return
		}
		w.Write(response)
		w.WriteHeader(code)
	}
}
