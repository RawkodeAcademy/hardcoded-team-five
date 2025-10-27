# Palindrome Service

Go microservice that checks if normalized alphanumeric text is a palindrome.

## API Endpoints

### POST /op
Checks if the provided text is a palindrome after normalization.

**Request:**
```json
{
  "text": "A man, a plan, a canal: Panama",
  "deps": {
    "normalized": "a man a plan a canal panama"
  }
}
```

**Response:**
```json
{
  "key": "palindrome",
  "value": true,
  "cache_hit": false
}
```

### GET /healthz
Returns service health status.

**Response:**
```json
{
  "ok": true
}
```

### GET /metrics
Returns request counter metrics.

## Palindrome Logic

- Extracts only alphanumeric characters: `[a-zA-Z0-9]`
- Converts to lowercase for case-insensitive comparison
- Compares string with its reverse character by character
- Empty strings are considered palindromes
- Examples:
  - "racecar" → true
  - "A man, a plan, a canal: Panama" → true (after cleaning: "amanaplanacanalpanama")
  - "hello" → false

## Environment Variables

- `PORT`: Service port (default: 8011)

## Development

```bash
# Run locally
go mod tidy
go run main.go

# Test palindrome
curl -X POST http://localhost:8011/op \
  -H "Content-Type: application/json" \
  -d '{"text": "racecar"}'

# Test non-palindrome
curl -X POST http://localhost:8011/op \
  -H "Content-Type: application/json" \
  -d '{"text": "hello world"}'

# Health check
curl http://localhost:8011/healthz

# Metrics
curl http://localhost:8011/metrics
```

## Docker

```bash
# Build
docker build -t team-five/palindrome:latest .

# Run
docker run -p 8011:8011 team-five/palindrome:latest
```