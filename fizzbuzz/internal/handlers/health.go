package handlers

import (
	"net/http"
	"time"

	"github.com/fizzbuzz/internal/models"
	"github.com/fizzbuzz/utils"
)

// Get Health godoc
//
//	@Summary	Handle Get Health
//	@Tags		Health
//	@Produce	json
//	@Success	200	{object}	models.HealthData
//	@Router		/health [get]
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	healthData := models.HealthData{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
		Version:   "1.0.0",
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "Service is healthy", healthData)
}
