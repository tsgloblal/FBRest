package handlers

import (
	"net/http"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/fizzbuzz/internal/models"
	"github.com/fizzbuzz/internal/services"
	"github.com/fizzbuzz/utils"
	"github.com/gorilla/schema"
	"github.com/rs/zerolog"
)

type FizzBuzzHandler struct {
	service services.Service
	Log     zerolog.Logger
}

func NewFizzBuzzHandler(service services.Service) *FizzBuzzHandler {
	return &FizzBuzzHandler{
		service: service,
		Log:     zerolog.New(os.Stderr).With().Timestamp().Logger(),
	}
}

// Get Fizz Buzz godoc
//
//	@Summary	Handle Get Fizz Buzz
//	@Tags		Fizz-Buzz
//	@Produce	json
//	@Success	200	{object}	string
//	@Router		/api/fizzbuzz [get]
//	@Param		int1	query	int		false	"int1"
//	@Param		int2	query	int		false	"int2"
//	@Param		limit	query	int		false	"limit"
//	@Param		string1	query	string	false	"string1"
//	@Param		string2	query	string	false	"string2"
func (h *FizzBuzzHandler) GetFizzBuzz(w http.ResponseWriter, r *http.Request) {

	qp := models.DefaultFizzBuzz
	if err := schema.NewDecoder().Decode(&qp, r.URL.Query()); err != nil {
		h.Log.Warn().Any("query", r.URL.Query()).Err(err).Msg("error parsing query params")
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Failed to parse query params", err.Error())
		return
	}

	if v, err := govalidator.ValidateStruct(qp); !v {
		h.Log.Warn().Any("query", r.URL.Query()).Err(err).Msg("error parsing query params")
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Failed to parse query params", err.Error())
		return
	}

	fizzBuzz, err := h.service.GetFizzBuzz(r.Context(), qp)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to get fizz-buzz", err.Error())
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "fizz-buzz retrieved successfully", fizzBuzz)
}

// Get Fizz Buzz godoc
//
//	@Summary	Handle Get Fizz Buzz
//	@Tags		Fizz-Buzz
//	@Produce	json
//	@Success	200	{object}	models.FizzBuzzRequest
//	@Router		/api/fizzbuzz/stats [get]
func (h *FizzBuzzHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	fizzBuzz, err := h.service.GetStats(r.Context())
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to get fizz-buzz stats", err.Error())
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "fizz-buzz stats retrieved successfully", fizzBuzz)
}
