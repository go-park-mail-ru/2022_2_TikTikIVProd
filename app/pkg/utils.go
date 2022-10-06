package pkg

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Body interface{} `json:"body"`
}

type Error struct {
	Message string `json:"error"`
}

func JSONresponse(w http.ResponseWriter, code int, body interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(Response {
									Body:body,
								  })
	if err != nil {
		return err
	}

	w.WriteHeader(code)
	w.Write(response)
	return nil
}

func ErrorResponse(w http.ResponseWriter, code int, errorMessage string) error {
	return JSONresponse(w, code, Error {
										Message: errorMessage,
									})
}
