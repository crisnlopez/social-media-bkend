package main

import (
	"errors"
	"net/http"
)

// Check the request method and call the corresponding handler function
func (apiClnt apiClient) endpointUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// call  GET handler
		apiClnt.handlerGetUser(w, r)
	case http.MethodPost:
		//call POST handler
		apiClnt.handlerCreateUser(w, r)
	case http.MethodPut:
		//call PUT handler
		apiClnt.handlerUpdateUser(w, r)
	case http.MethodDelete:
		//call DELETE handler
		apiClnt.handlerDeleteUser(w, r)
	default:
		responseWithError(w, 404, errors.New("method not supported"))
	}
}
