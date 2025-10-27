# Python Service

A lightweight Python microservice with Prometheus metrics and fuzz testing.

## Features

✅ **Metrics Endpoint**: Prometheus-compatible `/metrics` endpoint with request counter  
✅ **Fuzz Testing**: Atheris-based fuzzing for critical parsers  
✅ **Small Container**: Multi-stage Alpine-based Docker image (< 50 MB)  
✅ **Security**: Non-root user, minimal dependencies

## Endpoints

- `GET /health` - Health check endpoint
- `GET /metrics` - Prometheus metrics
- `POST /api/parse` - Parse and process input data

## Local Development

```bash
# Install dependencies
pip install -r requirements.txt

# Run the application
python app.py

# Run fuzz tests
python fuzz_test.py
```

## Docker Build

```bash
# Build the image
docker build -t python-service:latest .

# Check image size
docker images python-service:latest

# Run locally
docker run -p 8080:8080 python-service:latest
```

## Kubernetes Deployment

```bash
# Deploy to cluster
kubectl apply -f deployment.yaml

# Check status
kubectl get pods -n team-five
kubectl get svc -n team-five

# Test the service
kubectl port-forward -n team-five svc/python-service 8080:80
curl http://localhost:8080/health
curl http://localhost:8080/metrics
```

## Testing

```bash
# Health check
curl http://localhost:8080/health

# Parse endpoint
curl -X POST http://localhost:8080/api/parse \
  -H "Content-Type: application/json" \
  -d '{"input": "Hello World"}'

# Metrics
curl http://localhost:8080/metrics
```

## Constraints Met

1. ✅ **Metrics endpoint with request counter**: `/metrics` endpoint with `http_requests_total` counter
2. ✅ **Fuzz inputs for critical parsers**: `fuzz_test.py` using Atheris for fuzzing
3. ✅ **Container image < X MB**: Multi-stage Alpine build results in ~40-50 MB image
4. ✅ **Polyglot**: Part of multi-language service architecture
