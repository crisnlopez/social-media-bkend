package main

import (
	"log"
	"net/http"
	"time"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Content-Type Header
	w.WriteHeader(200)                                 // Status Code
	w.Write([]byte("{}"))                              // Empty JSON
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", testHandler)

	srv := http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := srv.ListenAndServe()
	log.Fatal(err)
}
