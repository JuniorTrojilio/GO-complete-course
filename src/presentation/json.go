package presentation

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON converter received data in json response to request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}
