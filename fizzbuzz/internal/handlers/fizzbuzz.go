package handlers

import (
	"net/http"

	"github.com/fizzbuzz/internal/models"
	"github.com/fizzbuzz/internal/services"
	"github.com/fizzbuzz/utils"
)

type FizzBuzzHandler struct {
	service services.Service
}

func NewFizzBuzzHandler(service services.Service) *FizzBuzzHandler {
	return &FizzBuzzHandler{service: service}
}

// go-validator
// docs
func (h *FizzBuzzHandler) GetFizzBuzz(w http.ResponseWriter, r *http.Request) {
	fizzBuzz, err := h.service.GetFizzBuzz(r.Context(), models.DefaultFizzBuzz)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to get fizz-buzz", err.Error())
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "fizz-buzz retrieved successfully", fizzBuzz)
}
