from flask import Flask, request, jsonify
from prometheus_client import Counter, generate_latest
import os

app = Flask(__name__)

# Prometheus metrics
request_counter = Counter('python_requests_total', 'Total requests', ['endpoint'])

@app.route('/healthz', methods=['GET'])
def health():
    request_counter.labels(endpoint='healthz').inc()
    return jsonify({"ok": True})

@app.route('/metrics', methods=['GET'])
def metrics():
    return generate_latest()

@app.route('/op', methods=['POST'])
def op():
    request_counter.labels(endpoint='op').inc()
    
    data = request.get_json()
    if not data or 'text' not in data:
        return jsonify({"key": "unique_chars", "value": 0, "cache_hit": False})
    
    text = data.get('text', '')
    
    # Count unique characters using a set
    unique_chars = len(set(text))
    
    return jsonify({
        "key": "unique_chars",
        "value": unique_chars,
        "cache_hit": False
    })

if __name__ == '__main__':
    port = int(os.getenv('PORT', 8085))
    app.run(host='0.0.0.0', port=port)
