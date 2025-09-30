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
	return &Repository{
		db: db,
	}
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
        SELECT result FROM fizzbuzz_requests 
        WHERE int1 = $1 AND int2 = $2 AND limit_value = $3 AND str1 = $4 AND str2 = $5
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

func (s *Repository) GetTop(ctx context.Context, limit int) ([]models.FizzBuzzRequest, error) {
	query := `
        SELECT * FROM fizzbuzz_requests 
        ORDER BY hit DESC
		LIMIT $1
    `

	rows, err := s.db.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get stats: %w", err)
	}
	defer rows.Close()

	var results []models.FizzBuzzRequest
	for rows.Next() {
		var result models.FizzBuzzRequest
		err := rows.Scan(
			&result.ID,
			&result.Int1,
			&result.Int2,
			&result.Limit,
			&result.Str1,
			&result.Str2,
			&result.Result,
			&result.Hit,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}
	return results, nil
}
