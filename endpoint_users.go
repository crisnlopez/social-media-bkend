package main

import (
	"errors"
	"net/http"
)

// Check the request method and call the corresponding handler function
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
