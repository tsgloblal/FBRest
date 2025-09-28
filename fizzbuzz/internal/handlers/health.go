package handlers

import (
	"net/http"
	"time"

	"github.com/fizzbuzz/utils"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	healthData := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"version":   "1.0.0",
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "Service is healthy", healthData)
}
