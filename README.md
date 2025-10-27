# hardcoded-team-five
Repository team-five for hardcoded

## Services

### Python Service

A lightweight Python microservice built with Flask, featuring Prometheus metrics and comprehensive fuzz testing.

**Key Features:**
- ✅ **Metrics Endpoint**: Prometheus-compatible `/metrics` endpoint with `http_requests_total` counter tracking all HTTP requests by method, endpoint, and status
- ✅ **Fuzz Testing**: Atheris-based fuzzing for critical input parsers to ensure robustness against malformed data
- ✅ **Small Container**: Multi-stage Alpine-based Docker image (~40-50 MB) for minimal resource footprint
- ✅ **Production Ready**: Gunicorn WSGI server, health checks, non-root user, and resource limits

**Endpoints:**
- `GET /health` - Health check endpoint for Kubernetes probes
- `GET /metrics` - Prometheus metrics endpoint
- `POST /api/parse` - Parse and process input data with validation

**Technology Stack:**
- Python 3.11 (Alpine-based)
- Flask web framework
- Prometheus client for metrics
- Atheris for fuzz testing
- Gunicorn for production serving

See [python-service/README.md](python-service/README.md) for detailed documentation, build instructions, and deployment guide.
