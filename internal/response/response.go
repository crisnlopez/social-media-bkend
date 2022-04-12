package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorBody struct {
	Error string `json:"error"`
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content Type", "application/json")
	w.Header().Set("Access Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if payload != nil {
		response, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling", err)

			w.WriteHeader(500)
			response, _ := json.Marshal(errorBody{
				Error: "Error Marshaling",
			})

			w.Write(response)
			return
		}

		w.WriteHeader(code)
		w.Write(response)
	}
}

func RespondWithError(w http.ResponseWriter, code int, err error) {
	if err == nil {
		log.Println("Don't call RespondWithError with nil err!")
		return
	}

	log.Println(err)
	RespondWithJSON(w, code, errorBody{
		Error: err.Error(),
	})
}
