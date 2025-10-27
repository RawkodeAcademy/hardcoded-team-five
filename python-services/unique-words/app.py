from flask import Flask, request, jsonify
from prometheus_client import Counter, generate_latest
import logging

app = Flask(__name__)

# Prometheus metrics
request_counter = Counter('python_requests_total', 'Total requests processed')

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


@app.route('/healthz', methods=['GET'])
def health():
    """Health check endpoint"""
    return jsonify({"ok": True}), 200


@app.route('/metrics', methods=['GET'])
def metrics():
    """Prometheus metrics endpoint"""
    metrics_output = generate_latest()
    return metrics_output, 200, {'Content-Type': 'text/plain; charset=utf-8'}


@app.route('/op', methods=['POST'])
def op():
    """Operation endpoint - counts unique words"""
    request_counter.inc()
    
    try:
        data = request.get_json()
        
        if not data or 'text' not in data or not isinstance(data['text'], str):
            return jsonify({
                "key": "unique_words",
                "value": 0,
                "cache_hit": False
            }), 200
        
        # Split by spaces and count unique words
        text = data['text']
        words = text.split()
        unique_count = len(set(words))
        
        return jsonify({
            "key": "unique_words",
            "value": unique_count,
            "cache_hit": False
        }), 200
    
    except Exception as e:
        logger.error(f"Processing error: {e}")
        return jsonify({
            "key": "unique_words",
            "value": 0,
            "cache_hit": False
        }), 200


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8085)
