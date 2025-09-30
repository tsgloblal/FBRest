# FBRest

## Overview
This API provides a flexible implementation of the classic Fizz Buzz game with customizable parameters and usage statistics.

## Old Fizz Buzz 
The traditional Fizz Buzz replaces numbers with words based on divisibility:
Multiples of 3 → "fizz"
Multiples of 5 → "buzz"
Multiples of 15 → "fizzbuzz"
Example output: 1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...

## New Fizz Buzz 
The new implementation allows full customization:
- int1: First divisor (replaces multiples with str1)
- int2: Second divisor (replaces multiples with str2)
- limit: Maximum number to process
- str1: String to replace int1 multiples
- str2: String to replace int2 multiples
- Combined: Multiples of both int1 and int2 are replaced with str1+str2

## Live Demo

**Production URL:** [fbrest-production.up.railway.app](https://fbrest-production.up.railway.app)

## API Endpoints
### GET /swagger/index.html

Interactive API documentation

### GET /health

Health check endpoint


### GET /api/fizzbuzz

Generate a customized Fizz Buzz sequence.

#### Parameters:
- int1 (integer): First divisor
- int2 (integer): Second divisor
- limit (integer): Maximum number to process
- str1 (string): Replacement for int1 multiples
- str2 (string): Replacement for int2 multiples

#### Example:
```
GET /getFizzBuzz?int1=3&int2=5&limit=15&str1=fizz&str2=buzz
```

#### Response:
```
{
    "data": "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz",
    "message": "fizz-buzz retrieved successfully"
}
```

### GET /api/stats

Return usage statistics showing the most frequent request.

#### Response:
```
{
    "data": {
        "id": 10,
        "int1": 3,
        "int2": 5,
        "limit": 10,
        "str1": "fizz",
        "str2": "buzz",
        "result": "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz",
        "hit": 10,
        "created_at": "2025-09-30T11:54:04.901263Z",
        "updated_at": "2025-09-30T11:54:04.901263Z"
    },
    "message": "fizz-buzz stats retrieved successfully"
}
```

## Local Development
### Prerequisites

- Docker & Docker Compose
- Go 1.21+ (for development)

### Running Locally

1. **Clone the repository**
   ```bash
   git clone https://github.com/tsgloblal/FBRest.git
   cd FBRest
   ```

2. **Start the development environment**
   ```bash
   make dev-docker
   ```
   This will start:
   - Fizz Buzz API (port 8080)
   - PostgreSQL database (port 5432)
   - Redis cache (port 6379)

3. **Access the API**
   - API: http://localhost:8080
   - Swagger UI: http://localhost:8080/swagger/index.html
   - Health Check: http://localhost:8080/health

### Development Commands

```bash
# Runs all checks before deployment
make pre-deploy

# Generate Swagger documentation
make swagg

# Start development environment
make dev-docker

# Stop development environment
make down
```


## Tech Stack
- #### Redis Cache
- #### Swagger/OpenAPI
- #### PostgreSQL
- #### Docker
- #### Railway

