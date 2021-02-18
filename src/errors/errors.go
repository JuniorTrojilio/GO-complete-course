package errors

import (
	"api/src/presentation"
	"net/http"
)

// Err returns an error in JSON format
func Err(w http.ResponseWriter, statusCode int, err error) {
	presentation.JSON(w, statusCode, struct {
		Err string `json:"erro"`
	}{
		Err: err.Error(),
	})
}
