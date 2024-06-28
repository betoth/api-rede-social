package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON return a response in JSON
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {

		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}

}

// ErrorJSON return an error in JSON
func ErrorJSON(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})

}
