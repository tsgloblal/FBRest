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

## API Endpoints
### GET /getFizzBuzz

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
1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz
```

### GET /stats

Return usage statistics showing the most frequent request.

#### Example:
```
GET /stats
```

#### Response:
```
{
  "mostFrequentRequest": {
    "int1": 3,
    "int2": 5,
    "limit": 100,
    "str1": "fizz",
    "str2": "buzz"
  },
  "hits": 42
}
```


Tech Stack:
- Redis Cache
- Swagger/OpenAPI
- PostgreSQL
- Docker
