# Project-root Makefile - Python & Go services

# Prevent caching issues and ensure all targets run fresh every time
.PHONY: health lint-python format-python test-python build-python run-docker-python ci-python test-go logs

# ✅ Ping /health endpoint to confirm service is running
health:
	./python-ai/health_check.sh

# 🧼 Lint Python code using ruff (report only)
lint-python:
	ruff check python-ai

# 🛠 Auto-fix style issues using ruff (safe fixes only)
format-python:
	ruff check --fix python-ai

# 🧪 Run Python unittests
test-python:
	PYTHONPATH=python-ai python3 -m unittest discover -s python-ai/tests -p "test_*.py"

# 🐳 Build Docker image for Python AI service
build-python:
	docker build -t python-ai-service ./python-ai

# 🐳 Run Python AI service in Docker
run-docker-python:
	docker run -p 8001:8001 --env-file=./python-ai/.env -e ENV=production python-ai-service

# 📄 View recent logs from the Python AI service container
logs:
	docker ps -q --filter ancestor=python-ai-service | xargs -r docker logs --tail=50

# ✅ CI target — run lint, test, and health checks in sequence
ci-python: lint-python test-python health

# Run go tests
test-go:
	go test ./go-backend/...
