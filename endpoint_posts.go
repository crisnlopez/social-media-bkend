package main

import (
	"errors"
	"net/http"
)

func (apiClnt apiClient) endpointPostsHandler(w http.ResponseWriter, r *http.Request) {
 switch r.Method {
	case http.MethodGet:
		// call  GET handler
		apiClnt.handlerRetrievePost(w, r)
	case http.MethodPost:
		//call POST handler
		apiClnt.handlerCreatePost(w, r)	
	case http.MethodDelete:
		//call DELETE handler
		apiClnt.handlerDeletePost(w, r)
	default:
		responseWithError(w, 404, errors.New("method not supported")) 
}
}
