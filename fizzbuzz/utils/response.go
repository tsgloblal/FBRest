package utils

import (
	"encoding/json"
	"net/http"
)

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]interface{}{
		"message": message,
		"data":    data,
	}

	json.NewEncoder(w).Encode(response)
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message, error string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]interface{}{
		"message": message,
		"error":   error,
	}

	json.NewEncoder(w).Encode(response)
}
