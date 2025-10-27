# CharCount Service

Go microservice that counts characters in normalized text.

## API Endpoints

### POST /op
Counts characters in the provided text.

**Request:**
```json
{
  "text": "Hello world",
  "deps": {
    "normalized": "hello world"
  }
}
```

**Response:**
```json
{
  "key": "char_count",
  "value": 11,
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
Returns request counter and Prometheus metrics.

## Character Counting Logic

- Uses UTF-8 rune counting for proper Unicode support
- Prefers `deps.normalized` text if available, falls back to `text`
- Returns 0 for empty or invalid input
- Gracefully handles malformed JSON

## Environment Variables

- `PORT`: Service port (default: 8007)

## Development

```bash
# Run locally
go mod tidy
go run main.go

# Test
curl -X POST http://localhost:8007/op \
  -H "Content-Type: application/json" \
  -d '{"text": "Hello world"}'

# Health check
curl http://localhost:8007/healthz

# Metrics
curl http://localhost:8007/metrics
```

## Docker

```bash
# Build
docker build -t team-five/charcount:latest .

# Run
docker run -p 8007:8007 team-five/charcount:latest
```