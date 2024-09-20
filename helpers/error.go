package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Error   bool   `json:"error"`
}

func ErrorCreate(message string) *Error {
	return &Error{Message: message, Error: true}
}

func ErrorResponse(message string, code int, res http.ResponseWriter) {
	logger := LoggerCreate("ErrorResponse")

	errorResponse, err := json.Marshal(ErrorCreate(message))
	if err != nil {
		logger.Error(fmt.Sprintf("Error marshalling error response: %s", err.Error()))
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Error(res, string(errorResponse), code)
}
