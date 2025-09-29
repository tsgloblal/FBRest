package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/fizzbuzz/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (s *Repository) SetFizzBuzz(ctx context.Context, req models.FizzBuzzRequest) error {
	query := `
        INSERT INTO fizzbuzz_requests (int1, int2, limit_value, str1, str2, result)
        VALUES ($1, $2, $3, $4, $5, $6)
        ON CONFLICT (int1, int2, limit_value, str1, str2) 
        DO UPDATE SET 
            hit = fizzbuzz_requests.hit + 1
		RETURNING id, created_at
    `

	var id int
	var createdAt time.Time
	err := s.db.QueryRowContext(ctx, query,
		req.Int1, req.Int2, req.Limit, req.Str1, req.Str2, req.Result,
	).Scan(&id, &createdAt)

	if err != nil {
		return fmt.Errorf("failed to store request: %w", err)
	}

	return nil
}

func (s *Repository) GetFizzBuzz(ctx context.Context, req models.FizzBuzz) (string, error) {
	query := `
        Select result From fizzbuzz_requests 
        Where int1 = $1 And int2 = $2 And limit_value = $3 And str1 = $4 And str2 = $5
    `

	var result string
	err := s.db.QueryRowContext(ctx, query,
		req.Int1, req.Int2, req.Limit, req.Str1, req.Str2,
	).Scan(&result)

	if err != nil {
		return "", fmt.Errorf("failed to get result: %w", err)
	}

	return result, nil
}

func (s *Repository) GetTop(ctx context.Context, limit int) (models.FizzBuzzRequest, error) {
	query := `
        Select * From fizzbuzz_requests 
        order by hit desc
		limit $1
    `

	var result models.FizzBuzzRequest
	err := s.db.QueryRowContext(ctx, query,
		limit,
	).Scan(&result)

	if err != nil {
		return models.FizzBuzzRequest{}, fmt.Errorf("failed to get result: %w", err)
	}

	return result, nil
}
