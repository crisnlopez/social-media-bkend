package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/crisnlopez/social-media-bkend/internal/database"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", testHandler)
	mux.HandleFunc("/err", testErrHanlder)

	srv := http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := srv.ListenAndServe()
	log.Fatal(err)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, database.User{
		Email:    "test@example.com",
		Password: "Password",
		Name:     "Pepe",
		Age:      20,
	})
}

func testErrHanlder(w http.ResponseWriter, r *http.Request) {
	err := errors.New("server error")
	responseWithError(w, 500, err)
}

type errorBody struct {
	Error string `json:"error"`
}

func responseWithError(w http.ResponseWriter, code int, err error) {
	log.Println(err)
	respondWithJSON(w, 500, errorBody{
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
