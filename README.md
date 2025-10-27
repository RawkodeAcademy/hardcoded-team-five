# hardcoded-team-five
Repository team-five for hardcoded

## Services

### Python Services

A collection of lightweight Python microservices built with Flask, featuring Prometheus metrics and production-ready deployment configurations.

**Common Features:**
- ✅ **Metrics Endpoint**: Prometheus-compatible `/metrics` endpoint with request counter
- ✅ **Health Checks**: `/healthz` endpoint returning `{"ok": true}`
- ✅ **Small Container**: Python 3.11-slim based images (~40-50 MB) for minimal resource footprint
- ✅ **Production Ready**: Gunicorn WSGI server with 2 workers and 2 threads, resource limits, and health probes

**Technology Stack:**
- Python 3.11 (slim-based)
- Flask 3.0.0 web framework
- Prometheus client 0.19.0 for metrics
- Gunicorn 21.2.0 for production serving

#### Services Implemented:

##### 1. Unique Words (`python-services/unique-words`)
Counts unique words from input text.
- **Endpoint**: `POST /op` with `{"text": "string"}`
- **Response**: `{"key": "unique_words", "value": number, "cache_hit": false}`

##### 2. Unique Chars (`python-services/unique-chars`)
Counts distinct characters in text.
- **Endpoint**: `POST /op` with `{"text": "string"}`
- **Response**: `{"key": "unique_chars", "value": number, "cache_hit": false}`

##### 3. Normalizer (`python-services/normalizer`)
NFKC normalize, lowercase, collapse whitespace, strip diacritics.
- **Endpoint**: `POST /op` with `{"text": "string"}`
- **Response**: `{"key": "normalized", "value": "string|null", "cache_hit": false}`

##### 4. Entropy (`python-services/entropy`)
Calculates Shannon entropy over code points (3 decimal places).
- **Endpoint**: `POST /op` with `{"text": "string"}`
- **Response**: `{"key": "entropy", "value": number, "cache_hit": false}`

##### 5. Palindrome (`python-services/palindrome`)
Checks if alphanumeric text equals its reverse.
- **Endpoint**: `POST /op` with `{"text": "string"}`
- **Response**: `{"key": "palindrome", "value": boolean, "cache_hit": false}`

**Deployment:**
Each service includes a Kubernetes deployment manifest with:
- 2 replicas for high availability
- Health probes (liveness and readiness)
- Resource limits (64-128 Mi memory, 50-100m CPU)
- Prometheus annotations for metrics scraping
