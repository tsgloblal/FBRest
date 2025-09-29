-- +goose Up
CREATE TABLE fizzbuzz_requests (
    id SERIAL PRIMARY KEY,
    int1 INTEGER NOT NULL CHECK (int1 >= 1 AND int1 <= 100),
    int2 INTEGER NOT NULL CHECK (int2 >= 1 AND int2 <= 100),
    limit_value INTEGER NOT NULL CHECK (limit_value >= 1 AND limit_value <= 1000),
    str1 VARCHAR(10) NOT NULL CHECK (LENGTH(str1) >= 1 AND LENGTH(str1) <= 10),
    str2 VARCHAR(10) NOT NULL CHECK (LENGTH(str2) >= 1 AND LENGTH(str2) <= 10),
    result TEXT NOT NULL,
    hit INTEGER DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- Unique constraint for upsert
    UNIQUE(int1, int2, limit_value, str1, str2)
);

CREATE INDEX idx_fizzbuzz_requests_params ON fizzbuzz_requests(int1, int2, limit_value, str1, str2);