package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fizzbuzz/internal/models"
	"github.com/fizzbuzz/internal/services/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetFizzBuzzDefault(t *testing.T) {

	service := mock.ServiceMock{}

	service.GetFizzBuzzFunc = func(ctx context.Context, fizzBuzz models.FizzBuzz) (string, error) {
		assert.Equal(t, models.DefaultFizzBuzz, fizzBuzz)
		return "fizzbuzz", nil
	}

	fizzBuzzHandler := SetupRouter(&service)

	req, err := http.NewRequest(http.MethodGet, "/api/fizzbuzz", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	fizzBuzzHandler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetFizzBuzzNokBigInt1(t *testing.T) {

	service := mock.ServiceMock{}

	fizzBuzzHandler := SetupRouter(&service)

	req, err := http.NewRequest(http.MethodGet, "/api/fizzbuzz?int1=101", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	fizzBuzzHandler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetFizzBuzz(t *testing.T) {

	service := mock.ServiceMock{}

	service.GetFizzBuzzFunc = func(ctx context.Context, fizzBuzz models.FizzBuzz) (string, error) {
		expected := models.DefaultFizzBuzz
		expected.Int1 = 1
		expected.Limit = 4
		assert.Equal(t, expected, fizzBuzz)
		return "fizzbuzz", nil
	}

	fizzBuzzHandler := SetupRouter(&service)

	req, err := http.NewRequest(http.MethodGet, "/api/fizzbuzz?int1=1&limit=4", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	fizzBuzzHandler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetStats(t *testing.T) {

	service := mock.ServiceMock{}

	service.GetStatsFunc = func(ctx context.Context) (models.FizzBuzzRequest, error) {
		return models.FizzBuzzRequest{}, nil
	}

	fizzBuzzHandler := SetupRouter(&service)

	req, err := http.NewRequest(http.MethodGet, "/api/fizzbuzz/stats", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	fizzBuzzHandler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}
