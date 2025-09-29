package services

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fizzbuzz/internal/models"
	"github.com/fizzbuzz/internal/repository"
)

func TestGetFizzBuzz(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	service := NewService(repository.NewRepository(db))

	service.GetFizzBuzz(context.Background(), models.DefaultFizzBuzz)

}

func TestGetStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	service := NewService(repository.NewRepository(db))

	service.GetStats(context.Background())
}
