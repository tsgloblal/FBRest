package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fizzbuzz/internal/models"
	"github.com/fizzbuzz/internal/repository"
	"github.com/fizzbuzz/utils"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

//go:generate moq -pkg mock -out ./mock/service_moq.go . Service
type Service interface {
	GetFizzBuzz(ctx context.Context, fizzBuzz models.FizzBuzz) (string, error)
	GetStats(ctx context.Context) (models.FizzBuzzRequest, error)
}

type service struct {
	repository  *repository.Repository
	redisClient *redis.Client
	Log         zerolog.Logger
}

func NewService(repository *repository.Repository, redisClient *redis.Client) Service {
	return &service{
		repository:  repository,
		redisClient: redisClient,
		Log:         zerolog.New(os.Stderr).With().Timestamp().Logger(),
	}
}

func (s *service) GetFizzBuzz(ctx context.Context, fizzBuzz models.FizzBuzz) (string, error) {
	result, err := s.getStored(ctx, fizzBuzz)
	if err != nil {
		s.Log.Error().Err(err).Msg("error fetching fizz-buzz")
		return "", fmt.Errorf("error fetching fizz-buzz")
	}

	if result == "" {
		result = s.newFizzBuzz(fizzBuzz)
	}

	fizzBuzzRequest := models.FizzBuzzRequest{
		Int1:   fizzBuzz.Int1,
		Int2:   fizzBuzz.Int2,
		Limit:  fizzBuzz.Limit,
		Str1:   fizzBuzz.Str1,
		Str2:   fizzBuzz.Str2,
		Result: result,
	}

	if err := s.setFizzBuzz(ctx, fizzBuzz, fizzBuzzRequest); err != nil {
		s.Log.Error().Err(err).Msg("error saving fizz-buzz")
		return "", fmt.Errorf("error saving fizz-buzz")
	}

	return fizzBuzzRequest.Result, nil
}

func (s *service) getStored(ctx context.Context, fizzBuzz models.FizzBuzz) (string, error) {
	encodedFizzbuzz, err := utils.EncodeStruct(fizzBuzz)
	if err != nil {
		s.Log.Error().Err(err).Msg("error encoding fizz-buzz")
		return "", fmt.Errorf("error encoding fizz-buzz")
	}

	cached := s.getCacheFizzBuzz(ctx, encodedFizzbuzz)
	if cached != "" {
		return cached, nil
	}
	stored, err := s.repository.GetFizzBuzz(ctx, fizzBuzz)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		s.Log.Error().Err(err).Msg("error fetching fizz-buzz")
		return "", fmt.Errorf("error fetching fizz-buzz")
	}

	return stored, nil

}

func (s *service) newFizzBuzz(fizzBuzz models.FizzBuzz) string {

	var result []string

	for i := 1; i <= fizzBuzz.Limit; i++ {
		var output string

		if i%fizzBuzz.Int1 == 0 {
			output += fizzBuzz.Str1
		}
		if i%fizzBuzz.Int2 == 0 {
			output += fizzBuzz.Str2
		}
		if output == "" {
			output = strconv.Itoa(i)
		}
		result = append(result, output)
	}

	return strings.Join(result, ",")

}

func (s *service) GetStats(ctx context.Context) (models.FizzBuzzRequest, error) {
	result, err := s.repository.GetTop(ctx, 1)
	if err != nil {
		s.Log.Error().Err(err).Msg("error fetching fizz-buzz stats")
		return models.FizzBuzzRequest{}, fmt.Errorf("error fetching fizz-buzz stats")
	}

	return result[0], nil
}

func (s *service) setFizzBuzz(ctx context.Context, fizzBuzz models.FizzBuzz, fizzBuzzRequest models.FizzBuzzRequest) error {
	if err := s.repository.SetFizzBuzz(ctx, fizzBuzzRequest); err != nil {
		s.Log.Error().Err(err).Msg("error saving fizz-buzz")
		return fmt.Errorf("error saving fizz-buzz")
	}

	encodedFizzbuzz, err := utils.EncodeStruct(fizzBuzz)
	if err != nil {
		s.Log.Error().Err(err).Msg("error encoding fizz-buzz")
		return fmt.Errorf("error encoding fizz-buzz")
	}

	s.setCacheFizzBuzz(ctx, encodedFizzbuzz, fizzBuzzRequest.Result)

	return nil
}

func (s *service) getCacheFizzBuzz(ctx context.Context, fizzBuzz string) string {
	return s.redisClient.Get(ctx, fizzBuzz).Val()
}

func (s *service) setCacheFizzBuzz(ctx context.Context, fizzBuzz, result string) error {
	return s.redisClient.Set(ctx, fizzBuzz, result, time.Hour).Err()
}
