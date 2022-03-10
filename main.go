package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	db "github.com/crisnlopez/social-media-bkend/internal/database"
)

type apiClient struct {
	dbClient db.Client
}

type errorBody struct {
	Error string `json:"error"`
}

func main() {
	// Create database
	dbClient := db.NewClient("db.json")
	err := dbClient.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}
	apiClient := apiClient{
		dbClient: dbClient,
	}

	// Router
	mux := http.NewServeMux()
	mux.HandleFunc("/users", apiClient.endpointUsersHandler)
	mux.HandleFunc("/users/", apiClient.endpointUsersHandler)
  mux.HandleFunc("/posts", apiClient.endpointPostsHandler)
  mux.HandleFunc("/posts/", apiClient.endpointPostsHandler)

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
