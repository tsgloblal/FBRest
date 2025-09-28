package handlers

import (
	"github.com/fizzbuzz/internal/services"
	"github.com/gorilla/mux"
)

func SetupRouter(service services.Service) *mux.Router {
	r := mux.NewRouter()

	r.Use(CORSMiddleware)
	r.Use(LoggingMiddleware)

	fizzBuzzHandler := NewFizzBuzzHandler(service)

	setupRoutes(r, fizzBuzzHandler)
	return r
}

func setupRoutes(r *mux.Router, fizzBuzzHandler *FizzBuzzHandler) {

	r.HandleFunc("/health", HealthHandler).Methods("GET")

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/fizzbuzz", fizzBuzzHandler.GetFizzBuzz).Methods("GET")
}
