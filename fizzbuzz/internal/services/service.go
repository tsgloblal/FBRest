package services

import (
	"context"

	"github.com/fizzbuzz/internal/models"
	"github.com/fizzbuzz/internal/repository"
)

//go:generate moq -pkg mock -out ./mock/service_moq.go . Service
type Service interface {
	GetFizzBuzz(ctx context.Context, fizzBuzz models.FizzBuzz) (string, error)
}

type service struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetFizzBuzz(ctx context.Context, fizzBuzz models.FizzBuzz) (string, error) {
	return "", nil
}
