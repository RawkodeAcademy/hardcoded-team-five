from flask import Flask, request, jsonify
from prometheus_client import Counter, generate_latest
import unicodedata
import re
import os

app = Flask(__name__)

# Prometheus metrics
request_counter = Counter('python_requests_total', 'Total requests', ['endpoint'])

def strip_diacritics(text):
    """Remove diacritical marks from text"""
    nfd = unicodedata.normalize('NFD', text)
    return ''.join(char for char in nfd if unicodedata.category(char) != 'Mn')

def normalize_text(text):
    """NFKC normalize, lowercase, collapse whitespace, strip diacritics"""
    if not text:
        return None
    
    # NFKC normalization
    normalized = unicodedata.normalize('NFKC', text)
    
    # Strip diacritics
    normalized = strip_diacritics(normalized)
    
    # Lowercase
    normalized = normalized.lower()
    
    # Collapse whitespace
    normalized = re.sub(r'\s+', ' ', normalized).strip()
    
    return normalized if normalized else None

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
        return jsonify({"key": "normalized", "value": None, "cache_hit": False})
    
    text = data.get('text', '')
    normalized = normalize_text(text)
    
    return jsonify({
        "key": "normalized",
        "value": normalized,
        "cache_hit": False
    })

if __name__ == '__main__':
    port = int(os.getenv('PORT', 8085))
    app.run(host='0.0.0.0', port=port)
