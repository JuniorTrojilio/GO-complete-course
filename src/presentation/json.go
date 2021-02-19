package presentation

import (
	"api/src/models"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// JSON converter received data in json response to request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// RemoveBlankSpace removes blank spaces
func RemoveBlankSpace(user *models.User) {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
