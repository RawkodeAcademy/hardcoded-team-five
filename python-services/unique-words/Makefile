.PHONY: help build run test fuzz deploy clean

help:
	@echo "Available targets:"
	@echo "  build   - Build Docker image"
	@echo "  run     - Run locally with Python"
	@echo "  test    - Run tests"
	@echo "  fuzz    - Run fuzz tests"
	@echo "  deploy  - Deploy to Kubernetes"
	@echo "  clean   - Clean up"

build:
	docker build -t python-service:latest .
	@echo "Image size:"
	@docker images python-service:latest --format "{{.Size}}"

run:
	python app.py

test:
	python -m pytest -v

fuzz:
	python fuzz_test.py -max_total_time=60

deploy:
	kubectl apply -f deployment.yaml
	kubectl rollout status deployment/python-service -n team-five

clean:
	find . -type d -name __pycache__ -exec rm -rf {} + 2>/dev/null || true
	find . -type f -name "*.pyc" -delete
	docker rmi python-service:latest 2>/dev/null || true
