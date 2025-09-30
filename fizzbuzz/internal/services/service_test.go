package services

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fizzbuzz/internal/models"
	"github.com/fizzbuzz/internal/repository"
	"github.com/fizzbuzz/utils"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestGetFizzBuzz(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	expectedFizzBuzz := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz"

	mock.ExpectQuery("SELECT result FROM fizzbuzz_requests").
		WithArgs(3, 5, 10, "fizz", "buzz").
		WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("INSERT INTO fizzbuzz_requests.*").
		WithArgs(3, 5, 10, "fizz", "buzz", expectedFizzBuzz).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).AddRow(1, time.Now()))

	smallFizzBuzz := models.DefaultFizzBuzz
	smallFizzBuzz.Limit = 10
	smallFizzBuzzEncoded, _ := utils.EncodeStruct(smallFizzBuzz)

	redis, redMock := redismock.NewClientMock()
	redMock.ExpectGet(smallFizzBuzzEncoded).RedisNil()
	redMock.ExpectSet(smallFizzBuzzEncoded, expectedFizzBuzz, time.Hour).RedisNil()

	service := NewService(repository.NewRepository(db), redis)

	fizzBuzz, err := service.GetFizzBuzz(context.Background(), smallFizzBuzz)
	assert.Nil(t, err)
	assert.Equal(t, expectedFizzBuzz, fizzBuzz)
	assert.NoError(t, redMock.ExpectationsWereMet())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetFizzBuzzSameInt(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	expectedFizzBuzz := "1,2,fizzbuzz,4,5,fizzbuzz,7,8,fizzbuzz,10"

	mock.ExpectQuery("SELECT result FROM fizzbuzz_requests").
		WithArgs(3, 3, 10, "fizz", "buzz").
		WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("INSERT INTO fizzbuzz_requests.*").
		WithArgs(3, 3, 10, "fizz", "buzz", expectedFizzBuzz).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).AddRow(1, time.Now()))

	sameFizzBuzz := models.DefaultFizzBuzz
	sameFizzBuzz.Limit = 10
	sameFizzBuzz.Int2 = sameFizzBuzz.Int1
	sameFizzBuzzEncoded, _ := utils.EncodeStruct(sameFizzBuzz)

	redis, redMock := redismock.NewClientMock()
	redMock.ExpectGet(sameFizzBuzzEncoded).RedisNil()
	redMock.ExpectSet(sameFizzBuzzEncoded, expectedFizzBuzz, time.Hour).RedisNil()

	service := NewService(repository.NewRepository(db), redis)

	fizzBuzz, err := service.GetFizzBuzz(context.Background(), sameFizzBuzz)
	assert.Nil(t, err)
	assert.Equal(t, expectedFizzBuzz, fizzBuzz)
	assert.NoError(t, redMock.ExpectationsWereMet())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	expectedFizzBuzzRequest := models.FizzBuzzRequest{
		ID:        1,
		Int1:      models.DefaultFizzBuzz.Int1,
		Int2:      models.DefaultFizzBuzz.Int2,
		Limit:     10,
		Str1:      models.DefaultFizzBuzz.Str1,
		Str2:      models.DefaultFizzBuzz.Str2,
		Hit:       5,
		Result:    "1,2,fizzbuzz,4,5,fizzbuzz,7,8,fizzbuzz,10",
		CreatedAt: time.Now(),
	}

	mock.ExpectQuery("SELECT \\* FROM fizzbuzz_requests ORDER BY hit DESC LIMIT \\$1").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "int1", "int2", "limit_value", "str1", "str2", "result", "hit", "created_at",
		}).AddRow(
			expectedFizzBuzzRequest.ID,
			expectedFizzBuzzRequest.Int1,
			expectedFizzBuzzRequest.Int2,
			expectedFizzBuzzRequest.Limit,
			expectedFizzBuzzRequest.Str1,
			expectedFizzBuzzRequest.Str2,
			expectedFizzBuzzRequest.Result,
			expectedFizzBuzzRequest.Hit,
			expectedFizzBuzzRequest.CreatedAt,
		))

	service := NewService(repository.NewRepository(db), nil)

	fizzBuzzStats, err := service.GetStats(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, expectedFizzBuzzRequest, fizzBuzzStats)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetFizzBuzzCached(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	expectedFizzBuzz := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz"

	mock.ExpectQuery("INSERT INTO fizzbuzz_requests.*").
		WithArgs(3, 5, 10, "fizz", "buzz", expectedFizzBuzz).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).AddRow(1, time.Now()))

	smallFizzBuzz := models.DefaultFizzBuzz
	smallFizzBuzz.Limit = 10
	smallFizzBuzzEncoded, _ := utils.EncodeStruct(smallFizzBuzz)

	redis, redMock := redismock.NewClientMock()
	redMock.ExpectGet(smallFizzBuzzEncoded).SetVal(expectedFizzBuzz)
	redMock.ExpectSet(smallFizzBuzzEncoded, expectedFizzBuzz, time.Hour).RedisNil()

	service := NewService(repository.NewRepository(db), redis)

	fizzBuzz, err := service.GetFizzBuzz(context.Background(), smallFizzBuzz)
	assert.Nil(t, err)
	assert.Equal(t, expectedFizzBuzz, fizzBuzz)
	assert.NoError(t, redMock.ExpectationsWereMet())
	assert.NoError(t, mock.ExpectationsWereMet())
}
