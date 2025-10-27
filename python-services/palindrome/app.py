from flask import Flask, request, jsonify
from prometheus_client import Counter, generate_latest
import re
import os

app = Flask(__name__)

# Prometheus metrics
request_counter = Counter('python_requests_total', 'Total requests', ['endpoint'])

def is_palindrome(text):
    """Check if alphanumeric text equals its reverse"""
    if not text:
        return False
    
    # Keep only alphanumeric characters and lowercase
    alphanumeric = re.sub(r'[^a-z0-9]', '', text.lower())
    
    if not alphanumeric:
        return False
    
    # Check if it's a palindrome
    return alphanumeric == alphanumeric[::-1]

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
        return jsonify({"key": "palindrome", "value": False, "cache_hit": False})
    
    text = data.get('text', '')
    palindrome = is_palindrome(text)
    
    return jsonify({
        "key": "palindrome",
        "value": palindrome,
        "cache_hit": False
    })

if __name__ == '__main__':
    port = int(os.getenv('PORT', 8085))
    app.run(host='0.0.0.0', port=port)
