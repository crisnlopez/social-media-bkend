package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	db "github.com/crisnlopez/social-media-bkend/internal/database"
)

func main() {
	// Create database
	dbClient := db.NewClient("db.json")
	err := dbClient.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}
	apiCfg := apiConfig{
		dbClient: dbClient,
	}

	// Router
	mux := http.NewServeMux()
	mux.HandleFunc("/", testHandler)
	mux.HandleFunc("/err", testErrHanlder)
	mux.HandleFunc("/users", apiCfg.endpointUsersHandler)
	mux.HandleFunc("/users/", apiCfg.endpointUsersHandler)

	srv := http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

type apiConfig struct {
	dbClient db.Client
}

func (apiCfg apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// call  GET handler
	case http.MethodPost:
		//call POST handler
		apiCfg.handlerCreateUser(w, r)
	case http.MethodPut:
		//call PUT handler
	case http.MethodDelete:
		//call DELETE handler
	default:
		responseWithError(w, 404, errors.New("method not supported"))
	}
}

// Post handler reads the request body and creates a user with the given parameters
func (apiCfg apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params) // Store decode JSON in params struct
	if err != nil {
		responseWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := apiCfg.dbClient.CreateUser(params.Email, params.Password, params.Name, params.Age)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
	}
	respondWithJSON(w, http.StatusCreated, user)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, db.User{
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
