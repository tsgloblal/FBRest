package handlers

import (
	_ "embed"

	_ "github.com/fizzbuzz/docs"
	"github.com/fizzbuzz/internal/services"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title			Fizz Buzz API
// @version		1.0
// @description	This is an API responsible for handling Fizz Buzz requests
// @Schemes		http https
// @host			localhost
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

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler())

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/fizzbuzz", fizzBuzzHandler.GetFizzBuzz).Methods("GET")
	api.HandleFunc("/fizzbuzz/stats", fizzBuzzHandler.GetStats).Methods("GET")
}
