from flask import Flask, request, jsonify
from prometheus_client import Counter, generate_latest
import math
import os

app = Flask(__name__)

# Prometheus metrics
request_counter = Counter('python_requests_total', 'Total requests', ['endpoint'])

def calculate_entropy(text):
    """Calculate Shannon entropy over code points, 3 decimal places"""
    if not text:
        return 0.0
    
    # Count frequency of each character
    char_counts = {}
    for char in text:
        char_counts[char] = char_counts.get(char, 0) + 1
    
    # Calculate entropy
    length = len(text)
    entropy = 0.0
    
    for count in char_counts.values():
        if count > 0:
            probability = count / length
            entropy -= probability * math.log2(probability)
    
    # Round to 3 decimal places
    return round(entropy, 3)

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
        return jsonify({"key": "entropy", "value": 0.0, "cache_hit": False})
    
    text = data.get('text', '')
    entropy = calculate_entropy(text)
    
    return jsonify({
        "key": "entropy",
        "value": entropy,
        "cache_hit": False
    })

if __name__ == '__main__':
    port = int(os.getenv('PORT', 8085))
    app.run(host='0.0.0.0', port=port)
