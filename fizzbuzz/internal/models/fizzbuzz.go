package models

import "time"

type FizzBuzz struct {
	Int1  int    `json:"int1" valid:"range(1|100)"`
	Int2  int    `json:"int2" valid:"range(1|100)"`
	Limit int    `json:"limit" valid:"range(1|100)"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

var DefaultFizzBuzz = FizzBuzz{
	Int1:  3,
	Int2:  5,
	Limit: 100,
	Str1:  "fizz",
	Str2:  "buzz",
}

type FizzBuzzRequest struct {
	ID        int       `json:"id" db:"id"`
	Int1      int       `json:"int1" db:"int1"`
	Int2      int       `json:"int2" db:"int2"`
	Limit     int       `json:"limit" db:"limit_value"`
	Str1      string    `json:"str1" db:"str1"`
	Str2      string    `json:"str2" db:"str2"`
	Result    string    `json:"result" db:"result"`
	Hit       int       `json:"hit" db:"hit"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
